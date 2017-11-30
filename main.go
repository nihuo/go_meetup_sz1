package main

import (
	"fmt"
)

func getHello() string {
	return "Hello world"
}

func main() {
	fmt.Println("hello world")
	fmt.Println(getHello())
}
