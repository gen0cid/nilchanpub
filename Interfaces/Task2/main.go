package main

import (
	"Task2/sending"
	"fmt"
	"log"
)

func main() {
	nm := sending.NewNotificationModule()

	notifications := []sending.Notifier{
		sending.EmailNotifier{
			To:      "user@example.com",
			Subject: "Привет из Go!",
			Body:    "Это тестовое письмо.",
		},
		sending.SMSNotifier{
			To:   "+1234567890",
			Text: "Ваш код: 9876",
		},
		sending.PushNotifier{
			DeviceID: "device_abc123",
			Title:    "Новое сообщение",
			Message:  "Вы получили push-уведомление!",
		},
		sending.TgNotifier{
			To:   "+1234567890",
			Text: "это сообщение из телеграмма",
		},
	}

	err := nm.SendNotifications(notifications)
	if len(err) == 0 {
		fmt.Println("✅ Все уведомления успешно отправлены!")
	} else {
		log.Printf("❌ Ошибки при отправке: %v", err)
	}
}
