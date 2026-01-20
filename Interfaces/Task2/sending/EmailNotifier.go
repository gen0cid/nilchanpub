package sending

import "fmt"

type EmailNotifier struct {
	To      string
	Subject string
	Body    string
}

func (e EmailNotifier) Send() error {
	fmt.Printf("📧 Отправка email на %s: %s\n", e.To, e.Subject)
	fmt.Printf("Текст сообщения:  %s", e.Body)
	fmt.Println("")
	return nil
}
