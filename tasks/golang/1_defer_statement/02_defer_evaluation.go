package main

import "fmt"

func main() {
	deferEvaluationAsUsual()  // 0
	deferEvaluationOnDemand() // 1
}

type data struct {
	value int
}

// deferEvaluationAsUsual откладывает вычисление значения value до конца функции.
func deferEvaluationAsUsual() {
	d := data{}
	defer fmt.Println(d.value) // 0

	d.value++
}

// deferEvaluationOnDemand откладывает вычисление значения value до конца функции.
func deferEvaluationOnDemand() {
	d := data{}
	defer func() {
		fmt.Println(d.value) // 1
	}()

	d.value++
}
