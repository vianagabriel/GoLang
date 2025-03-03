package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	ContaGabriel := Conta{
		Numero: 1,
		Saldo:  100000,
	}

	// Convertendo uma Struct para um Json
	res, err := json.Marshal(ContaGabriel)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))

	// Convertendo usando encoder
	ContaRebeca := Conta{
		Numero: 2,
		Saldo:  200000,
	}

	json.NewEncoder(os.Stdout).Encode(ContaRebeca) // posso converter direto e mandar para o terminal ou salvar em uma var

	// transformando em Struct
	Contajson := []byte(`{"n":3,"s":500}`)
	var contaPoup Conta
	err = json.Unmarshal(Contajson, &contaPoup)
	if err != nil {
		println(err)
	}

	fmt.Println(contaPoup.Saldo)
}
