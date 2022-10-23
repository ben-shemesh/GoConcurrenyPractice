package main

import (
	"sync"
	"testing"
)

func testDataRaceCondition(l *testing.T) {
	var state int32
	var mu sync.RWMutex

	// go x  takes state manipulate
	// go y  same thing 
		// both modify state
		for i := 0; i < 10; i++ {
			go func (i int)  {
				// hold acces to the var of the mu
				mu.Lock()
				state += int32(i)
				// base logic
				mu.Unlock()
			}(i)
		}
}