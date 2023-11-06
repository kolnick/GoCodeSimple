package regex

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	matched, _ := regexp.Match("^\\d{6,15}$", []byte("123456789"))
	fmt.Println(matched)
}
