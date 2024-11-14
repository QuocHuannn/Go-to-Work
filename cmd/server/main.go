package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}
	slice := arr[1:4]       // Slice từ vị trí 1 đến 3, dung lượng mặc định là từ 1 đến 5
	fullSlice := arr[1:4:5] // Slice từ vị trí 1 đến 3, nhưng dung lượng giới hạn từ 1 đến 4

	fmt.Println("Slice:", slice)                  // Kết quả: [20 30 40]
	fmt.Println("Length of slice:", len(slice))   // Kết quả: 3
	fmt.Println("Capacity of slice:", cap(slice)) // Kết quả: 5 (dung lượng mặc định từ 1 đến 5)

	fmt.Println("Full Slice:", fullSlice)                 // Kết quả: [20 30 40]
	fmt.Println("Length of fullSlice:", len(fullSlice))   // Kết quả: 3
	fmt.Println("Capacity of fullSlice:", cap(fullSlice)) // Kết quả: 4 (giới hạn dung lượng từ 1 đến 4)
}
