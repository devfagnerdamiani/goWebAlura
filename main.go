// AULA 7 - Modularizando o código

// go mod init // para iniciar o uso de modulo
// godoc.org // site de pesquisa de pacotes do GO
// comando para baixar o pacote de conector do postgreesql

package main

import (
	"goWebAlura/models"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil) // Liberar porta no servidor local
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
