package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 3
const monitorar = 3

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			monitorandSite()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}
func exibeIntroducao() {
	versao := 1.1
	fmt.Println("Olá usuario seja bem vindo, versão atual é", versao)
	fmt.Println(" ")
}

func exibeMenu() {
	fmt.Println("------MENU------")
	fmt.Println("1 - INICIAR MONITORAMENTO")
	fmt.Println("2 - EXIBIR LOGS")
	fmt.Println("0 - SAIR DO PROGRAMA")

}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Você Escolheu a opção:", comandoLido)
	return comandoLido
}
func monitorandSite() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "http://www.random-status-code.herokuapp.com/", "https://www.caelum.com.br"}

	for i := 0; i < monitorar; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
}
func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "Foi carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "Esta com problemas!")
	}
}
