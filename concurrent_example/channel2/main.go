package main

ipmort(
    "fmt"
    "time"
)

    var tick = time.Tick(1 * time.Second)
func main() {
    fmt.Println("Commencing countdown.")

    //tick := time.Tick(1 * time.Second)

    for countdown := 10; countdown > 0; countdown-- {
        fmt.Println(countdown)
        <-tick
    }
}
