package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello", math.Pi)
	// every exported name start with a capital(convention) or it
	// won't be accessible from outside the package
}
