package main

import "fmt"

type Figure struct {
	sideA int
	sideB int
}

func findP(receiver Figure) (P int) {
	P = 2 * (receiver.sideA + receiver.sideB)
	return P
}
func findS(receiver Figure) (S int) {
	S = receiver.sideA * receiver.sideB
	return S
}

func main() {
	var Triangle Figure = Figure{
		sideA: 2,
		sideB: 3,
	}
	fmt.Println(Triangle)
	fmt.Println(findP(Triangle))
	fmt.Println(findS(Triangle))
}
