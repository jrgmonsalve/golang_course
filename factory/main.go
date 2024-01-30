package main

import "fmt"

// Factory design pattern
type INotification interface {
	Send()
	GetSender() ISender
}

type ISender interface {
	GetMethod() string
	GetChannel() string
}

// SMS
type SMSNotification struct {
}

func (SMSNotification) Send() {
	fmt.Println("Sending SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetChannel() string {
	return "Twilio"
}

// EMAIL
type EmailNotification struct {
}

func (EmailNotification) Send() {
	fmt.Println("Sending Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetChannel() string {
	return "Gmail"
}

func getNotificator(notificationType string) (INotification, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "EMAIL":
		return &EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("invalid notification type")
	}
}

func sendNotification(notification INotification) {
	notification.Send()
}

func getMethod(notification INotification) {
	fmt.Println("Get Method: ", notification.GetSender().GetMethod())
}

func main() {
	smsNotification, _ := getNotificator("SMS")
	sendNotification(smsNotification)
	getMethod(smsNotification)

	emailNotification, _ := getNotificator("EMAIL")
	sendNotification(emailNotification)
	getMethod(emailNotification)
}
