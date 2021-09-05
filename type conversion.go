package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	var ints []int
	strs := strings.Split(a, "")
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}
	sort.Ints(ints)
	s := ints[len(ints)-1]

	fmt.Println(s)
}
