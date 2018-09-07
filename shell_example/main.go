package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 一个简单的shell
	//1.接收命令行输入
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	//对特殊命令特殊处理 如cd
	switch args[0] {
	case "cd":
		//跳转
		if len(args) < 2 {
			err := os.Chdir("$HOME")
			if err != nil {
				return err
			}
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		return nil

	case "exit":
		os.Exit(0)

	}
	//执行可带参数的命令
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
