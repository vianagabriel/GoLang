# Arrays em Go


## Sobre
Este repositório contém exemplos e explicações sobre diferentes maneiras de declarar e utilizar arrays na linguagem Go.

## Tipos de Declaração

### 1. Declarar um array com tamanho fixo e valores padrão

```go
var arr [2]int // Um array de 2 posições com valores padrão (0 para inteiros)

arr[0] = 1
arr[1] = 2
```

O array `arr` tem duas posições e armazena inteiros. Os valores padrão de um array são os valores zero do tipo correspondente.

### 2. Declarar e inicializar um array

```go
var array = [3]int{1, 2, 3} // Um array de 3 posições inicializado com valores específicos
```

O array `array` já é inicializado com os valores `{1, 2, 3}`.

### 3. Inferir o tamanho do array

```go
array1 := [...]int{1, 2, 3, 4} // O compilador determina automaticamente o tamanho do array
```

O operador `[...]` permite que o compilador conte o número de elementos e defina o tamanho do array automaticamente.

## Exemplo Completo

```go
package main

import "fmt"

func main() {
    var arr [2]int
    arr[0] = 1
    arr[1] = 2

    var array = [3]int{1, 2, 3}

    array1 := [...]int{1, 2, 3, 4}

    fmt.Println(arr)
    fmt.Println(array)
    fmt.Println(array1)
}
```

## Observações
- Arrays em Go possuem **tamanho fixo** e não podem ser redimensionados.
- Se precisar de uma estrutura dinâmica, considere usar **slices** (`[]int`).


