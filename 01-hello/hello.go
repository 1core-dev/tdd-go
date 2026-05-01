package main

import "fmt"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func main() {
	fmt.Println(Hello("English", "world"))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "spanish" {
		return SpanishHelloPrefix + name
	}
	return EnglishHelloPrefix + name
}
