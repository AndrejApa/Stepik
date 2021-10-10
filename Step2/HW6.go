//Аббревиатура
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := r.ReadString('\n')
	words := strings.Split(text, " ")
	res := ""
	for _, word := range words {
		res = res + string([]rune(word)[0])
	}
	fmt.Println(res)
}
