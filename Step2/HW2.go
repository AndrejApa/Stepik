/*Функция на вход получает целое положительное число (uint), возвращает число того же типа в
котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0,
то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(number uint) uint {
		var res string

		for number != 0 {
			if digit := number % 10; digit%2 == 0 && digit != 0 {
				res = strconv.Itoa(int(digit)) + res
			}
			number /= 10
		}

		if result, err := strconv.Atoi(res); err == nil && result != 0 {
			return uint(result)
		}
		return 100
	}
	fmt.Println(fn(222222))
	fmt.Println(fn(2343437))
	fmt.Println(fn(0))
}
