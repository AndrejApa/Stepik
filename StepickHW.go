package main
import "fmt"
func main() {
	var size int
	fmt.Scanln(&size)
	if(size < 4) {
		return
	}

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&elements[i])
	}
	fmt.Println(elements[3])
}