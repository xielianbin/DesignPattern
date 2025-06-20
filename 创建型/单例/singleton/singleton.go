package main

import (
	"fmt"
	"sync"
)

// 设置一个全局的独占锁
var lock = &sync.Mutex{}

type singleton struct {
	// 私有化构造函数
	Name string
	id   int
}

func (s *singleton) SetName(name string) {
	s.Name = name
}
func (s *singleton) GetId() int {
	return s.id
}

var singletonInstance *singleton

// GetInstance 返回单例实例
func GetInstance() *singleton {
	if singletonInstance == nil {
		// 使用互斥锁来保证线程安全
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating new singleton instance")
		singletonInstance = &singleton{Name: "Singleton Instance 1", id: 1}
		return singletonInstance
	} else {
		fmt.Println("Returning existing singleton instance")
		return singletonInstance
	}
}
