//chat is a server that lets clients chat with each other.
package main

import(
    "bufio"
    "net"
    "fmt"
    "log"
)

//广播
type client chan<- string

var (
    entering = make(chan client)
    leaving = make(chan client)
    messages = make(chan string) //所有传入的客户端消息
)

func broadcaster() {
    clients := make(map[client]bool) // 所有已连接的客户端
    for {
        select {
        case msg := <-messages:
            //广播传入的消息给所有的客户端
            for cli := range clients {
                cli <- msg
            }
        case cli := <-entering:
            clients[cli] = true
        case cli := <-leaving:
            delete(clients, cli)
            close(cli)
        }
    }
}

func handleConn(conn net.Conn) {
    ch := make(chan string) //传出客户端的消息
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages <- who + " has arrived"
    entering <- ch

    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ": " + input.Text()
    }

    leaving <- ch
    messages <- who + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg)
    }
}

func main() {
    fmt.Println("开始")
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }

    go broadcaster()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }

}
