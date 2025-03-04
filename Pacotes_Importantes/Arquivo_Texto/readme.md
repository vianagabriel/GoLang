# Manipulação de Arquivos em Go

## Sobre o Código
Este programa em Go demonstra como criar, escrever e ler um arquivo de texto. Ele faz uso de bibliotecas padrão da linguagem para manipulação de arquivos e entrada/saída de dados.

## Funcionalidades
1. **Criação de um arquivo** chamado `Arquivo.txt`.
2. **Escrita de texto** no arquivo.
3. **Leitura do arquivo** e exibição do conteúdo no console.
4. **Leitura parcial do arquivo** utilizando um buffer.

## Explicação do Código

### Importação de Pacotes
```go
import (
    "bufio"
    "fmt"
    "os"
)
```
- `os`: Pacote utilizado para manipulação de arquivos e operações do sistema operacional.
- `fmt`: Pacote usado para formatação e exibição de mensagens no console.
- `bufio`: Pacote utilizado para leitura eficiente de arquivos.

### Criação e Escrita no Arquivo
```go
f, err := os.Create("Arquivo.txt")
if err != nil {
    panic(err)
}
```
- `os.Create()`: Cria um novo arquivo. Se o arquivo já existir, ele será sobrescrito.
- `panic(err)`: Interrompe a execução caso ocorra um erro.

```go
tam, err := f.Write([]byte("Gravando texto"))
if err != nil {
    panic(err)
}
fmt.Printf("Arquivo criado com sucesso, tamanho de %v bytes\n", tam)
```
- `f.Write([]byte(...))`: Escreve os bytes no arquivo.
- `fmt.Printf()`: Exibe uma mensagem no console informando o tamanho do arquivo criado.

### Leitura Completa do Arquivo
```go
file, err := os.ReadFile("Arquivo.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(file))
```
- `os.ReadFile()`: Lê todo o conteúdo do arquivo e retorna os dados em um slice de bytes.
- `string(file)`: Converte os bytes lidos para string e imprime no console.

### Leitura Parcial com Buffer
```go
arq, err := os.Open("Arquivo.txt")
if err != nil {
    panic(err)
}
reader := bufio.NewReader(arq)
buffer := make([]byte, 10)
```
- `os.Open()`: Abre o arquivo para leitura.
- `bufio.NewReader()`: Cria um leitor bufferizado para leitura eficiente.
- `make([]byte, 10)`: Cria um buffer de 10 bytes para leitura parcial.

```go
for {
    n, err := reader.Read(buffer)
    if err != nil {
        break
    }
    fmt.Println(string(buffer[:n]))
}
```
- `reader.Read(buffer)`: Lê até 10 bytes do arquivo.
- `string(buffer[:n])`: Converte os bytes lidos para string.
- O loop continua lendo até que não haja mais dados no arquivo.

## Como Executar o Código
1. Instale o Go em seu sistema (https://go.dev/dl/).
2. Salve o código em um arquivo `main.go`.
3. No terminal, navegue até a pasta onde o arquivo foi salvo.
4. Execute o comando:
   ```sh
   go run main.go
   ```
5. O programa criará o arquivo, gravará o texto e exibirá o conteúdo no console.

