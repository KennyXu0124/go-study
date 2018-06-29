package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score > 100 || score < 0:
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score >=0 && score < 60:
			g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}

	return g
}
func main() {
	fmt.Println(
		grade(1),
		grade(59),
		grade(60),
		grade(99),
		grade(-3),
	)
}
