package main

import (
	"fmt"
	"log"
	"net/http"

	//implementa modelos baseados em dados para gerar saida html segura contra injacao de codigo
	"html/template"
)

// as funcoes chamadas pelo HandleFunc, recebem como parametro o corpo da requisicao e a requisicao
func index(w http.ResponseWriter, r *http.Request) { // o asterisco indica o apontamento ao local da memoria onde estao os dados da requisicao
	fmt.Fprintf(w, `
		<h1>Hello World!</h1>
		<a href="/segundaPagina">Segunda pagina</a>
	`)
}

func secondRoute(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	htmlContent := "<h2>Segunda Pagina!</h2><a href='/terceiraPagina'>ir para terceira pagina</a>"

	htmlBytes := []byte(htmlContent)

	w.Write(htmlBytes)
}

func thirdRoute(w http.ResponseWriter, r *http.Request) {
	// servindo um arquivo especifico
	// serve file recebe como parametros a resposta a requisicao e o caminho para o arquivo
	http.ServeFile(w, r, "page3.html")
}

func fourthRoute(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("page4.html")

	if err != nil {
		fmt.Println("ocorreu um erro na analise do template ou arquivo nao encontrado")

		fmt.Fprintf(w, `
		<h1>Pagina nao encontrada</h1>
		<a href="./">ir para pagina inicial</a>	
	`)

		return
	}

	data := map[string]string{
		"Title": "Quarta forma",
		"Link":  "Voltar para pagina inicial",
	}

	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func main() {

	//adiciona manipuladores ao servidor, recebendo como parametros uma rota e uma funcao, que sera executada quando ocorrer o acesso ao endereco
	http.HandleFunc("/", index)
	http.HandleFunc("/segundaPagina", secondRoute)
	http.HandleFunc("/terceiraPagina", thirdRoute)
	http.HandleFunc("/quartaPagina", fourthRoute)

	log.Println("Working server")
	http.ListenAndServe(":8080", nil)
}
