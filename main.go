package main

import (
	"fmt"
	"popcount/popcount"
)

func main() {
	fmt.Println(popcount.PopCountLoop(2))
	fmt.Println(popcount.PopCountLoop(255))
	fmt.Println(popcount.PopCountLoop(1023))
}
