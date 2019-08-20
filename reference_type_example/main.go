package main

import(
    "fmt"
)

func mkslice() []int {
    s := make([]int, 0, 10)
    s = append(s, 100)
    return s
}

func main() {
    s  := mkslice()
    fmt.Printf("s= %v\n", s)
}
