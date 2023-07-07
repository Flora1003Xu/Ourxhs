package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Let's sleep!")
	go func() {
		time.Sleep(10 * time.Second)
	}()
}
