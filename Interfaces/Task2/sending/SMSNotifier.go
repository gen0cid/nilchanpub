package sending

import "fmt"

type SMSNotifier struct {
	To   string
	Text string
}

func (s SMSNotifier) Send() error {
	fmt.Printf("📧 Отправка SMS пользователю %s: %s\n", s.To, s.Text)
	return nil
}
