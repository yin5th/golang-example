package main

import (
	"bufio"
	"fmt"
	"os"
)

func testOs() {
	var buffer [512]byte

	n, err := os.Stdin.Read(buffer[:])
	if err != nil {

		fmt.Println("read error:", err)
		return

	}

	fmt.Println("count:", n, ", msg:", string(buffer[:]))
}

func testBufio() {
	reader := bufio.NewReader(os.Stdin)

	result, err := reader.ReadString('\n')
	if err != nil {

		fmt.Println("read error:", err)
	}

	fmt.Println("result:", result)
}

func main() {
	//os 输入输出
	//testOs()

	//使用bufio 输入输出
	testBufio()
}
