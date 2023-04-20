package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go printNum(i)
	}

	time.Sleep(time.Second)
}

func printNum(j int) {
	fmt.Println(j)
}