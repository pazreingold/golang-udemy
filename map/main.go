package main

import (
	"fmt"
)

func main() {
	//Option 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	//Option 2
	//var colors map[string]string

	//Option 3
	//colors := make(map[string]string)

	colors["white"] = "#ffffff"
	delete(colors, "white")
	colors["white"] = "#ffffff"

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
