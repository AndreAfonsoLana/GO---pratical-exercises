package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheChamadas struct {
	mux      sync.Mutex
	chamadas []string
}

func (c *CacheChamadas) AdicionarSeguro(idChamada string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.chamadas = append(c.chamadas, idChamada)
}
func disparaChamada(id int, wg *sync.WaitGroup, cache *CacheChamadas) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)

	idFormatado := fmt.Sprintf("SIP/Ramal-%d", id)

	cache.AdicionarSeguro(idFormatado)
}

func main() {
	var wg sync.WaitGroup
	meuCache := CacheChamadas{}

	totalChamadas := 10000

	fmt.Println("🚀 Iniciando AGV Zakisoft Dispatcher...")

	for i := 1; i <= totalChamadas; i++ {
		wg.Add(1)
		go disparaChamada(i, &wg, &meuCache)
	}

	fmt.Println("⏳ Todas as goroutines foram disparadas. Aguardando processamento...")
	wg.Wait()

	fmt.Println("=======================================")
	fmt.Printf("🏁 Processamento finalizado! Total no Cache: %d chamadas\n", len(meuCache.chamadas))
}
