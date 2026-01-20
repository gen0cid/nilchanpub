package sending

type Notifier interface {
	Send() error
}
type NotificationModule struct{}

func NewNotificationModule() NotificationModule {
	return NotificationModule{}
}
func (nm NotificationModule) SendNotifications(notifiers []Notifier) []error {
	var errors []error
	for _, v := range notifiers {
		if err := v.Send(); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
