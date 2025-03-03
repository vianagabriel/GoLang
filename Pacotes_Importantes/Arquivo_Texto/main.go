package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Criando arquivo texto
	f, err := os.Create("Arquivo.txt")
	if err != nil {
		panic(err)
	}

	// escrevendo arquivo texto
	tam, err := f.Write([]byte("Gravando texto"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso, tamanho de %v bytes\n", tam)

	// Leitura
	file, err := os.ReadFile("Arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Leitura pouco a pouco

	arq, err := os.Open("Arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arq)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

}
