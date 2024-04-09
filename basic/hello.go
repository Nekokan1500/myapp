package main

import "fmt"

func main() {
	fmt.Printf("Hello, %s!\n", "world")
	// Ex. 2-1
	var i = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
	// Ex. 2-2
	const value = 5
	var i2 = value
	var f2 = value
	fmt.Println(i2, f2)
	// Ex. 2-3
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807
	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)

}
