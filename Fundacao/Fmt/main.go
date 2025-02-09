package main

import "fmt"

// O pacote fmt é muito utilizado para formatar os dados de entrada e saída, é muito utilizado para debugs, formatação, interação entre usuários e stdout.

// Para exibir apenas o valor no terminal afim de debug, pode se utilizar o ---Print---, existem algumas variáções.
// Print   -> É o mais simples, apenas imrpime no stdout o valor que for informado
// Println -> Tem a mesma funcionalidae do Print, mas com quebra de linha, também pode se passar alguns parâmetros.
// Printf  -> Usa se o Printf quando se quer passar variávies
// tipo %s String, %v Tipo generico de dado, %d tipo inteiro, %f Float, %.2f tipo Float com 2 causas decimais
func main() {
	Saudacao := "Hello Word!"
	texto := "em go!"
	soma := 10 + 10

	fmt.Print(Saudacao)
	fmt.Println("")

	fmt.Println(Saudacao, texto)
	fmt.Printf("O resultado de 10 + 10 é %v", soma)

}
