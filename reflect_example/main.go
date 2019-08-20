package main

import (
    "fmt"
    "reflect"
)

func main() {
    type X int

    var a X = 15
    t := reflect.TypeOf(a)
    v := reflect.ValueOf(a)

    fmt.Printf("name:%s value:%v kind:%s\n", t.Name(), v, t.Kind())
}
