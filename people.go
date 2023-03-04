package main

import "fmt"

type People struct {
	Name string
	Age  int
	Job  string
}

func (p *People) ShowCharacter() {
	fmt.Printf("My name is %v, I'm %v , my job is %v", p.Name, p.Age, p.Job)
}
