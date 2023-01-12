package main

import "fmt"

func swap(x, y string) (string, string) { //multiple return
	return y, x
}
func main() {
	a, b := swap("hello", "world")
	fmt.Println(swap("a", "b"))
	fmt.Println(a, b)

}
