package main

import (
	"fmt"
	"testing"
)

// Rectangle using value receiver
type rect1 struct {
	width, height int
}

func (r rect1) area() int {
	area := r.width * r.height
	r.width++
	r.height++

	// fmt.Printf("Inside rect1.area(); %p\n", &r)
	// fmt.Println(r)
	return area
}

func TestRect1(t *testing.T) {
	width := 10
	height := 17
	r1 := rect1{width: width, height: height}
	area1 := r1.area()

	r2 := &rect1{width: width, height: height}
	area2 := r2.area()

	fmt.Println(r1, area1)
	fmt.Println(r2, area2)
}

// Rectangle using pointer receiver
type rect2 struct {
	width, height int
}

func (r *rect2) area() int {
	// Side effect
	area := r.width * r.height
	r.width++
	r.height++

	return area
}

func TestRect2(t *testing.T) {
	width := 10
	height := 20
	r1 := rect2{width: width, height: height}
	area1 := r1.area()

	r2 := &rect2{width: width, height: height}
	area2 := r2.area()

	fmt.Println(r1, area1)
	fmt.Println(r2, area2)
}
