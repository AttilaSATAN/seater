package main

import (
	"github.com/attilasatan/seater"
	"github.com/attilasatan/seater/plane"
)

func main() {
	planes := []plane.PlanePrototype{
		{"THY800", 60, 10, []int{3, 7}}}

	seater.Init(planes...)
}
