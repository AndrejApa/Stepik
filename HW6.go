//Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func main() {
	urlDownload := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_files/task_sep_1/task.data"
	resp, err := http.Get(urlDownload)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	r.Comma = ';'
	lines, _ := r.ReadAll()

	for index, val := range lines[0] {
		if val == "0" {
			fmt.Println("Result:", index+1)
		}
	}
}
