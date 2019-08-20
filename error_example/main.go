package main

import(
    "log"
    "runtime/debug"
    "errors"
)

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
    if y==0 {
        return 0, errDivByZero
    }

    return x/y, nil
}

func test() {
    panic("there are something wrong")
}

func main() {
    //z, err := div(5, 0)
    z, err := div(5, 2)
    if err == errDivByZero {
        log.Fatalln(err)
    }

    println(z)

    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
        }
    }()

    test()
    println("222")
}
