package main

import "fmt"

const eng = "english"
const spanish = "spanish"

const engHelloPrefix = "Hello, "
const esHelloPrefix = "Hola, "

func main() {

	fmt.Println(hello("world", eng))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Printf("sum of %d and %d is %d ", 5, 6, add(5, 6))
}

func hello(name string, language string) string {
	if name == "" {
		return engHelloPrefix + "World"
	}
	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = esHelloPrefix
	case eng:
		prefix = engHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
