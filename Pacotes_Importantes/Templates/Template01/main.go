package main

import (
	"html/template"
	"os"
)

type Aluno struct {
	Nome          string
	Turma         string
	Identificador int
}

func main() {

	aluno1 := Aluno{"Gabriel", "4a", 1}

	templat := template.New("TemplateAluno")
	templat, _ = templat.Parse("Nome do aluno: {{.Nome}}, Turma: {{.Turma}}, Id: {{.Identificador}}")
	err := templat.Execute(os.Stdout, aluno1)
	if err != nil {
		panic(err)
	}

}
