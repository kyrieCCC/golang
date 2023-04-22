package main

import "sync"

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func (j int)  {
			defer wg.Done()
			println(j)
		}(i)
	}

	wg.Wait()
}