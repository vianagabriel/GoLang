package main

import "fmt"

func main() {

	lista := []int{10, 20, 30, 40, 50}
	for k, v := range lista {
		fmt.Printf("Index: %v Valor: %v\n", k, v)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
