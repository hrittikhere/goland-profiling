package main

import "testing"

func TestMutx(t *testing.T) {
	result := Mutx(4)
	t.Logf("%d", result)
}

func TestBlock(t *testing.T) {
	Block()
}

func TestFibRecursive(t *testing.T) {

	for i := 0; i < 100; i++ { // running it a 1000 times
		if FibRecursive(31) != 1346269 {
			//The above lines test the FibRecursive function to see if it returns the correct result when given the input 31. If it does not return 1346269, it will throw an error.
			t.Error("Incorrect!")
		}
	}
}

func TestFib(t *testing.T) {
	for i := 0; i < 1000000; i++ { // running it a 1000 times
		if Fib(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}
