# Explicação sobre Ponteiros em Go

Um ponteiro nada mais é do que uma variável que, ao invés de armazenar um valor (como `true` ou `"hello world"`), ela armazena um **endereço de memória** onde um valor está alocado.

Neste exemplo, criamos duas funções em Go para explicar o conceito de ponteiros. A situação envolve um cliente de um banco que deseja fazer um PIX. Antes de confirmar a transação, é realizada uma simulação do PIX para mostrar ao usuário o valor que restará em sua conta. Dependendo da decisão do usuário, o PIX é confirmado ou não.

---

## Funções Criadas

### 1. Simulação do PIX (Passagem por Valor)
A função `visualizaSaldoRestantePix` realiza uma simulação do PIX. Ela recebe uma **cópia** do saldo e do valor do PIX, calcula o saldo restante e exibe o resultado **sem alterar** o saldo original.

### 2. Confirmação do PIX (Passagem por Referência)
A função `confirmaPix` realiza a confirmação do PIX. Ela recebe um **ponteiro** para o saldo e o valor do PIX. Dessa forma, ela **altera diretamente** o valor do saldo original.

---

## Definindo Ponteiros

Quando se usa um ponteiro para uma variável, há alguns elementos de sintaxe importantes:

1. **Operador `&`**:
   - Quando você coloca um `&` na frente de uma variável, está obtendo o **endereço de memória** (ponteiro) daquela variável.

2. **Operador `*`**:
   - O `*` é usado para **desreferenciar** um ponteiro, ou seja, acessar o valor armazenado no endereço de memória apontado pelo ponteiro.
   - Também é usado para declarar uma variável como um ponteiro.

### Exemplo de Função com Ponteiro

```go
func confirmaPix(saldo *float64, valorPix float64) {
    err := saldoSuficiente(*saldo, valorPix)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        *saldo -= valorPix
        fmt.Printf("PIX realizado com sucesso! Saldo atual: R$ %.2f\n", *saldo)
    }
}