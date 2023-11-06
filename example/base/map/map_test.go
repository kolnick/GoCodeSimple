package main

import (
	"fmt"
	"sync"
	"testing"
)

var syncMap = sync.Map{}

func TestSyncMap(t *testing.T) {
	// put
	syncMap.Store("aa", "a")
	// get
	load, _ := syncMap.Load("aa")
	fmt.Println(load)

}

func TestMap(t *testing.T) {

	countryMap := make(map[string]string, 10)
	// 添加元素
	countryMap["China"] = "Beijing"
	countryMap["Japan1"] = "Tokyo"
	countryMap["Japan2"] = "Tokyo"
	countryMap["Japan3"] = "Tokyo"
	countryMap["Japan4"] = "Tokyo"
	countryMap["Japan5"] = "Tokyo"
	countryMap["Japan6"] = "Tokyo"
	countryMap["Japan7"] = "Tokyo"
	countryMap["Japan8"] = "Tokyo"
	countryMap["Japan9"] = "Tokyo"
	countryMap["Japan10"] = "Tokyo"
	countryMap["Japan11"] = "Tokyo"
	countryMap["Japan12"] = "Tokyo"

	// 删除key
	delete(countryMap, "China")

	// 遍历Map
	for k, v := range countryMap {
		fmt.Printf("key %v value %v \n", k, v)
	}

	// 查找元素
	value, flag := countryMap["Japan233"]
	fmt.Println(value, flag)

	// 获取元素个数
	size := len(countryMap)
	fmt.Println(size)
	create()

}

func create() {
	// 声明方式1
	map1 := map[int]string{
		1: "caochaojie",
		2: "zhoujielun",
	}
	fmt.Println(map1)
	// 声明方式2
	map2 := make(map[string]string)
	map2["one"] = "one"
	map2["two"] = "two"
	map2["three"] = "three"

}
