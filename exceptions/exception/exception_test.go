package exception_test

import (
	"errors"
	"testing"

	"github.com/sheason2019/spoved/exceptions/exception"
)

func TestException(t *testing.T) {
	err := errors.New("Error")
	e := exception.New(err)

	e2 := e.Wrap()
	e2.Print()
}

func TestExceptionPanic(t *testing.T) {
	e := exception.New(errors.New("Panic Test"))
	defer func() {
		recover()
	}()

	e.Panic()
}
