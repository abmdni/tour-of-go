package main

import "fmt"

func main() {
	defer fmt.Println("world")
	// defer : defer the execution of a function until the
	// 		 : surrouding function return
	// deferred call's arguments are evaluated immediately.
	// but the function call in not executed until the surrounding function returns
	fmt.Println("hello")
}
