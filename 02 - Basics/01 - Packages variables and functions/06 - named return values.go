package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // nacked return (used only in short functions cause hurt readability)
}
func main() {
	fmt.Println(split(17))
}
