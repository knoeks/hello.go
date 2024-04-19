package main

import "fmt"

func main() {
	var a, b = 6, "Hello"
	const c, d int = 6, 5
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Printf("%#v\n", c)

	//declaring an array
	var arr1 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(arr1[i])
	}

}
