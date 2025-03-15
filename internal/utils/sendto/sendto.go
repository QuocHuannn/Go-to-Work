package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/config"
)

// SMTP configuration
const (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = "587"
)

// Get SMTP credentials from environment variables or use defaults
var (
	// Temporarily hardcode credentials for testing
	SMTPUsername = "truonghuan0709@gmail.com" // Temporarily hardcoded
	SMTPPassword = "jonw tvvu frcz vjbt"      // Temporarily hardcoded App Password
)

// Original code with environment variables - commented out for now
/*
var (
	SMTPUsername = getEnvOrDefault("SMTP_USERNAME", "your-email@gmail.com")
	SMTPPassword = getEnvOrDefault("SMTP_PASSWORD", "your-app-password")
)

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	fmt.Printf("DEBUG ENV: Reading %s = '%s'\n", key, value) // Debug log
	if value == "" {
		return defaultValue
	}
	return value
}
*/

// Placeholder for the getEnvOrDefault function to avoid compiler errors
func getEnvOrDefault(key, defaultValue string) string {
	return defaultValue
}

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s", mail.Body)

	return msg
}

// SendTextEmailOTP sends OTP verification email using HTML template
func SendTextEmailOTP(to []string, from string, otp string) error {
	// Get SMTP settings from config
	smtpConfig := config.Cfg.SMTP

	// Log current settings for debugging
	global.Logger.Info("Attempting to send email",
		zap.String("smtp_host", smtpConfig.Host),
		zap.String("smtp_port", smtpConfig.Port),
		zap.String("smtp_username", smtpConfig.Username),
		zap.Strings("recipients", to),
		zap.String("otp", otp))

	// Load HTML template
	templatePath := filepath.Join("templates", "otp-auth.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		global.Logger.Error("Failed to parse email template", zap.Error(err))
		return err
	}

	// Execute template with OTP data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, map[string]string{"OTP": otp}); err != nil {
		global.Logger.Error("Failed to execute email template", zap.Error(err))
		return err
	}

	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: smtpConfig.FromName},
		To:      to,
		Subject: "OTP Verification Code",
		Body:    body.String(),
	}

	messageMail := BuildMessage(contentEmail)

	// Create authentication
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	// Send email
	err = smtp.SendMail(
		smtpConfig.Host+":"+smtpConfig.Port,
		auth,
		smtpConfig.Username,
		to,
		[]byte(messageMail),
	)

	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	global.Logger.Info("Email sent successfully", zap.Strings("to", to))
	return nil
}

func SendTemplateEmailOTP(to []string, from string, htmlTemplate string, templateData map[string]interface{}) error {
	body, err := getMailTemplate(htmlTemplate, templateData)
	if err != nil {
		global.Logger.Error("Failed to get mail template", zap.Error(err))
		return err
	}

	return send(to, from, body)
}

func getMailTemplate(nameTemplate string, templateData map[string]interface{}) (string, error) {
	htmlTemplate, err := template.ParseFiles("templates/" + nameTemplate)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := htmlTemplate.Execute(&body, templateData); err != nil {
		return "", err
	}

	return body.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	// Get SMTP settings from config
	smtpConfig := config.Cfg.SMTP

	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: smtpConfig.FromName},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)

	// Create authentication
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	// Send email
	err := smtp.SendMail(
		smtpConfig.Host+":"+smtpConfig.Port,
		auth,
		smtpConfig.Username,
		to,
		[]byte(messageMail),
	)

	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	global.Logger.Info("Email sent successfully", zap.Strings("to", to))
	return nil
}
