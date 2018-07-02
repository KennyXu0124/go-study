package main

import "fmt"

func printArr(arr [5]int) {
	arr[2] = 1873
	for i, v := range arr  {
		fmt.Println(i, v)
	}
}

func printArrPoint(arr *[5]int) {
	arr[2] = 1873
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var grid [4][5] int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArr(arr1)

	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	printArrPoint(&arr1)

	printArrPoint(&arr1)
}
