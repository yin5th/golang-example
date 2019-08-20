package main

import(
    "fmt"
)
type In int

func (i In) add(num In) {
    i += num
}

func (i *In) addP(num In) {
    *i += num
}

func main() {
    var s In = 10
    s.add(15)
    println(s)
    s.addP(20)
    println(s)

    var a = [3]int{1,2,3}
    var b = &a

    b[1]++
    fmt.Println(a,*b)
}
