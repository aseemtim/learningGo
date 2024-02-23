package main

import "fmt"

// const englishHelloPrefix = "Hello, "
// const spanishHelloprefix = "Hola, "
// const frenchHelloprefix = "Bonjour, "

const (
	french = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}