package main

import (
    "fmt"
)

type cat struct {
    name string
    age int
}

func (n cat) sayHello () {
    fmt.Println("Hello mio ~, i am ", n.name)
}

func demo() {
    newCat := cat{name: "Kitty", age: 3}
    newCat.age = 4
    fmt.Println(newCat)
    newCat.sayHello()
}

func main() {
    demo()
}
