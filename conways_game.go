package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make(Universe, height)

	for i := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}

func (u Universe) Show() {
	for i := range u {
		for j := range u[i] {
			if u[i][j] {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Println("")
	}
}

func (u Universe) Seed() {
	for i := 0; i < (width*height)/4; i++ {
		rand_h := rand.Intn(height)
		rand_w := rand.Intn(width)
		u[rand_h][rand_w] = true
	}
}

func (u Universe) Alive(x, y int) bool {
	y = (y + width) % width
	x = (x + height) % height

	return u[x][y]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0

	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if (x != i || y != j) && u.Alive(i, j) {
				count++
			}

		}
	}

	return count
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	if u.Alive(x, y) {
		return n == 3 || n == 2
	}
	return n == 3
}

func Step(a, b Universe) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(i, j)
		}
	}
}

func main() {
	universe, alternate_universe := NewUniverse(), NewUniverse()
	universe.Seed()

	for i := 0; i < 300; i++ {
		Step(universe, alternate_universe)
		universe.Show()
		clearScreen()
		time.Sleep(time.Second / 30)
		universe, alternate_universe = alternate_universe, universe
	}
}
