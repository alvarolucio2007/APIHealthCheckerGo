package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EscanearVerificar() {
	fmt.Println("Digite as URLs separadas por vírgula...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()
	urls := strings.Split(entrada, ",")
	canal := make(chan ChecagemAPI)
	contagemValidas := 0
	for _, u := range urls {
		urlLimpa := strings.TrimSpace(u)
		if ValidarURL(urlLimpa) {
			contagemValidas++
			go ChecarAPI(urlLimpa, canal)
		} else {
			fmt.Printf("Url inválida pulada %s\n", urlLimpa)
		}
	}
	for range contagemValidas {
		res := <-canal
		cor := "\033[32m" // verde
		if !res.IsOnline {
			cor = "\033[31m" // vermelho
		}
		reset := "\033[0m"
		fmt.Printf("%s[%d] %-30s | %v%s\n", cor, res.StatusCode, res.LinkAPI, res.Latency, reset)
	}
}

func ValidarURL(u string) bool {
	if !strings.HasPrefix(u, "https://") && !strings.HasPrefix(u, "http://") {
		return false
	}
	extensoes := []string{".com", ".net", ".gov", ".org", ".io", ".br"}
	valido := false
	for _, ext := range extensoes {
		if strings.Contains(u, ext) {
			valido = true
			break
		}
	}
	return valido
}
