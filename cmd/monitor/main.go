package main

import (
	"github.com/seu-usuario/APIHealthCheckerGo/internal" // Importa seu pacote
)

func main() {
	// 1. Conecta e Migra
	internal.ConectarBanco()

	// 2. Agora você pode usar o internal.DB para salvar dados
	novaChecagem := internal.ChecagemAPI{
		LinkAPI:  "https://google.com",
		IsOnline: true,
	}
	internal.DB.Create(&novaChecagem)
}
