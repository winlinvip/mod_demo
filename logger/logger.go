package logger

import (
	"fmt"
	"private.me/show/core"
	"github.com/pkg/errors"
)

type Logger struct {
}

func New(tank string) (*Logger,error) {
	if tank != "console" {
		return nil, errors.Errorf("invalid tank %v", tank)
	}

	return &Logger{}, nil
}

func (v *Logger) Printf(format string, args... interface{}) {
	fmt.Println(fmt.Sprintf("[%v] %v", core.Version(), fmt.Sprintf(format, args...)))
}
