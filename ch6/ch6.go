package main

import "fmt"

// Exercise 1
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// Exercise 2
func UpdateSlice(slice []string, input string) {
	length := len(slice)
	slice[length-1] = input
	fmt.Println(slice)
}

func GrowSlice(slice []string, input string) {
	slice = append(slice, input)
	fmt.Println(slice)
}

// Exercise 3
func BuildPersons() {
	//persons := make([]Person, 10000000)
	var persons []Person
	for i := 0; i < 10000000; i++ {
		persons = append(persons, Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       20,
		})
	}
}

func main() {
	// Test of exercise 1
	fmt.Println(MakePerson("John", "Smith", 36))
	fmt.Println(MakePersonPointer("John", "Smith", 36))

	fmt.Println("================================================================")

	// Test of exercise 2
	slice := []string{"Apple", "Banana"}
	fmt.Println(slice)
	GrowSlice(slice, "Hello")
	fmt.Println(slice)
	UpdateSlice(slice, "Goodbye")
	fmt.Println(slice)

	// Test of exercise 3
	BuildPersons()
}
