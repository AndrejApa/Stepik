package main

import "fmt"

func f(ch chan int) {
	ch <- 4
}
func main() {
	ch := make(chan int)
	go f(ch)
	fmt.Println(<-ch / 2)
}
