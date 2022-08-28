package main

import (
	"net/http"

	"github.com/hallex-abreu/Go-crie-de-uma-aplica-o-web/controllers"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.DeletarProduto)
	http.HandleFunc("/edit", controllers.EditarProduto)
	http.ListenAndServe(":8000", nil)
}
