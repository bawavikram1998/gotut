package main

import "fmt"

type person struct {
	name string

	age int
}

func initPerson() *person {
	m := person{name: "noname", age: 50}
	return &m
}
func main() {
	fmt.Printf("gvgvhg %p\n", initPerson())

}
