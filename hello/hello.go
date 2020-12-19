package main

import "fmt"

const vietnamese = "vietnamese"
const french = "french"
const englishHelloPrefix = "hello"
const vietnameseHelloPrefix = "xin chao"
const frenchHelloPrefix = "Bonjour"

func main() {
	fmt.Println(Hello("Nam", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + " " + name
}

func greetingPrefix(language string) (prefix string){
	switch language {
	case vietnamese:
		prefix =  vietnameseHelloPrefix;
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix =  englishHelloPrefix
	}
	return
}
