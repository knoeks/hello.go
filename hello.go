package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y float32 = 19.99
	var z float32 = float32(x)

	fmt.Printf("%.3f", z+y)
}
