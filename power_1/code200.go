package main

import (
	"fmt"
)

type WeekDay int
type Catch int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
)

func main() {
	fmt.Println(Monday, Tuesday, Wednesday, Thursday)
	//c := Catch(100)
	//fmt.Println(c*Monday, c*Tuesday, c*Wednesday, c*Thursday)
}
