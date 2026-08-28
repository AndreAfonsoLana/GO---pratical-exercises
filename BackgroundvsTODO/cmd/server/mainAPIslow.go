package main

import (
	"net/http"
	"fmt"
	"time"
)

func main(){
	fmt.Printf(" Start main APIslow ")

	http.HandleFunc("/lerda", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("⏳ Requisição recebida! Simulando processamento de 5 segundos...")
		ctx := r.Context()

		select {
		case <- time.After(5 * time.Second):
			fmt.Println("✅ Processamento concluído com sucesso!")
			fmt.Fprintln(w, "Desculpe a demora! Aqui está sua resposta após 5 segundos.")
		
		case <-ctx.Done():
			// Cenário 2: O cliente cancelou/fechou a conexão antes dos 5 segundos
			fmt.Println("🛑 O cliente abortou a requisição! Economizando recursos e parando por aqui.")
		}

	})

	
	fmt.Println("🚀 Servidor da API Lenta rodando na porta 8081...")

	http.ListenAndServe(":8081", nil)
}