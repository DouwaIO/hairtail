package task

import (
//	"github.com/DouwaIO/hairtail/src/pipeline"
)

type Pipeline interface {
	Pipeline() error
}

type Service interface {
	Service() error
}

