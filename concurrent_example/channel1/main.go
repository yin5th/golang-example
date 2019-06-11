package main

import (
    "fmt"
)

func main() {
    naturals := make(chan int)
    squares := make(chan int)

    go func() {
        for i := 0; i < 50;i++ {
            naturals <- i
        }
        close(naturals)
    } ()

    go func() {
        for s := range naturals {
            squares <- s*s
        }
        close(squares)
    } ()

    for x := range squares {
        fmt.Println(x)
    }
}
