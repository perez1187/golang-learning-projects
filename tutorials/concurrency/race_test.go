package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// using mutex
// disadvantages of mutex, if you have many variables
// you lock all variables and your program will be slow
// instead of mu, you can use channels to communicate
func TestDataRaceConditions(t *testing.T) {
	var state int32

	var mu sync.RWMutex // we can name it lock as well

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			// business logic
			mu.Unlock()
		}(i)
	}
}

// atomic value
// is antoher way (mu or channels)

func TestAtomicValue(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			// state += int32(i)
			//in atomic we need to specify address of space
			// &state => pointer in memory -> 8 bytes
			//( addres, what we add)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
