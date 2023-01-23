package exception

import (
	"fmt"
	"runtime"
	"strings"
)

type Exception struct {
	CallStack []string
	// Code -1表示不向外展示的异常
	Code int
	error
}

func (e *Exception) Wrap() *Exception {
	_, file, line, _ := runtime.Caller(1)
	e.CallStack = append(e.CallStack, fmt.Sprintf("%s line: %d", file, line))

	return e
}

func New(err error) *Exception {
	_, file, line, _ := runtime.Caller(1)
	return &Exception{
		error:     err,
		CallStack: []string{fmt.Sprintf("%s line: %d", file, line)},
	}
}

func (e *Exception) Print() string {
	row := []string{}
	row = append(row, fmt.Sprintf("\033[1;37;41m ERROR \033[0m \033[31m%s", e.error))
	for _, str := range e.CallStack {
		row = append(row, fmt.Sprint("\t\033[31m"+str+"\033[0m"))
	}

	return strings.Join(row, "\n")
}

func (e *Exception) Panic() {
	_, file, line, _ := runtime.Caller(1)
	e.CallStack = append(e.CallStack, fmt.Sprintf("%s line: %d", file, line))
	panic(e)
}
