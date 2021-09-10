//Если время события до 13-00, то ничего менять не нужно
//Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
//а затем вывести на стандартный вывод в том же формате.

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	layout := "2006-01-02 15:04:05"
	cur := input
	t, err := time.Parse(layout, cur)
	if err != nil {
		fmt.Println(err)
	}
	if t.Hour() < 13 {
		fmt.Println(input)
	} else {
		e := t.Add(time.Hour * 24)
		fmt.Println(e.Format("2006-01-02 15:04:05"))
	}

}
