package main

import (
	"testing"
	"math/rand"
	"time"
)

// run cmd in terminal for test
// "go test quick_sort_test.go quick_sort.go -v"


// test empty array
func Test_empty(t *testing.T) {
	a := []int{}
	QuickSort(a)
	for i := 0; i < len(a) - 1; i++ {
		if a[i] > a[i + 1] {
			t.Errorf("Error")
		}
	}
}

// test array of the same value
func Test_equal(t *testing.T) {
	a := make([]int, 10000)
	QuickSort(a)
	for i := 0; i < len(a) - 1; i++ {
		if a[i] > a[i + 1] {
			t.Errorf("Error")
		}
	}
}

// test random for 100 times
func Test_random(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := [10000]int {0}
	for i:= 0; i < 100; i++ {
		r = rand.New(rand.NewSource(int64(r.Int())))
		for j, _ := range a {
			a[j] = r.Int()
		}
		QuickSort(a[:])
		for j := 0; j < len(a) - 1; j++ {
			if a[i] > a[i+1] {
				t.Errorf("Error")
			}
		}
	}
}
