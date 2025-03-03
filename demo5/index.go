package main

import "fmt"

func main() {
	fn1()
	fn2("good morning")
	fn3("Good ", 1)
	const a, b = 2, 3
	fmt.Printf("%d + %d = %d\n", a, b, sum(a, b))

	var x, y int = swap(1, 0)
	fmt.Printf("%d,%d => %d,%d\n", x, y, x, y)
	c, d := swap(a, b)
	fmt.Printf("%d,%d => %d,%d\n", a, b, c, d)

	_c, _d ,title := swapV2(10, 20)
	fmt.Printf("%d,%d => %d,%d,%s\n", 10, 20, _c, _d,title)
}

func fn1() {
	fmt.Println("456")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Print(title)
	fmt.Println(version)
}

func sum(a, b int) int {
	return a + b
}
func sum2(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}
func swapV2(a, b int) (x, y int,title string) {
	y = a
	x = b
	title = "resulth"
	return
}
