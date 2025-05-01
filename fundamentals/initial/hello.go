package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(s string) string {
	if s == "" {
		s = "world"
	}
	return helloPrefix + s
}

func main() {
	fmt.Println(Hello("world"))
}
