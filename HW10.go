//Блокировка, которая приостановит выполнение функции, пока не завершит
//выполнение вся группа горутин.
package main

import (
	"fmt"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Начали")
	time.Sleep(time.Second * 2)
	fmt.Println("Закончили")
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(wg)
	}

	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}
