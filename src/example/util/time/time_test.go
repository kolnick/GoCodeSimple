package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	time1 := time.Now()
	fmt.Println("1 ", time1)

	fmt.Println("2 localTime ", time1.Local())

	fmt.Println("3 utcTime ", time1.UTC())

	time2 := time.Now()
	// sub time
	fmt.Println("4 sub time ", time2.Sub(time1).Milliseconds())
	// millisecond
	fmt.Println("5 millisecond ", time1.UnixMilli())
}

func TestParseTime(t *testing.T) {

	duration, _ := time.ParseDuration("1h30m")
	fmt.Println(duration)

	time, _ := time.Parse("2006-01-02 12:12:12", "2023-10-19 19:19:19")
	fmt.Println(time.Local())
}
