
# Golang - Primeiros passos

Estudando GOlang, desde o princípio, como declarar variável, pacotes, função, etc.

### Módulo 1 - Introdução ao Golang.

Todo método que vem de um pacote externo, inicia com letra maiúscula.

fmt.Println("Olá Sr.", nome)

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

### Módulo 3 - Controlando o fluxo do script

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

### Módulo 4 - Fazendo requisições para a web

	nome, idade := getNomeEIdade()
	fmt.Println(nome, idade)
}

func getNomeEIdade() (string, int) {
	nome := "Dan"
	idade := 32
	return nome, idade
}

no exemplo acima, temos uma função onde retorna  dois valores. Para informar os tipos, é necessário colocar dentro de uma parênteses. 

Ao declarar, separamos por vírgula, como no exemplo acima. Mas digamos que eu queira apenas receber a idade? 
No Golang, usamos o underline quando não queremos usar uma variável

_, idade := getNomeEIdade()
fmt.Println(idade)

Existe um pacote que facilita a requisição web, é o "net/http". Ele tem um método Get, onde podemos passar o link do site. 
E ele retorna algumas informações referente a esse site, como body, header, status e statuscode. 

O método get retorna o response e erro. Podemos trabalhar apenas com o response, de imediato. 

Caso eu queira fazer um loop infinito, no golang, não usamos o while, por exemplo, mas o for, sem passar uma condicional para ele, da seguinte maneira.

func main() {

	exibeInformacoes()
	for {
		exibeMenu()
		printaTipoDeComando(readComando())
	}
}

os dois métodos dentro do for vão ficar repetindo até que o usuário saia do programa.

