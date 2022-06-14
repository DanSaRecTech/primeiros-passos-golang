package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorar int = 1
const iniciarLogs int = 2
const sairDoPrograma int = 0

const monitoramentos = 3
const delayDoMonitoramento = 5

func main() {

	exibeInformacoes()
	for {
		exibeMenu()
		printaTipoDeComando(readComando())
	}
}

func exibeInformacoes() {
	nome := "Dan"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Versão do programa", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do sistema")
}

func readComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func printaTipoDeComando(comando int) {
	switch comando {
	case monitorar:
		iniciarMonitoramento()
	case iniciarLogs:
		fmt.Println("Iniciando os logs...")
	case sairDoPrograma:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando!")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for indice, site := range sites {
			fmt.Println("")
			fmt.Println("Testando site", indice, ":", site)
			verificaSiteOnline(site)
		}
		time.Sleep(delayDoMonitoramento * time.Second)
	}
	fmt.Println("")
}

func verificaSiteOnline(site string) {

	resp, err := http.Get(site)
	isErro(err)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso! Status:", resp.Status)
	} else {
		fmt.Println("Site:", site, "está com problemas!!! Staus:", resp.StatusCode)
	}
}

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	isErro(err)
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n') //até o fim da linha
		linha = strings.TrimSpace(linha)      //tira o quebra linha no fim de cada linha

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	return sites
}

func isErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro!!!", err)
	}
}
