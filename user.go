package main

import "fmt"

type User struct {
}

// Run 跑路
func (u User) Run() {
	fmt.Println("run run run")
	fmt.Println("跑路")
}
