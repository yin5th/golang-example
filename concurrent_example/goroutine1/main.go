package main

import (
    "fmt"
    "time"
)

func main() {
    go fmt.Println("goroutine")
    fmt.Println("main")
    time.Sleep(time.Second)
}
