// AULA 5 - Conectando com o banco

// go mod init // para iniciar o uso de modulo
// godoc.org // site de pesquisa de pacotes do GO
// comando para baixar o pacote de conector do postgreesql

package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComOBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil) // Liberar porta no servidor local
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3}, {"Fone", "Muito bom", 59, 2},
		{"Produto Novo", "Muito legal", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}

func conectaComOBancoDeDados() *sql.DB {
	conexao := "user=postgress dbname=alura_lojas password=Kurosaki80 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {

		panic(err.Error())
	}

	return db
}
