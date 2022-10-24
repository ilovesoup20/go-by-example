package main

import "fmt"

type rect1 struct {
	width, height int
}

func (r rect1) area() int {
	fmt.Println("rect1 value receiver")
	return r.width * r.height
}

type rect2 struct {
	width, height int
}

func (r *rect2) area() int {
	fmt.Println("rect2 pointer receiver")
	return r.width * r.height
}

func main() {
	// Using 'rect1' struct
	r1 := rect1{width: 10, height: 50}
	r1p := &r1

	fmt.Printf("area r1:%d\n\n", r1.area())
	fmt.Printf("area r1p:%d\n\n", r1p.area())

	r2p := &rect1{width: 20, height: 30}
	fmt.Printf("area r2:%d\n\n", r2p.area())

	// Using 'rect2' struct
	r3 := rect2{width: 1, height: 2}
	r3p := &r3

	fmt.Printf("area r3:%d\n\n", r3.area())
	fmt.Printf("area r3p:%d\n\n", r3p.area())

	r4p := &rect2{width: 2, height: 3}
	fmt.Printf("area r4p:%d\n\n", r4p.area())
}
