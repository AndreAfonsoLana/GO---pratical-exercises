package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func callAPIexternal(ctx context.Context, nameDOTest string) {
	fmt.Printf("%s Start call external API...\n",nameDOTest)
	
		ctxTimeout, cancel := context.WithTimeout(ctx, 6*time.Second)
		defer cancel()

		url := "http://localhost:8081/lerda"

		request, error_ := http.NewRequestWithContext(ctxTimeout, http.MethodGet, url, nil)

		if error_ != nil {
			return
		}

		response, erro := http.DefaultClient.Do(request)

		if erro != nil {
			fmt.Printf("[%s] Falha na chamada (Cancelada ou Erro): %v\n", nameDOTest, erro)
			return
		}
		defer response.Body.Close()
}

func main() {
	fmt.Printf("Start main APP.")

	http.HandleFunc("/certo", func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()
		callAPIexternal(ctx, "CERTO")
		fmt.Printf("Finalizado.")
	})

	http.HandleFunc("/errado", func(w http.ResponseWriter, r *http.Request){
		ctx := context.Background()
		//context.Background() sempre gera uma "raiz" nova e completamente vazia. Ele não tem valores embutidos (como IDs de tracing), não tem prazo de validade (timeout) e não escuta sinais de cancelamento. É como a semente de uma nova árvore.
		callAPIexternal(ctx, "ERRADO")
		fmt.Printf("Finalizado.")
	})

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}