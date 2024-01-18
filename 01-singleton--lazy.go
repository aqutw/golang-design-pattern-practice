package main

import "fmt"

type Singleton struct {
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		fmt.Println("debug: nil")
		instance = &Singleton{}
	}
	fmt.Println("debug: return")
	return instance
}

func main() {
	fmt.Println("Lazy:")
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
}
