package main

import (
    "fmt"
)

func reverse_integer(number int) int {
    new_int := 0
    for number > 0 {
        remainder := number % 10
        new_int *= 10
        new_int += remainder 
        number  /= 10
    }
    return new_int 
}

func main() {
    fmt.Println(reverse_integer(1234))
}