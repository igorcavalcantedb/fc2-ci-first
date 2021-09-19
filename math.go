package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 13))
	fmt.Println(pot(3))

}
func Soma(a int, b int) int {
	return a + b

}

func pot(a int) int {
	return a * a
}
