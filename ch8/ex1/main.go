package main

import "fmt"

type IntFloat interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr |
		~float32 | ~float64
}

func Double[T IntFloat](v T) T {
	return v * 2
}

func main() {
	fmt.Println(Double(10))
}
