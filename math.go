package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 13))
	fmt.Println(Square(3))

}
func Soma(a int, b int) int {
	return a + b

}

func Square(a int) int {
	return a * a
}
