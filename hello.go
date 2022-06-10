package main

import (
	"fmt"
	"os"
)

const MONITORAR int = 1
const INICIAR_LOGS int = 2
const SAIR_DO_PROGRAMA int = 0

func main() {

	exibeInformacoes()
	exibeMenu()
	printaTipoDeComando(readComando())
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
	case MONITORAR:
		fmt.Println("Monitorando...")
	case INICIAR_LOGS:
		fmt.Println("Iniciando os logs...")
	case SAIR_DO_PROGRAMA:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando!")
		os.Exit(-1)
	}
}
