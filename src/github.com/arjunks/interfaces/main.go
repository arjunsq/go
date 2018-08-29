package main

import (
	"fmt"
)

func main() {
	var two twoincrementer
	two = 0
	tw := &two
	for i := 0; i < 10; i++ {
		fmt.Println(tw.increment())
	}

	var one oneincrementer
	one = 0
	on := &one
	for i := 0; i < 10; i++ {
		fmt.Println(on.increment())
	}
}

type incrementer interface {
	increment() int
}

type twoincrementer int
type oneincrementer int

func (num *twoincrementer) increment() int {
	*num = *num + 2
	return int(*num)
}

func (num *oneincrementer) increment() int {
	*num = *num + 1
	return int(*num)
}
