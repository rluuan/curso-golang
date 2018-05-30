package main

import (
	"fmt"
)

func a() {
	fmt.Println("start")
}
func b(a, b string) {
	fmt.Println(a, b)
}
func c(a string, b bool) {
	fmt.Println(a, b)

}
func d() int {
	return 1
}
func e() (int, float64) {
	return 1, 2
}
func main() {
	a()
	b("rodrigo", "luan")
	c("rodrigo", true)
	a := d()
	c, d := e()

	fmt.Println(a, c, d)

}
