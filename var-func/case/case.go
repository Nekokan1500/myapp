package _case

import "errors"

type User struct {
	Name string
	Age  uint
}

func NewUser(name string, age uint) *User {
	return &User{Name: name, Age: age}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() uint {
	return u.Age
}

// passing parameters by value
func SumCase(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("Two inputs cannot be both zero.")
		return
	}
	sum = a + b
	return
}

// passing parameters by reference
func ReferenceCase(a int, b *int) {
	a += 1
	*b += 1
}

// private global variable (to the current package)
var g int

// public global variable
var G int

// variable scope
func ScopeCase(a, b int) {
	c := 100
	g = a + b + c
	G = g
}
