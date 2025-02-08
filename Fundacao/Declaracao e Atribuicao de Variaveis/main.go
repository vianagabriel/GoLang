package main

// Tipos de dados

// Escopo global
var (
	exist        bool    //   Boleano
	num          int     //   Inteiro
	texto        string  //   texto
	numFlutuante float64 //   numero flutuante
)

func main() {
	exist = true
	num = 10
	texto = "Tipo Texto"
	numFlutuante = 10.2

	local := "Variável escopo local" // escopo local

	println(exist)
	println(num)
	println(texto)
	println(numFlutuante)
	println(local)

}
