package main

import "fmt"

func main() {
	username := ""
	age := 0
	count, _ := fmt.Scanln(&username, &age)
	fmt.Println("账号信息为:", username, age, count)
}
