package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Initial data"}
	})
	return instance
}

func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 == instance2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Instances are different.")
	}

	instance1.data = "Updated data"
	fmt.Println("Data in instance1:", instance1.data)
	fmt.Println("Data in instance2:", instance2.data)
}
