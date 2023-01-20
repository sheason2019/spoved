package exception

import (
	"fmt"
	"runtime"
	"strings"
)

type Exception struct {
	CallStack []string
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
	row = append(row, fmt.Sprintf("\033[1;37;41m ERROR \033[0m %s\n", e.error))
	for _, str := range e.CallStack {
		row = append(row, fmt.Sprintln("\t"+str))
	}

	return strings.Join(row, "\n")
}

func (e *Exception) Panic() {
	fmt.Printf("\033[1;37;41m PANIC! \033[0m %s\n", e.error)
	for _, str := range e.CallStack {
		fmt.Println("\t" + str)
	}
	panic(fmt.Sprintf("\033[1;31m%s\033[0m", e.error))
}
