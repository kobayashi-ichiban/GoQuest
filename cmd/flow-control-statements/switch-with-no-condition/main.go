package main

import (
	"fmt"
	"time"
)

func main() {
	times := time.Now()
	switch {
	case times.Hour() < 12:
		fmt.Println("Good morning!")
	case times.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
