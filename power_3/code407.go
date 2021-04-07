package main

// interfaces

import "fmt"

// start OMIT
type plain struct{}

func (plain) Print(s string) {
	fmt.Println(s)
}

func (plain) Sing() {}

type decorated struct {
	Name string
}

func (d decorated) Print(s string) {
	fmt.Println(d.Name, "says ***", s, "***")
}

// gen OMIT
type myInterface interface {
	Print(string)
}

func hi(m myInterface) {
	m.Print("hello")
}

// plain OMIT
func main() {
	xx := plain{}
	zz := decorated{"Ringo"}
	hi(xx)
	hi(zz)
}

// end OMIT
