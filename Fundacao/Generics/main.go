package main

import "fmt"

type Number interface {
	int | float64
}

func Soma[T int | float64](array []T) T {
	var soma T

	for _, valor := range array {
		soma += valor
	}

	return soma
}

func Soma1[T Number](array []T) T {
	var soma T

	for _, valor := range array {
		soma += valor
	}

	return soma
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	arrayF := []float64{2.2, 3.5, 5.6}

	fmt.Println(Soma(array))
	fmt.Println(Soma(arrayF))

}
