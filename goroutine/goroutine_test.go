package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	a := 0
	var mutex = new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			//fmt.Println("hello goroutine:" , i)
			mutex.Lock()
			defer mutex.Unlock()
			a++
			//runtime.Gosched() 交出控制权
		}(i)
	}

	for runtime.NumGoroutine() != 1 {
	}

	fmt.Println(a)

}
