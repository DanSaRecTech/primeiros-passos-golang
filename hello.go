package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
		exibirOsLogs()
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
		registraLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas!!! Staus:", resp.StatusCode)
		registraLogs(site, false)
	}
}

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	isErro(err)
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	isErro(err)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func exibirOsLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	isErro(err)
	fmt.Println(string(arquivo))
}

func isErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro!!!", err)
		log.Fatal(err)
	}
}
