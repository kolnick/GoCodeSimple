package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCpu(t *testing.T) {

	fmt.Println("cpu个数:", runtime.NumCPU())

}
