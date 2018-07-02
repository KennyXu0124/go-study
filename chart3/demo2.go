package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 1001
}
func main() {
	arr := [...]int{0, 2, 4, 6, 8, 10, 12}
	for i, v := range arr[2:5]  {
		fmt.Println(i, v)
	}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	arrFinal := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arrFinal[2:6] //2, 3, 4, 5
	s4 := s3[3:5]

	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d", s4, len(s4), cap(s4))
}
