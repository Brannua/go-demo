package main

import (
    "fmt"
)

func add(a * int) {
    * a = * a + 2
}

func main() {
    n := 0
    add(&n)
    fmt.Println(n)
}
