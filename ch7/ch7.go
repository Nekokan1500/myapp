package main

import (
	"fmt"
	"time"
)

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

type BitField int

const (
	Field6 BitField = 1 << iota
	Field7
	Field8
	Field9
	Field10
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
	Increment()
}

func main() {
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
	fmt.Println(Field6, Field7, Field8, Field9, Field10)

	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}
	myStringer = pointerCounter // ok
	myStringer = valueCounter   // ok

	myIncrementer = pointerCounter // ok
	//myIncrementer = valueCounter // compile-time error!
}
