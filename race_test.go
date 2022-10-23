package main

import (

	"sync/atomic"
	"testing"
)

func testDataRaceCondition(l *testing.T) {
	var state int32
	// state = 1
	// s+1
	// s+1
	// var mu sync.RWMutex

	// go x  takes state manipulate
	// go y  same thing 
		// both modify state
		for i := 0; i < 10; i++ {
			go func (i int)  {
				// add state to the already existing value  
				atomic.AddInt32(&state, int32(i))

					// hold acces to the var of the mu
				// mu.Lock()
				state += int32(i)
					// base logic
				// mu.Unlock()
			}(i)
		}
}