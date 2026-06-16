// Package notifier sends alerts to humans (Slack, PagerDuty, email, webhook).
package notifier

import "log"

// Notifier delivers a message to an external channel.
type Notifier interface {
	Notify(msg string) error
}

// LogNotifier writes notifications to the logger. Real channels (Slack,
// PagerDuty) come later behind this same interface.
type LogNotifier struct{ logger *log.Logger }

// NewLogNotifier returns a Notifier that logs messages.
func NewLogNotifier(logger *log.Logger) *LogNotifier { return &LogNotifier{logger: logger} }

// Notify writes the message to the log.
func (n *LogNotifier) Notify(msg string) error {
	n.logger.Printf("NOTIFY %s", msg)
	return nil
}
