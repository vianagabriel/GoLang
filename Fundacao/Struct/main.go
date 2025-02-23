package main

import "fmt"

type Funcionario struct {
	nome    string
	cargo   string
	salario float64
}

func aumentoSal(f *Funcionario, percentual float64) {
	aumento := f.salario * (percentual / 100)
	f.salario += aumento
	fmt.Println()
	fmt.Printf("Aumentando o salario de %s em %.0f%%\n", f.nome, percentual)
	fmt.Printf("Salário atual do %s é R$%.2f", f.nome, f.salario)
}

func main() {

	funcionarios := []Funcionario{
		{"Gabriel", "Analista", 3520},
		{"Maria", "Enfermeira", 5000},
	}

	for _, f := range funcionarios {
		fmt.Printf("Nome: %s | Cargo: %s | Salário: %2.f\n", f.nome, f.cargo, f.salario)
	}

	aumentoSal(&funcionarios[0], 10)

}
