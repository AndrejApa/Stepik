//Анализ результатов спортивных матчей и подсчет очков заданной команды.
package main

import "fmt"

func main() {

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var s string
	sum := 0
	fmt.Scan(&s)
	results = append(results, s)
	for _, el := range results {
		switch {
		case el == "w":
			sum += 3
		case el == "d":
			sum += 1
		}
	}
	fmt.Print(sum)
}
