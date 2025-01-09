package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// for consumer
func getKafkaReader(kafkaURL, topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset: kafka.LastOffset,
	})
}

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}


type StockInfo struct {
	Message string `json:"message"`
	Type string `json:"type"`
}

// mua bán chứng khoán 
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s
}

// Consumer hóng mua ATC
func RegisterConsumerATC(id int) {
	// group consumer
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer (%d) Hong Phien ATC::", id)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v\n", id, err)
			continue
		}

		fmt.Printf("Consumer(%d), hong topic: %v, partition: %d, offset: %d, time: %s = %s\n",
			id, m.Topic, m.Partition, m.Offset,
			time.Unix(m.Time.Unix(), 0).String(),
			string(m.Value))
	}
}

func actionStock(c *gin.Context) {
	var body StockInfo
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonBody, _ := json.Marshal(body)
	msg := kafka.Message{
		Key:   []byte(body.Type),
		Value: jsonBody,
	}
	// view message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON (200, gin.H{
		"error": "",
		"message": "Stock action sent to Kafka",
	})
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("/action-stock", actionStock)

	// regist hong
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8080")
}