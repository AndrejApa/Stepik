package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
func isPalindrome(str string) interface{} {
	if str == Reverse(str) {
		return true
	}
	return false
}
func main() {
	var str string
	fmt.Scan(&str)
	if isPalindrome(strings.ToUpper(str)) == true {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
