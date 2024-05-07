// AULA 8 - Modularizando um pouco mais

// go mod init // para iniciar o uso de modulo
// godoc.org // site de pesquisa de pacotes do GO
// comando para baixar o pacote de conector do postgreesql

package main

import (
	"goWebAlura/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil) // Liberar porta no servidor local
}
