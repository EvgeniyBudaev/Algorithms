package main

import "fmt"

func main() {
	fmt.Println(deferChangeValue().value)       // 42
	fmt.Println(deferChangeReturnValue().value) // 43
}

type returnData struct {
	value int
}

func deferChangeValue() returnData {
	d := returnData{}

	defer func() {
		d.value = 43
	}()

	d.value = 42
	return d
}

func deferChangeReturnValue() (d returnData) {
	defer func() {
		d.value = 43
	}()

	d.value = 42
	return
}
