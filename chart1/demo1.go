package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	aa = 3
	bb = "kkk"
	cc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c ,s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "ghi"
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 +4i
	fmt.Println(cmplx.Abs(c))
}

func perfectEuler() {
	fmt.Printf("%3f\n", cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func constExample() {
	const (
		a, b = 3, 4
		filename = "test.txt"
	)

	c := int(math.Sqrt(a * a + b * b))
	fmt.Println(c, filename)
}

func enums() {
	const (
		cpp = iota
		_
		python
		java
		php
		golang
	)

	fmt.Println(cpp, python, java, php, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	euler()
	perfectEuler()
	triangle()
	constExample()
	enums()
}
