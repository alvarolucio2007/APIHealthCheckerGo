package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ChecarAPI(url string, canal chan ChecagemAPI) {
	inicio := time.Now()
	res, err := http.Get(url)
	checagem := ChecagemAPI{
		LinkAPI: url,
		Latency: time.Since(inicio),
	}
	if err != nil {
		fmt.Printf("Erro na URL %s: %v \n", url, err)
		checagem.IsOnline = false
	} else {
		checagem.IsOnline = (res.StatusCode == http.StatusOK)
		checagem.StatusCode = res.StatusCode
		err := res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	canal <- checagem
}
