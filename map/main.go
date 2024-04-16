package main

import "fmt"

func main() {
	// var syntax-- used if we want to know what key-pair values are added later on
	// var colors map[string]string
	// can also use built in make function
	// colors := make(map[string]string)

	// plain, literal syntax
	// all keys and values are of type string
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF000",
		"blue":  "#0000FF",
	}

	colors["white"] = "#FFFFFF"
	colors["black"] = "#000000"
	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
