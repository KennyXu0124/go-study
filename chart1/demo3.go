package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}

	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println(contents)
	} else {
		fmt.Println(err)
	}
 }
