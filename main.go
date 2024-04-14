package main

import "fmt"

const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "

func Hello(name, language string) string {
	res := englishHelloPrefix
	if name == "" {
		res += "World"
	} else {
		res += name
	}
	if language == portuguese {
		return portugueseHelloPrefix + name
	}

	return res
}

func main() {
	fmt.Println(Hello("", ""))
}
