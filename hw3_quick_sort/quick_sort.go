package main

import "fmt"

func partition(array []int) int {
	end, i := len(array), -1
	for x, j := array[end-1], 0; j < end-1; j++ {
		if array[j] <= x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[end-1] = array[end-1], array[i+1]
	return i + 1
}


// input is a slice of type []int
func QuickSort(array []int) {
	if begin, end := 0, len(array); begin < end - 1  {
		q := partition(array)
		QuickSort(array[:q])
		QuickSort(array[q+1:])
	}
}

func main() {
	a := [...]int{1, 3, 1, 2, 835, 3, 87, 35, -35, 33}
	fmt.Println(a)
	QuickSort(a[:])
	fmt.Println(a)
}
