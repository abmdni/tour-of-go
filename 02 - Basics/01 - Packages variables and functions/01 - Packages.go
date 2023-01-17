package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite nubmer is", rand.Intn(10))

	// rand.Intn not math.rand.Intn cause by convention the
	// package name is the same the last one of the imported path
}
