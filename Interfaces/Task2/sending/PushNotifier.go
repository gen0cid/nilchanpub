package sending

import "fmt"

type PushNotifier struct {
	DeviceID string
	Title    string
	Message  string
}

func (p PushNotifier) Send() error {
	fmt.Printf("📧 Отправка push на %s: %s\n", p.DeviceID, p.Title)
	fmt.Printf("Текст сообщения:  %s", p.Message)
	fmt.Println()
	return nil
}
