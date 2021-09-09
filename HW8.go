//Необходимо вычислить результат выполнения функции work для каждого из полученных чисел
package main

import (
	"fmt"
	"time"
)

func work(x int) int {
	time.Sleep(time.Second)
	if x >= 4 {
		return x + 1
	} else {
		return x - 1
	}
}
func main() {
	m := map[int]int{}
	for i := 0; i < 10; i++ {
		var n int
		fmt.Scan(&n)
		if _, ok := m[n]; !ok {
			m[n] = work(n)
		}
		fmt.Print(m[n], " ")
	}

}
