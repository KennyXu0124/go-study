package main

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
)

func convertToBin(num int) string {
	result := ""
	for ; num > 0; num = num / 2 {
		v := num % 2
		result =  strconv.Itoa(v) + result
	}

	return result
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("Love")
	}
}

func main() {
	println(
		convertToBin(12),
		convertToBin(78214),
	)

	readFile("abc.txt")

	//forever()
}
