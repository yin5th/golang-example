package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Printf("redis connect failed, err:%v\n", err)
		return
	}

	defer conn.Close()

	_, err = conn.Do("Set", "name", "bocai")
	if err != nil {
		fmt.Printf("redis set failed, err:%v\n", err)
		return
	}

	res, err := redis.String(conn.Do("Get", "name"))

	if err != nil {
		fmt.Printf("redis get failed, err:%v\n", err)
		return
	}

	fmt.Printf("name is :%v\n", res)
}
