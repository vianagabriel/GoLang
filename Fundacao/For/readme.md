# Loops `for` em Go

O `for` é a única estrutura de repetição disponível na linguagem Go. Ele pode ser utilizado de diferentes formas para atender diversas necessidades.

## Estruturas do `for` em Go

### 1. Iteração sobre uma lista com `range`

O `for` pode ser utilizado com `range` para iterar sobre elementos de um slice, array ou map. No exemplo abaixo, a variável `k` representa o índice e `v` o valor de cada elemento da slice `lista`:

**OBS:** Pode-se omitir o valor da variável `k` passando o `_` no seu lugar, e para omitir o valor e exibir apenas o índice, basta informar somente `k`.

```go
lista := []int{10, 20, 30, 40, 50}
for k, v := range lista {
    fmt.Printf("Index: %v Valor: %v\n", k, v)
}
```

### 2. Loop com condição

O `for` pode funcionar como um `while` de outras linguagens, onde a repetição ocorre enquanto uma condição for verdadeira:

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
```

Neste caso, o loop imprimirá os números `1, 2 e 3` e então sairá do laço quando `i` for maior que `3`.

### 3. Loop com inicialização, condição e incremento

Esse é o formato mais comum do `for`, semelhante ao `for` do C, Java e outras linguagens:

```go
for j := 7; j <= 9; j++ {
    fmt.Println(j)
}
```

Aqui, o loop começa com `j = 7`, e continua enquanto `j <= 9`, incrementando `j` em cada iteração.

### 4. Loop infinito

Se não fornecermos uma condição para o `for`, ele se tornará um loop infinito. Podemos usar `break` para sair do loop:

```go
for {
    fmt.Println("loop")
    break
}
```

Esse código imprime `loop` uma vez e sai imediatamente, pois há um `break` dentro do bloco.

## Conclusão

O `for` em Go é versátil e pode ser usado de várias maneiras para percorrer listas, repetir ações baseadas em condições ou criar loops infinitos. Com essas variações, conseguimos implementar diferentes tipos de repetições de forma clara e eficiente.

