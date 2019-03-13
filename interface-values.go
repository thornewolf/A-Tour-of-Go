package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	var f F
	f = 1.0
	f.M()
	describe(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
