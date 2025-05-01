package main

import (
	"fmt"
)

const (
	helloPrefix         = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "Konnichiwa, "

	spanish  = "spanish"
	french   = "french"
	japanese = "japanese"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
