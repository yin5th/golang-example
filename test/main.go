package main

import(
    "crypto/sha256"
    "fmt"
)

func main() {
    a := []byte("x")
    b := []byte("y")
    
    fmt.Printf("a的SHA256哈希码为：%x\n", sha256.Sum256(a))
    fmt.Printf("b的SHA256哈希码为：%x\n", sha256.Sum256(b))
    c := countDiff(a, b)
    fmt.Printf("a和b的不同数为：%d\n", c)
}

func countDiff(a, b []byte) int {
    aSha := sha256.Sum256(a)
    bSha := sha256.Sum256(b)
    fmt.Printf("函数中a的sha256为：%#v\n", aSha)
    fmt.Printf("函数中b的sha256为：%#v\n", bSha)
    fmt.Printf("a的sha256长度为：%d\n", len(aSha))
    fmt.Printf("b的sha256长度为：%d\n", len(bSha))
    count := 0
    for i := 0; i < len(aSha); i++ {
        fmt.Printf("aSha的第%d个bit值为：%#v\n", aSha[i])
        fmt.Printf("bSha的第%d个bit值为：%#v\n", bSha[i])
            if (aSha[i] != bSha[i]) {
                count++
            }
       
    }
    return count
}
