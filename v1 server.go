package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { // ResponseWriter = Retorna dados para o usuario // http.Request = solicitacoes para a API
		w.Write([]byte("Servidor funcionando!")) // Write = Retorna a situacao do servidor ao usuario // bytes = formato ideal
	})

	fmt.Println("Servidor rodando na porta 5555...")
	http.ListenAndServe(":5555", nil) // 1- porta no localhost // 2- nulo
}
