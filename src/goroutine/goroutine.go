package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var a int = 0
	var mutex = new(sync.Mutex)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			//fmt.Println("hello goroutine:" , i)
			mutex.Lock()
			defer mutex.Unlock()
			a++
		}(i)
	}

	for runtime.NumGoroutine() != 1 {
	}

	fmt.Println(a)

}
