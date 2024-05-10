package main

import "fmt"

// declare type called bot
type bot interface {
	// if you are a type in the program with getGreeting and it returns a string, you are a member of type bot
	// because of above, you can now call function where bot is the expected arg type
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// the bad function calls
	// eb.printGreeting()
	// sb.printGreeting()

	printGreeting(eb)
	printGreeting(sb)
}

// Below are two functions with shared name doing the same thing, but they take in different types
// This can cause naming and type clashes
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// Below is the shared logic expecting bot interface type
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
