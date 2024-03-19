package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ { // Последнее значение 5
		go func() {
			defer wg.Done() // Замыкание, wg прокидывается по указателю
			fmt.Println(i)  // i прокидывается по указателю
		}()
	}

	wg.Wait()
}
