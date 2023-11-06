package cron_robfig

//https://cron.qqe2.com/  cron net

// https://github.com/robfig/cron
import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
)

func TestSecondJob(t *testing.T) {
	c := cron.New()

	c.AddFunc("* * * * * ?", func() {
		fmt.Println("执行了定时任务")
	})
	c.Run()
}
