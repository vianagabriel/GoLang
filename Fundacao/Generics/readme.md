# Generics em Go

## O que são Generics?

Go é uma linguagem de tipagem forte, ou seja, sempre precisamos declarar o tipo da variável, dos parâmetros de uma função e dos seus retornos. Isso pode levar à duplicação de código quando queremos que uma função opere sobre diferentes tipos de dados.

Por exemplo, ao criar uma função para somar elementos de um array, sem Generics, precisaríamos de duas versões para lidar com `int64` e `float64`:

```go
func SomaInt(array []int64) int64 {
    var soma int64
    for _, valor := range array {
        soma += valor
    }
    return soma
}

func SomaFloat(array []float64) float64 {
    var soma float64
    for _, valor := range array {
        soma += valor
    }
    return soma
}
```

## Usando Generics

Com Generics, podemos escrever apenas uma função que atenda a diferentes tipos de dados:

```go
func Soma[T int64 | float64](array []T) T {
    var soma T
    for _, valor := range array {
        soma += valor
    }
    return soma
}
```

### Exemplo de Uso

```go
package main

import "fmt"

func main() {
    array := []int64{1, 2, 3, 4, 5, 6, 7}
    fmt.Println(Soma(array))
}
```

## Criando uma Interface Genérica

Podemos evitar a repetição dos tipos de dados dentro dos colchetes utilizando uma interface:

```go
type Number interface {
    int64 | float64
}

func Soma1[T Number](array []T) T {
    var soma T
    for _, valor := range array {
        soma += valor
    }
    return soma
}
```

### Exemplo de Uso

```go
package main

import "fmt"

func main() {
    array := []float64{2.2, 3.5, 5.6}
    fmt.Println(Soma1(array))
}
```

## Conclusão

Generics permitem escrever funções reutilizáveis em Go sem precisar duplicar código para diferentes tipos. Isso melhora a legibilidade e a manutenção do código, tornando-o mais flexível e eficiente.

