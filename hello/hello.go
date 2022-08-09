package main

import "fmt"

const englishHelloPrefix = "Hello, "
const FrenchHelloPrefix = "Bonjour, "
const SpanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = SpanishHelloPrefix
	case "French":
		prefix = FrenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	a := F()
	a[0]() //0xc00004c080 0
	a[1]() //0xc00004c088 1
	a[2]() //0xc00004c090 2
	a[0]()
}
func F() []func() {
	b := make([]func(), 3)
	for i := 0; i < 3; i++ {
		j := i
		b[i] = func() {
			fmt.Println(&j, j)
			j++
			fmt.Println(&j, j)
		}
	}
	return b
}
