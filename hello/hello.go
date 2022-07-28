package main

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
}
