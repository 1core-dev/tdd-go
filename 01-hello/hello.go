package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("English", "world"))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return SpanishHelloPrefix + name
	}

	if language == french {
		return FrenchHelloPrefix + name
	}

	return EnglishHelloPrefix + name
}
