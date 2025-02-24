package main

import (
	"errors"
	"fmt"
)

// Função auxiliar para verificar se o saldo é suficiente
func saldoSuficiente(saldo, valorPix float64) error {
	if valorPix < 0 {
		return errors.New("o valor do PIX não pode ser negativo.")
	}
	if saldo < valorPix {
		return errors.New("saldo insuficiente para realizar o PIX.")
	}
	return nil
}

// Função para visualizar o saldo restante (passagem por valor)
func visualizaSaldoRestantePix(saldo, valorPix float64) {
	err := saldoSuficiente(saldo, valorPix)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		saldoAtual := saldo - valorPix
		fmt.Printf("Simulação de PIX: Saldo restante será R$ %.2f\n", saldoAtual)
	}
}

// Função para confirmar o PIX (passagem por referência)
func confirmaPix(saldo *float64, valorPix float64) {
	err := saldoSuficiente(*saldo, valorPix)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		*saldo -= valorPix
		fmt.Printf("PIX realizado com sucesso! Saldo atual: R$ %.2f\n", *saldo)
	}
}

func main() {
	saldo := 50.00
	valorPix := 25.00

	fmt.Printf("Saldo antes que qualquer alteração: %2.f\n\n", saldo)

	// Visualiza o saldo restante sem alterar o saldo original
	visualizaSaldoRestantePix(saldo, valorPix)
	fmt.Printf("Saldo depois de simular pix:  %2.f\n\n", saldo)

	// Confirma o PIX, alterando o saldo original
	confirmaPix(&saldo, valorPix)
	fmt.Printf("Saldo depois de confirmar pix:  %2.f\n", saldo)

}
