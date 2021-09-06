package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var myRune = []rune(text)

	if unicode.IsUpper(myRune[0]) && string(myRune[len(myRune)-1]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
