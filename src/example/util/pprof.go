package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go func() {
		if err := http.ListenAndServe(":7080", nil); err != nil {
			fmt.Printf("http.ListenAndServe failed, err:%+v", err)
		}
	}()
	//runtime.SetBlockProfileRate(100)
	log.Println("===> Start")
	readMemStatus()
	test()
	readMemStatus()
	log.Println("===> force gc")
	runtime.GC()
	log.Println("====> done")

	readMemStatus()
	go func() {
		for {
			readMemStatus()
			time.Sleep(10 * time.Second)
		}
	}()
	time.Sleep(3600 * time.Second)

}
