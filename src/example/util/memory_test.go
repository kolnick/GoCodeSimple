package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
	"time"
)

func readMemStatus() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf("===> Alloc:%d(Bytes) HeapIdle:%d(Bytes) HeapReleased %d(Bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test() {
	containers := make([]int, 8)
	log.Println("===> loop begin")

	for i := 0; i < 32*1000*1000; i++ {
		containers = append(containers, i)
		if i == 16*1000*1000 {
			readMemStatus()
		}
	}
	log.Println("===> loop end")
}

func TestMemStats(t *testing.T) {
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

func TestPPROF(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()
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
