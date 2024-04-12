package main

import (
	"fmt"
	//depend "myapp/ch7/dependency"
	//"net/http"
	ex "myapp/ch7/exercises"
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
	/*
		l := depend.LoggerAdapter(depend.LogOutput)
		ds := depend.NewSimpleDataStore()
		logic := depend.NewSimpleLogic(l, ds)
		c := depend.NewController(l, logic)
		http.HandleFunc("/hello", c.SayHello)
		http.ListenAndServe(":8080", nil)
	*/
	league := ex.League{
		Teams: []string{"Manchester United", "Chelsea", "Arsenal"},
		Wins: map[string]int{
			"Manchester United": 0,
			"Chelsea":           0,
			"Arsenal":           0,
		},
	}

}
