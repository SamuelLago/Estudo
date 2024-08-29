package main

import (
	"fmt"
	"net/http"
)

type MeuHandler struct{}

func (h MeuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // O * significa o ponteiro, uma referencia ao valor em vez uma copia dele, aumentando eficiencia e mutabilidade
	if r.URL.Path == "/hello" { // URL = campo na estrutura do http.Request (como host, caminho, query, etc.) // Path = Faz parte da URL representa o caminho solicitado no caso hello
		w.Write([]byte("Ola, mundo!"))
		return
	}
	http.NotFound(w, r)
}

func main() {
	handler := MeuHandler{}

	fmt.Println("Servidor rodando na porta 5555...")
	http.ListenAndServe(":5555", handler)
}
