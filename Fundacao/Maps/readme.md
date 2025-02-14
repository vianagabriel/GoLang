# Maps em Go

## Introdução
Maps (ou dicionários) são estruturas de dados que associam chaves a valores. Em Go, eles são implementados como tipos embutidos e oferecem alta performance para busca, inserção e remoção de elementos.

## Criando um Map
Os maps em Go são declarados usando a sintaxe `map[KeyType]ValueType`. Eles podem ser inicializados de diferentes maneiras:

### Declaração Simples
```go
var DDDs map[string]int
```
Neste caso, `DDDs` é um map, mas está nil e precisa ser inicializado antes do uso.

### Inicialização com `make`
```go
cep := make(map[string]int)
```
Agora `cep` está pronto para ser usado.

### Inicialização com Valores
```go
DDDs := map[string]int{
    "São Paulo":      11,
    "Rio de Janeiro": 21,
    "Belo Horizonte": 31,
    "Curitiba":       41,
    "Porto Alegre":   51,
}
```
Isso cria e inicializa o map com valores.

## Operações com Maps

### Inserindo e Atualizando Valores
```go
DDDs["Paraná"] = 91
cep["São Miguel"] = 113999049
```

### Obtendo Valores
```go
ddd := DDDs["Curitiba"]
fmt.Println("DDD de Curitiba:", ddd)
```
Se a chave não existir, o map retorna o valor zero do tipo do valor.



### Removendo um Elemento
```go
delete(cep, "Boituva")
```

### Iterando sobre um Map
```go
for cidade, ddd := range DDDs {
    fmt.Printf("Estado de %s, ddd %d\n", cidade, ddd)
}

for _, valor := range cep {
    fmt.Println("Cep:", valor)
}
```

## Considerações
- Maps são passados por referência em Go.
- A ordem de iteração sobre um map não é garantida.
- Maps não são seguros para concorrência. Para uso concorrente, utilize `sync.Map`.

## Exemplo Completo
```go
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

    for _, valor := range cep {
        fmt.Println("Cep:", valor)
    }
}
```

Isso imprime os valores armazenados nos maps.

## Conclusão
Maps são uma ferramenta poderosa para armazenar e manipular dados associativos em Go. Eles são eficientes, fáceis de usar e essenciais para muitas aplicações práticas.

