package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {

	set := Set{}
	set.init()
	set.Add("1")
	set.Add(2)
	set.Add(3)

	fmt.Printf("%v", set)

	set.Remove("2")
	fmt.Printf("%v", set)

	set.Remove(2)
	fmt.Printf("%v", set)

}
