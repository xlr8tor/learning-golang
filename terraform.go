package main

import "fmt"

type planets []string

func (p planets) terraform() {
	for i := range p {
		p[i] = "New "  + p[i]
	}
}
func main() {
	planets := planets{"Mercury",
	"Venus",
	"Earth",
	"Mars",
	"Jupiter",
	"Saturn",
	"Uranus",
	"Neptune",}

	planets.terraform()
	fmt.Println(planets)
}
