package main

import "fmt"

func add(a, b int) int { return a + b }

func swap(a, b string) (string, string) { return b, a }

func divmod(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func main() {
	fmt.Println("add:", add(3, 4))
	x, y := swap("a", "b")
	fmt.Println("swap:", x, y)
	q, r := divmod(10, 3)
	fmt.Println("divmod:", q, r)
}
