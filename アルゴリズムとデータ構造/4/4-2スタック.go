package main

import (
	"fmt"
	"strconv"
)

var top int
var S[1000] int

func main() {
	var a, b int
	var s string
	top = 0

	for {
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		
		if (s == "+") {
			a = pop()
			b = pop()
			push(b + a)
		} else if s == "-" {
			a = pop()
			b = pop()
			push(b - a)
		} else if s == "*" {
			a = pop()
			b = pop()
			push(b * a)
		} else {
			a, _ = strconv.Atoi(s)
			push(a)
		}
	}
	fmt.Printf("%d\n", pop())
}

func push(x int) {
    top = top + 1
	S[top] = x
}

func pop() int {
	top = top - 1
	return S[top + 1]
}