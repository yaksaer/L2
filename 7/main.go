package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			for v := range ch {
				result <- v
			}
			fmt.Println("Close")
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}
