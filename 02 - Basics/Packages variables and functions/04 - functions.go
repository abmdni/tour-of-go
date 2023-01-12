package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// we can write func add (x,y int) int{ return x+y}

func main() {
	fmt.Println(add(1, 2))
}
