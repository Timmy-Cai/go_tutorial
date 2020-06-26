package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	simpleHash()
	multiHash()
}

//处理简单的hash,一个个放入redis
func simpleHash() {
	fmt.Println("simple redis hash")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(fmt.Sprintf("redis dial err: %s", err))
	}
	defer conn.Close()

	//redis hash ->  key field value
	_, err = conn.Do("HSet", "timmy", "sex", "male")
	if err != nil {
		panic(fmt.Sprintf("redis HSet command error: %s", err))
	}
	_, err = conn.Do("HSet", "timmy", "age", 10)
	if err != nil {
		panic(fmt.Sprintf("redis HSet command error: %s", err))
	}

	age, err := redis.Int(conn.Do("HGet", "timmy", "age"))
	sex, err := redis.String(conn.Do("HGet", "timmy", "sex"))
	fmt.Printf("timmy's sex is: %v ,age is: %v", sex, age)
}

//批量处理hash
func multiHash() {
	fmt.Println("multi redis hash")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(fmt.Sprintf("redis dial err: %s", err))
	}
	defer conn.Close()

	//redis hmset ->  key field value field value field value
	_, err = conn.Do("HMSet", "tom", "sex", "female", "age", 20)
	if err != nil {
		panic(fmt.Sprintf("redis HMSet command error: %s", err))
	}
	result, err := redis.Strings(conn.Do("HMGet", "tom", "sex", "age"))
	for i, v := range result {
		fmt.Printf("result[%d]=%s \n", i, v)
	}
}
