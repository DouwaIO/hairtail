package task

import (
)

type Task interface {
	MQ(string, string, string, string, string) error
}
