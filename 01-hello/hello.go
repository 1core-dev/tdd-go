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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = SpanishHelloPrefix
	case french:
		prefix = FrenchHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}

	return
}
