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

	prefix := EnglishHelloPrefix

	switch language {
	case spanish:
		prefix = SpanishHelloPrefix
	case french:
		prefix = FrenchHelloPrefix
	}

	return prefix + name
}
