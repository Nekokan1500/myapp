package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	add = func(i, j int) (result int, err error) { return i + j, nil }
	sub = func(i, j int) (result int, err error) { return i - j, nil }
	mul = func(i, j int) (result int, err error) { return i * j, nil }
	div = func(i, j int) (result int, err error) {
		if j == 0 {
			err = errors.New("The denominator cannot be zero!")
			return -999, err
		}
		return i / j, nil
	}
)

var opMap = map[string]func(int, int) (result int, err error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// Exercise 2
func fileLen(fileName string) (length int, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var size int
	buffer := make([]byte, 2048)
	for {
		count, err := f.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			size += count
			break
		}
		size += count
	}
	return size, err
}

// Exercise 3
func prefixer(prefix string) func(string) string {
	return func(input string) string {
		return prefix + input
	}
}

func main() {
	// Exercise 1 and test
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}

	fmt.Println("==========================")

	// Exercise 2 test
	fileSize, err := fileLen("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File size is", fileSize, "bytes.")

	fmt.Println("==========================")

	// Exercise 3 test
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))

}
