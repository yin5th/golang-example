package main

import(
    "fmt"
)

func romanToInt(s string) int {
    m := make(map[string]int)
    m["I"] = 1
    m["V"] = 5
    m["X"] = 10
    m["L"] = 50
    m["C"] = 100
    m["D"] = 500
    m["M"] = 1000
    
    m["IV"] = 4
    m["IX"] = 9
    m["XL"] = 40
    m["XC"] = 90
    m["CD"] = 400
    m["CM"] = 900
    
    length := len([]rune(s))
    //fmt.Println(length)
    var total int
    for i:=0;i<length;i++ {
        now := string(s[i]) 
        if (i != length -1) {
            next := string(s[i+1])
            two := now + next 
            //fmt.Printf("%T\n", s[i])
            //fmt.Printf("two:%s now:%s next:%s\n", two,now, next)
            if v, ok := m[two]; ok {
                total += v
                total -= m[next]
                continue
            }
        }

        total += m[now]
    }
    return total
}

func main() {
    s := romanToInt("CDXCV")
    fmt.Println(s)
}
