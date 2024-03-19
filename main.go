package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ { // Последнее значение i это 5, при котором for завершится
		go func(i int) {
			defer wg.Done() // Замыкание, wg прокидывается по указателю
			fmt.Println(i)  // Здесь i создвнная в функции, не та i которая есть в for
		}(i)
	}

	wg.Wait()
}
