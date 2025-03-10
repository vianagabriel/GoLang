# Consulta de CEP com Go

Este é um projeto simples em Go (Golang) que cria uma API para consultar informações de endereços a partir de um CEP, utilizando a API pública do [ViaCEP](https://viacep.com.br/). O programa implementa um servidor HTTP que recebe um CEP via query string e retorna os dados do endereço em formato JSON.

## Funcionalidades

- Recebe um CEP como parâmetro na URL (ex.: ?cep=01001000).
- Consulta a API do ViaCEP e retorna informações como logradouro, bairro, cidade, estado, entre outros.
- Responde com dados no formato JSON.
- Inclui tratamento básico de erros para rotas inválidas ou CEPs não informados.

## Pré-requisitos

- [Go](https://golang.org/dl/) instalado (versão 1.16 ou superior recomendada).
- Conexão com a internet para consultar a API do ViaCEP.

## Como Executar

1. **Clone o repositório ou copie o código**:  
   git clone <URL_DO_REPOSITORIO>  
   cd <NOME_DO_DIRETORIO>

2. **Execute o programa**:  
   go run main.go  
   O servidor será iniciado na porta 8080.

3. **Teste a API**:  
   Abra um navegador ou use uma ferramenta como curl para fazer uma requisição:  
   curl "http://localhost:8080/?cep=01001000"  
   Ou acesse diretamente no navegador:  
   http://localhost:8080/?cep=01001000

## Exemplo de Resposta

Ao consultar o CEP 01001000, você receberá algo como:  
{  
  "cep": "01001-000",  
  "logradouro": "Praça da Sé",  
  "complemento": "lado ímpar",  
  "unidade": "",  
  "bairro": "Sé",  
  "localidade": "São Paulo",  
  "uf": "SP",  
  "estado": "São Paulo",  
  "regiao": "Sudeste",  
  "ibge": "3550308",  
  "gia": "1004",  
  "ddd": "11",  
  "siafi": "7107"  
}

## Estrutura do Código

- **main.go**:  
  - **Pacotes utilizados**: encoding/json, io, net/http.  
  - **Struct CEP**: Representa os dados retornados pela API ViaCEP.  
  - **Função main**: Configura o servidor HTTP na porta 8080 e define a rota raiz (/).  
  - **Função BuscaCepHandle**: Manipula as requisições HTTP, valida o CEP e retorna os dados em JSON.  
  - **Função BuscaCep**: Faz a requisição ao ViaCEP e deserializa a resposta JSON.

## Endpoints

| Método | Endpoint       | Parâmetro       | Descrição                          |  
|--------|----------------|-----------------|------------------------------------|  
| GET    | /              | cep (query)     | Consulta o endereço pelo CEP       |

### Exemplos de Requisições  
- Sucesso: GET /?cep=01001000 → Retorna os dados do CEP.  
- Erro (CEP não fornecido): GET / → Retorna status 404.  
- Erro (rota inválida): GET /outra-rota → Retorna status 404.

