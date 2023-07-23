package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// test one by one
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}
	// table with test cases
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d and expected %d", total, item.n)
		}
	}
}

/*
commands
  - go test -cover
    it will shows overall coverage
  - go test --coverprofile=coverage.out
    creates coverage.out file
  - go tool cover --func=coverage.out
    resume of your code in terms of coverage
  - go tool cover --html=coverage.out
    view each line to show you what hasn't been tested yet
  - go test -cpuprofile=cpu.out
    shows a file with the metrics of performance
  - go tool pprof cpu.out
    shows the performance metrics
  - top shows slowest process
  - list Fibonacci makes zoom in that routine
*/
func TestGetMax(t *testing.T) {
	// test one by one
	x := 1
	y := 2
	max := GetMax(x, y)
	if max != 2 {
		t.Errorf("error GetMax expected 2 and got %d ", max)
	}

	// table with test cases
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 2},
		{2, 2, 2},
		{26, 25, 26},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d and expected %d", max, item.n)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d and expected %d", fib, item.n)
		}
	}
}
