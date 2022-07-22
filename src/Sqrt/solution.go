package main

import (
    "fmt"
)

func Sqrt(x int) int {
    const times int = 100
    var z float64 = 1.0

    for i := 0; i < times; i++ {
        z -= (z*z - float64(x))/(2*z)
    }
    return int(z)
}

func main() {
    fmt.Println(Sqrt(8))
    fmt.Println(Sqrt(4))
    fmt.Println(Sqrt(8192))
}
