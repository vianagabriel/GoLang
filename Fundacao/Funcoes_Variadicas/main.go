package main

import "fmt"

// Função variádica que soma vários números
func soma(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

func main() {
    fmt.Println(soma(1, 2, 3))       // Saída: 6
    fmt.Println(soma(10, 20, 30, 40)) // Saída: 100
}
