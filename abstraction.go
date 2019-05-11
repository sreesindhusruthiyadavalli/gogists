package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}
func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (a Bird) Eat() {
	fmt.Println("worms")
}
func (a Bird) Move() {
	fmt.Println("fly")
}
func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (a Snake) Eat() {
	fmt.Println("mice")
}
func (a Snake) Move() {
	fmt.Println("slither")
}
func (a Snake) Speak() {
	fmt.Println("hsss")
}

func callMethod(a Animal, method string) {
	fmt.Println("callmethod %+v, method %s", a, method)
	if method == "eat" {
		a.Eat()
	} else if method == "speak" {
		a.Speak()
	} else if method == "move" {
		a.Move()
	}
}

func main() {
	//reading user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	for {
		fmt.Println("-> ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		animals := strings.Split(text, " ")
		fmt.Println(animals[0], animals[1])
		if animals[0] == "cow" {
			a := Cow{}
			callMethod(a, strings.TrimSpace(animals[1]))
		} else if animals[0] == "bird" {
			a := Bird{}
			callMethod(a, strings.TrimSpace(animals[1]))
		} else if animals[0] == "snake" {
			a := Snake{}
			callMethod(a, strings.TrimSpace(animals[1]))
		}
	}

	//f1 := GenDisplaceFn(a, vel, disp)
	//fmt.Println(f1(time))

}
