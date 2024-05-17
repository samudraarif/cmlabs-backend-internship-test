package main

import (
    "fmt"
)

func main() {
    for i := 1; i <= 100; i++ {
        var output string
        switch {
        case i % 3 == 0 && i % 5 == 0:
            output = "FizzBuzz"
        case i % 3 == 0:
            output = "Fizz"
        case i % 5 == 0:
            output = "Buzz"
        default:
            output = fmt.Sprint(i)
        }
        fmt.Println(output)
    }
}
