package sending

import "fmt"

type TgNotifier struct {
	To   string
	Text string
}

func (t TgNotifier) Send() error {
	fmt.Printf("📧 Отправка SMS пользователю через телеграмм %s: %s\n", t.To, t.Text)
	return nil
}
