package date

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	now := time.Now()
	millSecond := now.UnixMilli()
	fmt.Printf("时间戳: %v\n", millSecond)
	fmt.Printf("%v~~对应的类型是%T\n", now, now)
	fmt.Printf("年 %v \n", now.Year())
	fmt.Printf("月 %v \n", int(now.Month()))
	fmt.Printf("日 %v \n", now.Day())
	fmt.Printf("时 %v \n", now.Hour())
	fmt.Printf("分 %v \n", now.Minute())
	fmt.Printf("秒 %v \n", now.Second())
	fmt.Println("--------")
	format := now.Format(time.UnixDate)
	fmt.Println(format)
}

func TestTemplateFormat(t *testing.T) {
	template := "20060102"
	format := time.Now().Format(template)
	fmt.Println(format)
}
