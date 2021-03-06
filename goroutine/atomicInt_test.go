package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func TestAtomicInt(* testing.T) {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	go func() {
		fmt.Println(a.get())
	}()

	for runtime.NumGoroutine() != 1 {
	}
}
