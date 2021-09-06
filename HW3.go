//Даются две строки X и S.
//Нужно найти и вывести первое вхождение подстроки S в строке X.
//Если подстроки S нет в строке X - вывести -1
package main

import (
	"fmt"

	"strings"
)

func main() {
	x, s := "", ""
	_, err := fmt.Scan(&x)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&s)
	if err != nil {
		return
	}

	res := strings.Index(x, s)
	if res >= 0 {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}

}
