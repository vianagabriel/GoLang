# Interfaces em Go

## Sobre o Código
Este código demonstra o uso de **interfaces** em Go para calcular a área e o perímetro de diferentes formas geométricas. A interface `geometria` define os métodos `area()` e `perim()`, que devem ser implementados por qualquer estrutura que deseje ser considerada uma "geometria".

## Estrutura do Código

### 1. Definição da Interface `geometria`
A interface `geometria` exige que qualquer estrutura que a implemente tenha os métodos:
- `area() float64`
- `perim() float64`

```go
 type geometria interface {
 	area() float64
 	perim() float64
 }
```

### 2. Implementação das Estruturas `quadrado` e `circulo`
Duas estruturas implementam a interface `geometria`:

- `quadrado`: Representa um retângulo com `largura` e `altura`.
- `circulo`: Representa um círculo com um `raio`.

```go
type quadrado struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}
```

### 3. Implementação dos Métodos
Cada estrutura implementa os métodos exigidos pela interface:

#### Quadrado
```go
func (q quadrado) area() float64 {
	return q.largura * q.altura
}

func (q quadrado) perim() float64 {
	return q.largura*2 + q.altura*2
}
```

#### Círculo
```go
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perim() float64 {
	return 2 * math.Pi * c.raio
}
```

### 4. Função `medir`
A função `medir` recebe qualquer tipo que implemente a interface `geometria`, chamando seus métodos `area()` e `perim()`.

```go
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro:", g.perim())
}
```

### 5. Função `main`
Criamos um `quadrado` e um `circulo`, passando-os para `medir()`.

```go
func main() {
	q := quadrado{largura: 3, altura: 4}
	c := circulo{raio: 5}

	medir(q)
	medir(c)
}
```

## Saída Esperada
```
{3 4}
Área: 12
Perímetro: 14
{5}
Área: 78.53981633974483
Perímetro: 31.41592653589793
```

## Conceitos Abordados
- Uso de **interfaces** para definir comportamentos comuns.
- Implementação de **métodos** em structs.
- **Polimorfismo implícito**, já que qualquer struct que implementa os métodos exigidos automaticamente satisfaz a interface.

## Como Executar
1. Instale o Go caso ainda não tenha instalado.
2. Salve o código em um arquivo `main.go`.
3. No terminal, navegue até o diretório onde o arquivo está salvo e execute:
   ```sh
   go run main.go
   ```

