package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var n int
	fmt.Println("Enter the number of elements in the array:")
	fmt.Scan(&n)

	arr := make([]float64, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Float64s(arr)

	part1 := arr[:len(arr)/4]
	fmt.Println(part1)
	part2 := arr[len(arr)/4 : 2*len(arr)/4]
	part3 := arr[2*len(arr)/4 : 3*len(arr)/4]
	part4 := arr[3*len(arr)/4:]

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		sort.Float64s(part1)
		wg.Done()
	}()
	go func() {
		sort.Float64s(part2)
		wg.Done()
	}()
	go func() {
		sort.Float64s(part3)
		wg.Done()
	}()
	go func() {
		sort.Float64s(part4)
		wg.Done()
	}()

	wg.Wait()

	sortedArr := make([]float64, 0, len(arr))
	sortedArr = append(sortedArr, part1...)
	sortedArr = append(sortedArr, part2...)
	sortedArr = append(sortedArr, part3...)
	sortedArr = append(sortedArr, part4...)

	fmt.Println(sortedArr)
}
