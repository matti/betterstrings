package main

import (
	"fmt"

	"github.com/matti/betterstrings"
)

func main() {
	a, b := betterstrings.Split("a:b:c", ":").Two()
	fmt.Println(a, b)

	a, b, c := betterstrings.Split("a:b:c", ":").Three()
	fmt.Println(a, b, c)
}
