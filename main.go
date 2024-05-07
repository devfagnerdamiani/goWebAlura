// AULA 6 - Exibindo dados do banco

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
	Id         int
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

	db := conectaComOBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {

			panic(err.Error())

		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	temp.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()

}

func conectaComOBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_lojas password=Kurosaki@80 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {

		panic(err.Error())
	}

	return db
}
