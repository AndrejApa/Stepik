//Получаем 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64. Третье значение в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/" (определенная математическая операция).
//Но такие идеальные случаи будут не всегда, вы можете получить значения других типов, а также строка в третьем значении может не относится к одной из указанных математических операций.

package main

import (
	"fmt"
)

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/"
}

func main() {
	value1, value2, operation := readTask()
	if x, ok := value1.(float64); !ok {
		fmt.Printf("value=%v: %T", value1, value1)
	} else if y, ok := value2.(float64); !ok {
		fmt.Printf("value=%v: %T", value2, value2)
	} else if z, ok := operation.(string); !ok {
		fmt.Print("неизвестная операция")
	} else {
		if len(z) > 1 {
			fmt.Print("неизвестная операция")
		} else if z[0] != '+' && z[0] != '-' && z[0] != '*' && z[0] != '/' {
			fmt.Print("неизвестная операция")
		} else {
			switch z[0] {
			case '+':
				fmt.Printf("%.4f", x+y)
			case '-':
				fmt.Printf("%.4f", x-y)
			case '*':
				fmt.Printf("%.4f", x*y)
			case '/':
				fmt.Printf("%.4f", x/y)
			}
		}
	}
}
