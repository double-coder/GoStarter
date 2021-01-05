package main

import (
	"errors"
	"fmt"
)

const (
	first = iota
	second
	third
)

func learning() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var firstName1 *string = new(string)

	*firstName1 = "Arthur"

	fmt.Println(*firstName1)

	// constants - it is

	fmt.Println(first, second, third)
	datatypes()

}
func datatypes() {
	arr1 := [3]int{1, 2, 3}

	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 4
	fmt.Println(arr, arr1)

	slice := arr[:]
	fmt.Println(slice)

	slice1 := []int{1, 4, 9}

	slice1 = append(slice, 4)

	fmt.Println(slice1)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

	m := map[string]int{"foo": 32}
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Arther"
	u.LastName = "Dent"
	fmt.Println(u)

	u2 := user{ID: 2,
		FirstName: "Aditya",
		LastName:  "Path",
	}
	fmt.Println(u2)

	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println(port)
	return 2, errors.New("something went wrong")
}
