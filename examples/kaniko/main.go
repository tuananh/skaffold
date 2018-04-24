package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello kaniko!")
		time.Sleep(time.Second * 1)
	}
}
