package main

import "fmt"

type Singleton struct {
}

var instance *Singleton = &Singleton{}

func GetInstance() *Singleton {
	fmt.Println("debug: return")
	return instance
}

func main() {
	fmt.Println("Eager:")
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
}
