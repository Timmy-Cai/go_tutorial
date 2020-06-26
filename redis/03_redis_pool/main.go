package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//定义全局redis连接池
var pool *redis.Pool

//设置操作超时时间
var doWithTimeout = 500 * time.Millisecond

//初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     10,              //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   30,              //在给定的时间,最大的连接数，表示同时最多有N个连接。0表示不限制。
		IdleTimeout: 1 * time.Minute, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭。如果设置成0，空闲连接将不会被关闭。应该设置一个比redis服务端超时时间更短的时间。
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//获取连接
	conn1 := pool.Get()
	defer conn1.Close() //释放连接，放回连接池

	_, err := conn1.Do("Set", "key", "hello")
	if err != nil {
		panic(fmt.Sprintf("redis set command error: %s", err))
	}
	result, err := redis.String(conn1.Do("Get", "key"))
	fmt.Println("redis get command success , result: " + result)

	conn2 := pool.Get()
	defer conn2.Close() //释放连接，放回连接池
	_, err = conn2.Do("Set", "key", "hello")
	if err != nil {
		panic(fmt.Sprintf("redis set command error: %s", err))
	}
	result, err = redis.String(conn2.Do("Get", "key"))
	fmt.Println("redis get command success , result: " + result)

	//设置执行操作超时时间
	result, err = redis.String(redis.DoWithTimeout(conn1, doWithTimeout, "Get", "key"))
	if err != nil {
		panic(fmt.Sprintf("redis get command error: %s", err))
	}
	fmt.Println(result)
}
