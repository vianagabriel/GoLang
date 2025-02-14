package main

import "fmt"

func main() {
	DDDs := map[string]int{
		"São Paulo":      11,
		"Rio de Janeiro": 21,
		"Belo Horizonte": 31,
		"Curitiba":       41,
		"Porto Alegre":   51,
	}

	DDDs["Paraná"] = 91

	for cidade, ddd := range DDDs {
		fmt.Printf("Estado de %s, ddd %d\n", cidade, ddd)
	}

	cep := make(map[string]int)

	cep["São Miguel"] = 113999049
	cep["Paulista"] = 33342256
	cep["Boituva"] = 00445660

	delete(cep, "Boituva")

	for _, cep := range cep {
		fmt.Println("Cep:", cep)
	}

}
