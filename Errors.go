package main

import (
	"fmt"
)

func main() {

	var input, inputy int
	fmt.Scan(&input)
	fmt.Scan(&inputy)
	d, err := divide(input, inputy)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(d)
	}
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("ошибка")
	}
	return a / b, nil
}
