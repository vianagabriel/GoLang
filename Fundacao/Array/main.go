package main

import "fmt"

func main() {
	//1. Declarar um array com tamanho fixo e valores padrão
	var arr [2]int // Um array de 5 posições com valores padrão

	arr[0] = 1
	arr[1] = 2

	//2. Declarar e inicializar um array
	var array = [3]int{1, 2, 3} // Um array de 3 posições inicializado com valores

	//3. Inferir o tamanho do array
	array1 := [...]int{1, 2, 3, 4} // O compilador determina o tamanho automaticamente

	fmt.Println(arr)
	fmt.Println(array)
	fmt.Println(array1)

}
