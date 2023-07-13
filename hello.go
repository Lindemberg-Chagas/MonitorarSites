package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5
const monitorar = 2

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
			imprimeLogs()
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
	versao := 2.0
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
	sites := lerArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}
	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "Foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "Esta com problemas!")
		registraLog(site, false)
	}
}

func lerArquivo() []string {

	var site []string

	arquivo, err := os.Open("site.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		site = append(site, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return site
}
func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + "-" + site + "- ESTAVA: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
