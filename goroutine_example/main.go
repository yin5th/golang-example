package main

import(
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "time"
    "os"
)

func main() {
    startTime := time.Now()
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch) //获取信息
    }

    costTime := time.Since(startTime).Seconds()
    fmt.Printf("total cost %.2f s\n", costTime)
}

func fetch(url string, ch chan<- string) {
    startTime := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("url: %s fetch err. msg: %v", url, err)
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)//将内容丢入discard 如丢入垃圾桶
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    costTime := time.Since(startTime).Seconds()
    ch <- fmt.Sprintf("url %s cost %.2fs, content length: %d", url, costTime , nbytes)
}
