package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type StringerInt int

func (si StringerInt) String() string {
	return strconv.Itoa(int(si))
}

type StringerFloat float64

func (sf StringerFloat) String() string {
	return fmt.Sprintf("%f", sf)
}

func PrintIt[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var i StringerInt = 10
	PrintIt(i)
	var f StringerFloat = 1.5
	PrintIt(f)
}
