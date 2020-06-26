package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		panic(fmt.Sprintf("redis dial err: %s", err))
	}
	fmt.Println("redis dial success")
	defer conn.Close()

	_, err = conn.Do("Set", "key", "hello")
	if err != nil {
		panic(fmt.Sprintf("redis set command error: %s", err))
	}
	result, err := redis.String(conn.Do("Get", "key"))
	fmt.Println("redis get command success , result: " + result)
}
