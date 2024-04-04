package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5.0, 6.0
	fmt.Println("Comparing numbers of nongeneric int types", getMaxNumInt(a, b))
	fmt.Println("Comparing numbers of nongeneric float64 types", getMaxNumFloat64(c, d))

	// type determined by compiler
	fmt.Println("Comparing numbers of generic int types", getMaxNum(a, b))
	// explicitly specify type
	fmt.Println("Comparing numbers of generic float64 types", getMaxNum[float64](c, d))
}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CustomNumberT interface {
	// line 1 and 2 are disjunctive
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// a derived type of int64
type MyInt64 int64

// a type alias for int32
type MyInt32 = int32

func CustomNumberTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = 5.0, 6.0
	fmt.Println("Comparing numbers of custom generic types", getMaxCusNum(a, b))
	fmt.Println("Comparing numbers of custom generic types with alias", getMaxCusNum(a1, b1))

	var c, d float64 = 5.0, 6.0
	fmt.Println("Comparing numbers of custom generic types", getMaxCusNum(c, d))
	var e, f int64 = 7, 8
	var g, h MyInt64 = 7, 8
	fmt.Println("Comparing numbers of generic types", getMaxCusNum(e, f))
	fmt.Println("Comparing numbers of custom generic types", getMaxCusNum(g, h))
}

func getMaxCusNum[T CustomNumberT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func BuiltInCase() {
	var a, b string = "abc", "efg"
	fmt.Println("Built-in comprable generic types", getBuiltInComparable(a, b))
	var c, d float64 = 100, 100
	fmt.Println("Built-in comprable generic types", getBuiltInComparable(c, d))
	var f = 100.023
	printBuiltInAny(f)
	printBuiltInAny(a)
}

func getBuiltInComparable[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func printBuiltInAny[T any](a T) {
	fmt.Println("Built-in any generic type", a)
}
