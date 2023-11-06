package os

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExec(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("cmd", "/C", "start D:")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
