package main
import "fmt"
func main() {

	var a, b int

	var numA1, numA2, numA3, numA4, numA5 int

	var numB1, numB2, numB3, numB4, numB5 int

	fmt.Scan(&a, &b)

	numA1 = a / 10000
	numA2 = (a / 1000) % 10
	numA3 = (a / 100) % 10
	numA4 = (a / 10) % 10
	numA5 = a % 10

	numB1 = b / 10000
	numB2 = (b / 1000) % 10
	numB3 = (b / 100) % 10
	numB4 = (b / 10) % 10
	numB5 = b % 10

	var num1 = [5]int{numA1, numA2, numA3, numA4, numA5}
	var num2  = [5]int{numB1, numB2, numB3, numB4, numB5}

	for i := 0; i != 5; i++ {
		for j := 0; j != 5; j++ {
			if num1[i] == 0{
				continue
			}else if num2[j] == 0 {
				continue
			}else if num1[i] == num2[j]  {

				fmt.Print(num1[i], " ")
			}
		}
	}
}

