package main

import (
	"fmt"
	_case "myapp/var-func/case"
)

func main() {
	a := 10
	b := 20
	fmt.Println(_case.SumCase(a, b))
	fmt.Println(a, b)
	_case.ReferenceCase(a, &b)
	fmt.Println(a, b)
	fmt.Println(_case.G)
	_case.ScopeCase(a, b)
	fmt.Println(_case.G)

	user := _case.NewUser("Arthur", 18)
	fmt.Println(user.GetName(), user.GetAge())

}
