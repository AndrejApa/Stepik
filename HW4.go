//конвертировать  строку в Time, а затем вывести в формате UnixDate

package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(time.UnixDate))
}
