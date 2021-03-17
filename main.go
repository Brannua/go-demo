package main

import (
    "fmt"
)

func helloworld() {
    fmt.Println("Hello world ~")
}

func demo1() {
    var a int = -1 // a := -1 ( 使用类型推断的简写形式 )
    if a > 1 && a < 3 {
        fmt.Println("a")
    } else if a < 0 {
        fmt.Println(a)
    } else {
        fmt.Println("oh my god.")
    }
}

func demo2() {
    arr := [3]int{1, 2, 3}
    for i := 0; i < 3; i ++ {
        fmt.Println(arr[i])
    }

    j := 0
    for ; j < 3 ; {
        fmt.Println(arr[j])
        j++
    }

    arr2 := []int{1, 2 ,3}
    fmt.Println(arr2)
    arr2 = append(arr2, 4)
    fmt.Println(arr2)
}

func demo3() {
    numbers := make(map[string]int)
    numbers["one"] = 1
    numbers["two"] = 2
    numbers["three"] = 3
    // fmt.Println(numbers)
    // delete(numbers, "one")
    // fmt.Println(numbers)

    for key, val := range numbers {
        fmt.Println(key, val)
    }
}

func demo4(a int, b int) (int, int) {
    return a + b, a * b
}

func main() {
    // helloworld()
    // demo1()
    // demo2()
    // demo3()
    fmt.Println(demo4(1, 2))
}
