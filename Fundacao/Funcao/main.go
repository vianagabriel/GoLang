package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(dobro(10))

	resultado, err := dividir(50, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func dobro(num int) int {
	return num * 2
}

func dividir(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("não é possivel dividir por zero")

	}

	return num1 / num2, nil
}
