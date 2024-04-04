package _case

import "fmt"

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID    int
	City  string
	State string
}

// collection to list or list to collection
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{ID: 1, Name: "John", Age: 20}
	userMp[2] = user{ID: 2, Name: "Jane", Age: 21}
	userList := mapToList[int64, user](userMp)

	ch := make(chan user)
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	fmt.Println("======================")

	addressMp := make(map[int64]address, 0)
	addressMp[1] = address{ID: 1, City: "New York", State: "NY"}
	addressMp[2] = address{ID: 2, City: "Los Angeles", State: "CA"}
	addressList := mapToList[int64, address](addressMp)

	ch1 := make(chan address)
	go myPrintln(ch1)
	for _, addr := range addressList {
		ch1 <- addr
	}
}
