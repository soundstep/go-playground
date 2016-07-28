/*
To run the benchmark: go test -run=XXX -bench=.
*/

package intpkg //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_TheAddIntsFunction(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		Add2Ints(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	b.StopTimer() //stop the performance timer temporarily while doing initialization

	//do any time consuming initialization functions here ...
	//database connection, reading files, network connection, etc.

	b.StartTimer() //restart timer
	for i := 0; i < b.N; i++ {
		Add2Ints(4, 5)
	}
}
