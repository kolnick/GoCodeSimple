package simple_factory

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	handler := getMatchHandler(1)
	match := handler.doMatch(1)

	fmt.Println(match.matchCode)
	handler2 := getMatchHandler(2)
	match2 := handler2.doMatch(2)
	fmt.Println(match2.matchCode)
}
