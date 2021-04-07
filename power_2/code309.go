package main

// interfaces

import "fmt"

type plain struct{}

func (plain) Print(s string) {
	fmt.Println(s)
}

func (plain) Sing() {

}

type decorated struct{}

func (decorated) Print(s string) {
	fmt.Println("***", s, "***")
}

type myInterface interface {
	Print(string)
}

func hi(m myInterface) {
	m.Print("hello")
}

func main() {
	var xx plain
	var zz decorated
	hi(xx)
	hi(zz)

}
