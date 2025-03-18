# Conversão de Struct para JSON em Go

Este projeto demonstra como converter structs para JSON e vice-versa em Go, utilizando os pacotes `encoding/json` e `os`.

## Funcionalidades
- Conversão de uma struct para JSON usando `json.Marshal`.
- Impressão do JSON gerado no terminal.
- Uso de `json.NewEncoder(os.Stdout).Encode` para converter e exibir diretamente.
- Conversão de JSON para struct utilizando `json.Unmarshal`.

## Estrutura do Código

O programa define a struct `Conta`:
```go
 type Conta struct {
     Numero int `json:"n"`
     Saldo  int `json:"s"`
 }
```
A struct tem dois campos: `Numero` e `Saldo`, que são mapeados para `n` e `s` no JSON.

### Convertendo Struct para JSON
Usamos `json.Marshal` para transformar uma struct em JSON:
```go
res, err := json.Marshal(ContaGabriel)
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(res))
```

Também podemos usar `json.NewEncoder(os.Stdout).Encode()` para exibir o JSON diretamente no terminal:
```go
json.NewEncoder(os.Stdout).Encode(ContaRebeca)
```

### Convertendo JSON para Struct
Para converter um JSON em struct, usamos `json.Unmarshal`:
```go
Contajson := []byte(`{"n":3,"s":500}`)
var contaPoup Conta
err = json.Unmarshal(Contajson, &contaPoup)
if err != nil {
    println(err)
}
fmt.Println(contaPoup.Saldo)
```

## Como Executar
1. Instale o Go em seu sistema.
2. Salve o código em um arquivo `main.go`.
3. No terminal, execute:
   ```sh
   go run main.go
   ```

## Saída Esperada
```json
{"n":1,"s":100000}
{"n":2,"s":200000}
500
```
Isso demonstra a conversão de struct para JSON e de JSON para struct com sucesso.



