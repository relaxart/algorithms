package main

import "fmt"

func main() {
	testArray := []int{1, 23, 6, 4, 83, 7, 9, 5, 4, 0, 5}
	bubbleSort(testArray)
	fmt.Println(testArray)
}

func bubbleSort(arr []int) {
	var previous int
	unsorted := false
	for i, v := range arr {
		if previous > -1 && arr[previous] > v {
			holder := arr[previous]
			arr[previous] = v
			arr[i] = holder
			unsorted = true
		}
		previous = i
	}
	if unsorted {
		fmt.Println(arr)
		bubbleSort(arr)
	}
}
