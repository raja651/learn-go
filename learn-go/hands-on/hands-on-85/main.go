package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	//var mu sync.Mutex
	var incrementor int64
	gs := 100
	wg.Add(gs)

	for range gs {
		go func() {
			//mu.Lock()
			// v := incrementor
			// v++
			// incrementor = v
			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			//fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())
			//mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value", incrementor)
}
