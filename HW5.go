// Нужно удалить все символы, которые встречаются более одного раза
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}
	for _, ch := range a {
		if strings.Count(a, string(ch)) == 1 {
			fmt.Print(string(ch))
		}
	}
}
