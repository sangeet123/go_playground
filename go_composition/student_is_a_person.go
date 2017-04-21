package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

type student struct {
	person
	address_line_1 string
	address_line_2 string
	zip            string
	apt_no         string
}

func (p person) get_name() string {
	return p.first_name + " " + p.last_name
}

func main() {
	p := student{}
	p.first_name = "Hello"
	p.last_name = "World"
	fmt.Println(p.get_name())
}
