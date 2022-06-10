
# Golang - Primeiros passos

Estudando GOlang, desde o princípio, como declarar variável, pacotes, função, etc.

### Módulo 1 - Introdução ao Golang.

Todo método que vem de um pacote externo, inicia com letra maiúscula.

![alt text](https://raw.githubusercontent.com/DanSaRecTech/primeiros-passos-golang/master/curso%20golang/fmt_Println.png)

nota-se o uso do pacote fmt e seu método Println.

Para buildar um arquivo em Go, é necessário ir até a pasta onde se encontra o projeto, via terminal:
-- go build arquivo.go -> esse comando irá criar o arquivo.exe

-- ./arquivo -> irá executar o arquivo. 
Mas tem um caminho mais fácil para buildar e executar no mesmo comando
-- go run arquivo.go

### Módulo 2 - Variáveis

Quando não atribui um valor à variável, do tipo numérico, ele assume 0. Para string, um campo vazio (''), para boolean false.
O var não se faz necessários, desde que inicie a variável da seguinte forma: nome := "Dan". A linguagem infere e identifica que se trata de uma string. 
Caro seja declarado uma variável e não usemos, ele informa um erro. Assim, não é possível buildar o projeto com uma variável sem tá em uso.

### Módulo 2 - Controlando o fluxo do script

O if é implementado sem os parênteses, da seguinte forma:

	if readComando() == 1 {
		fmt.Println("Monitorando...")
	} else if readComando() == 2 {
		fmt.Println("Iniciando os logs...")
	} else if readComando() == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço esse comando!")
	}

Também tem o switch, que é implementado da seguinte forma:

	switch comando {
	case MONITORAR:
		fmt.Println("Monitorando...")
	case INICIAR_LOGS:
		fmt.Println("Iniciando os logs...")
	case SAIR_DO_PROGRAMA:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço esse comando!")
	}

O fmt, tem uma função que se chama scanf, onde você passa, como parâmetro, o tipo de dados que ele vai receber e o ponteiro da variável. 

Exemplo. Acima, tem a variável comando. E quero receber do usuário qual comando ele quer executar.

var comando int
fmt.scanf("%d", &comando) -> O & aponta para o local da memória onde a variável comando ficará alocada.

Porém, como eu informo o tipo na declaração da variável, não é necessário, no parâmetro, passar "%d" para indicar que o parâmetro é do tipo inteiro.

Assim, podemos usar o scan (sem f) e passar apenas o ponteiro comando. 
-- fmt.scan(&comando)

