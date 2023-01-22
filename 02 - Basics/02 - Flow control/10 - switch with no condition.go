package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// cleaner then if else!?
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning!! ")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("good evening.")

	}
}
