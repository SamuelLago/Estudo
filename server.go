package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler { // Middleware para logar as requisições
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path) // 1- Mostra o método (GET) // 2- Mostra o caminho usado (/hello)
		next.ServeHTTP(w, r)                               // Chama o próximo manipulador (handler) na cadeia
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Página não encontrada", http.StatusNotFound) // http.Error = Função que facilita o envio de mensagem de erro para o usuário
}

func Meuheadler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!")) // Envia uma resposta simples "Olá mundo!" para o cliente
}

func main() {
	r := mux.NewRouter() // Cria um novo roteador para gerenciar as rotas

	r.Use(LoggingMiddleware) // Adiciona o Middleware globalmente no roteador

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler) // Define um handler personalizado para rotas não encontradas

	r.HandleFunc("/hello", Meuheadler).Methods("GET") // Registra uma rota para "/hello" que aceita apenas o método GET

	//ATENCAO /HELLO VIR ANTES PARA SERVIR ARQUIVOS ESTATICOS !!!

	fs := http.FileServer(http.Dir("../../Front-gerencir/public")) // FileServer = Serve arquivos estáticos a partir da pasta especificada
	r.PathPrefix("/").Handler(fs)                                  // PathPrefix = Serve todos os arquivos estáticos com o prefixo "/"

	fmt.Println("Servidor rodando na porta 5555...")
	http.ListenAndServe(":5555", r) // Inicia o servidor HTTP na porta 5555 usando o roteador definido
}
