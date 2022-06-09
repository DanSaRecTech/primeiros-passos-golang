package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Dan"
	versao := 1.1
	idade := 32

	fmt.Println("Olá Sr.", nome, "Sua idade é:", idade)
	fmt.Println("Este programa se encontra na versão: ", versao)

	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
}
