package main

import (
	"fmt"
)

var outside string = "outside the \"main()\""

func main() {
	var x int = 2
	var y float32 = 2.999
	var z float32 = float32(x)

	fmt.Printf("%.3f ", z+y)

	gatve := "verkių g."
	fmt.Println(gatve)
	println(outside)
}
