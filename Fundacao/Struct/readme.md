# Structs em Go

## Sobre o Código
Este código demonstra o uso de **structs** em Go para representar funcionários de uma empresa, armazenando seu nome, cargo e salário. Além disso, há uma função que permite aplicar um aumento salarial a um funcionário específico.

## Estrutura do Código

### 1. Definição da Struct `Funcionario`
A struct `Funcionario` contém três campos:
- `nome` (string): Nome do funcionário.
- `cargo` (string): Cargo ocupado pelo funcionário.
- `salario` (float64): Salário do funcionário.

```go
 type Funcionario struct {
 	nome    string
 	cargo   string
 	salario float64
 }
```

### 2. Função `aumentoSal`
A função `aumentoSal` recebe um ponteiro para um `Funcionario` e um percentual de aumento, aplicando o reajuste ao salário do funcionário.

```go
func aumentoSal(f *Funcionario, percentual float64) {
	aumento := f.salario * (percentual / 100)
	f.salario += aumento
	fmt.Println()
	fmt.Printf("Aumentando o salário de %s em %.0f%%\n", f.nome, percentual)
	fmt.Printf("Salário atual do %s é R$%.2f\n", f.nome, f.salario)
}
```

### 3. Função `main`
- Cria uma lista de funcionários.
- Exibe os dados dos funcionários.
- Aplica um aumento salarial de 10% ao primeiro funcionário da lista.

```go
func main() {
	funcionarios := []Funcionario{
		{"Gabriel", "Analista", 3520},
		{"Maria", "Enfermeira", 5000},
	}

	for _, f := range funcionarios {
		fmt.Printf("Nome: %s | Cargo: %s | Salário: %.2f\n", f.nome, f.cargo, f.salario)
	}

	aumentoSal(&funcionarios[0], 10)
}
```

## Saída Esperada
```
Nome: Gabriel | Cargo: Analista | Salário: 3520.00
Nome: Maria | Cargo: Enfermeira | Salário: 5000.00

Aumentando o salário de Gabriel em 10%
Salário atual do Gabriel é R$3872.00
```

## Conceitos Abordados
- Uso de **structs** para representar objetos do mundo real.
- Manipulação de **slices de structs**.
- Passagem de **ponteiros** para modificar valores dentro de uma função.
- Formatação de saída com `fmt.Printf`.

## Como Executar
1. Instale o Go caso ainda não tenha instalado.
2. Salve o código em um arquivo `main.go`.
3. No terminal, navegue até o diretório onde o arquivo está salvo e execute:
   ```sh
   go run main.go
   ```

