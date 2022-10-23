package main

import "testing"

func testDataRaceCondition(l *testing.T) {
	var state int32

	// go x  takes state manipulate
	// go y  same thing 
		// both modify state
	
}