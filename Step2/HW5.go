//Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое
//число подается на стандартный ввод с новой строки (количество чисел не известно).
//Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	s, v := bufio.NewScanner(os.Stdin), 0
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		v += n
	}
	io.WriteString(os.Stdout, strconv.Itoa(v))
}
