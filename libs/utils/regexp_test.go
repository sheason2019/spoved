package utils

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	reg := regexp.MustCompile("([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]")
	name := "service-proj-id-1"

	fmt.Println(reg.Match([]byte(name)))
}
