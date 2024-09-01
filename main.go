package main

import (
	"fmt"
)

func construct2DArray(original []int, m int, n int) [][]int {
	var length_of_original_array int = len(original)
	var result_array [][]int = make([][]int, m)
	// fmt.Println("if return:", float32(length_of_original_array) / float32(m))
	if ((float32(length_of_original_array) / float32(m)) != float32(n)) {
		return make([][]int, 0)
	}
	var f int = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result_array[i] = append(result_array[i], original[f])
			f++
		}
	}
	return result_array
}

func main() {
	var array []int = []int{1,2,3,4}
	var m int = 2
	var n int = 2
	var array_2d [][]int = construct2DArray(array, m, n)
	fmt.Println(array_2d)

	array = []int{1,2,3}
	m = 1
	n = 3
	array_2d = construct2DArray(array, m, n)
	fmt.Println(array_2d)

	array = []int{1,2}
	m = 1
	n = 1
	array_2d = construct2DArray(array, m, n)
	fmt.Println(array_2d)

	array = []int{1,3,3,5}
	m = 3
	n = 1
	array_2d = construct2DArray(array, m, n)
	fmt.Println(array_2d)

}