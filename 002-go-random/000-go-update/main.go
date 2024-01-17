package main

import "fmt"

func main() {
	const count = 10
	for i := range count{
		fmt.Println(i)
	}
}

/*
output:

before the new update, 1.22
cannot range over count (untyped int constant 10)
*/