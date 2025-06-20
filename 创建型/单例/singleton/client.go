package main

import "fmt"

func Client(number int) {
	singleton := GetInstance()
	singleton.SetName(fmt.Sprintf("Singleton Instance %d", number))
	fmt.Println(singleton.Name)
	fmt.Println(singleton.GetId())

}
