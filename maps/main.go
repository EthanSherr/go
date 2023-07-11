package main

import "fmt"

func main() {
	colors2 := make(map[string]string)
	colors2["blue"] = "#0000ff"

	fmt.Println("colors2 length", len(colors2))

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"poo":   "#444400",
	}
	if colors["poo"] == "" {
		fmt.Println("WHAT")
	}

	what, ok := colors["poo"]
	if ok {
		fmt.Println("what", what, ok)
	}
	for key, value := range colors {
		fmt.Println("key is", key, "value is", value)
	}

	// delete(colors, "poo")

	fmt.Println("map", colors)
}
