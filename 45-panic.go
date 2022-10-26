package main

import "os"

func main() {

	a := 1 / 0
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
