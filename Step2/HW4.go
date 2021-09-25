package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	_, err := fmt.Scanf("%s %d", &text, &width)
	if err != nil {
		return
	}

	res := text
	if len(text) > width {
		res = text[:width] + "..."
	}

	fmt.Println(res)
}
