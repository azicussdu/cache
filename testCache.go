package main

import (
	"fmt"
	"github.com/azicussdu/cache/task2"
	"time"
)

func main() {
	cache := task2.New()

	err := cache.Set("userId", 42, time.Second*1)
	if err != nil {
		return
	}

	err = cache.Set("userId2", 46, time.Second*3)
	if err != nil {
		return
	}

	userId, err := cache.Get("userId")
	if err == nil {
		fmt.Println(userId)
	}

	time.Sleep(time.Second * 2)

	userId, err = cache.Get("userId")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(userId)
	}

	userId, err = cache.Get("userId2")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(userId)
	}
}

//func main() {
//cache := cache.New()
//
//cache.Set("userId", 42)
//userId, err := cache.Get("userId")
//
//if err == nil {
//	fmt.Println(userId)
//}
//
//cache.Delete("userId")
//userId, err = cache.Get("userId")
//if err != nil {
//	fmt.Println(err.Error())
//}
//
//fmt.Println(userId)
//}
