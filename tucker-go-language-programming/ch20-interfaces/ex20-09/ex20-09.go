package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	// ...
}

func (s *Student) String() string {
	return "Student"
}

type Actor struct {
	// ...
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	stduent := stringer.(*Student)
	fmt.Println(stduent)
}

func main() {
	actor := &Actor{}
	ConvertType(actor)
}
