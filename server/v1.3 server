package main

import (
	"fmt"
	"net/http"
)

func Meuheadler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo!"))
}

func main() {
	fs := http.FileServer(http.Dir("../../Front-gerencir/public")) //FileServer = Serve arquivos com um manipulador HTTP // Dir = direciona o FileServer
	http.Handle("/", fs)                                           // Handle = Hosteia o FileServer // 1- Escolha de arquivo especifico // 2- O FileServer

	http.HandleFunc("/hello", Meuheadler) // 1- Porta no localhost // 2- HandleFunc chama uma funcao

	fmt.Println("Servidor rodando na porta 5555...")
	http.ListenAndServe(":5555", nil)
}
