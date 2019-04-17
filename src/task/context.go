package task

import (
	//"fmt"
	"context"
	//"github.com/Sirupsen/logrus"
)

const key = "task"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Remote associated with this context.
func FromContext(c context.Context) Task {
	return c.Value(key).(Task)
}

// ToContext adds the Task to this context if it supports
// the Setter interface.
func ToContext(c Setter, t Task) {
	c.Set(key, t)
}
