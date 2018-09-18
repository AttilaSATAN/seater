package main

import (
	"github.com/attilasatan/seater"
	"github.com/attilasatan/seater/plane"
)

func main() {
	planes := []plane.PlanePrototype{
		{"THY800", 60, 10, []int{3, 7}},
		{"AVA240", 120, 12, []int{4, 8}}}

	seater.Init(planes...)
}
