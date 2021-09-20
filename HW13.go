//Sample Input: 1 149,6088607594936;1 179,0666666666666
//Sample Output: 0.9750

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var numbers [2]float64
	sc.Scan()
	txt := sc.Text()

	txts := strings.Split(txt, ";")

	for index, txtItem := range txts {
		txtItem = strings.TrimSpace(txtItem)
		txtItem = strings.Replace(txtItem, " ", "", -1)
		txtItem = strings.Replace(txtItem, ",", ".", -1)
		numbers[index], _ = strconv.ParseFloat(txtItem, 64)
	}

	fmt.Printf("%.4f \n", numbers[0]/numbers[1])
}
