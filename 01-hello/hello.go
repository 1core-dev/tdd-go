package main

import "fmt"

const EnglishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return EnglishHelloPrefix + name
}
