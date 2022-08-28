package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/hallex-abreu/Go-crie-de-uma-aplica-o-web/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 8)
		if err != nil {
			log.Println("Erro na conversão da preco:", preco)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", quantidade)
		}

		models.CriarNovoDadoNoBanco(nome, descricao, precoConvertido, quantidadeConvertida)

		http.Redirect(w, r, "/", 301)
	}
}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	produto_id := r.URL.Query().Get("id")
	models.DeletaProduto(produto_id)
	http.Redirect(w, r, "/", 301)
}

func EditarProduto(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "edit", nil)
}
