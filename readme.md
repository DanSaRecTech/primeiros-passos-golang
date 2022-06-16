
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

### As principais coleções em GO

Arrays - Um array em go tem um tamanho fixo, limitando assim o uso do mesmo => var sites [4]string
Caso não queira informar um array com tamanho fixo, o go disponibiliza o slice, que é uma abstração do array, mas com algumas funcionalidades a mais
Abaixo algumas funções do slice

func exibeNomes() {
	nomes := []string{"Dan", "Maju", "Lorena"}

	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O tamanho do meu slice:", len(nomes))
	fmt.Println("O meu slice tem a capacidade de:", cap(nomes))
	nomes = append(nomes, "Alice")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O tamanho do meu slice:", len(nomes))
	fmt.Println("O meu slice tem a capacidade de:", cap(nomes))
}

o len é para informar o tamanho do array, naquele momento.
o cap é para informar a capacidade que o array tem. No exemplo acima, no primeiro cap, ele tem tamanho 3 e capacidade 3. 
Depois que eu adiciono mais 1 elemento, na função append, o len do meu array para para 4 mas o cap dobra para 6. Isso acontece
quando um array é estourado. O Go duplica o tamanho dele e coloca o novo valor.

Existem 2 formas de interar. A tradicional e range

## Abaixo a forma tradicional:

for i := 0; i < len(sites); i++ {

		resp, _ := http.Get(sites[i])

		if resp.StatusCode == 200 {
			fmt.Println("Site:", sites, "foi carregado com sucesso! Status:", resp.Status)
		} else {
			fmt.Println("Site:", sites[i], "está com problemas!!! Staus:", resp.StatusCode)
		}
	}

## Abaixo a forma com range

	for _, site := range sites {

		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso! Status:", resp.Status)
		} else {
			fmt.Println("Site:", site, "está com problemas!!! Staus:", resp.StatusCode)
		}
	}

O primeiro parâmetro é o índice que ele vai tá passando naquele momento dentro do for e a segunda variável é o valor daquele índice. 

-- O Go tem um pacote times onde ele dispõe de alguns métodos interessantes. Um deles é o sleep. Com esse método, podemos pedir para uma parte do nosso código
-- aguardar um período para poder ser executado de novo. No exemplo abaixo, usamos 5 segundos. Mas pode ser minutos, horas, etc.

for i := 0; i < 5; i++ {
		for indice, site := range sites {
			fmt.Println("")
			fmt.Println("Testando site", indice, ":", site)
			verificaSiteOnline(site)
		}
		time.Sleep(5 * time.Second)
	}

### Lendo arquivos

Em go, o pacote responsável por se comunicar com o sistema de operação do computador é o "os". 
Nele, temos o método open, onde passamos o arquivo que queremos ler e ele busca as informações. 
Esse Open() tem dois parâmetros: o arquivo, em si, e um erro, como já vimos acima. 

arquivo, err := os.Open("sites.txt")
result -> &{0xc0000d0f00}

Observando acima, em vez de informações válidas do arquivo, retornou os bytes do arquivo, pois, o Open() literalmente pega o arquivo da memória do computador e taca dentro da variável arquivo. Ou seja, um ponteiro puro. 
Há outras formas de ler o arquivo, como o da biblioteca "ioutil" que tem a função ReadFile().

A funçao acima retorna um array de bytes, mas ele acaba sendo mais fácil de transformar numa string. Basta envolver a variável arquivo em uma string. Abaixo um exemplo.

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo, err := ioutil.ReadFile("sites.txt")
	isErro(err)
	fmt.Println(string(arquivo))
	return sites
}

Com o ReadFile, ele retorna todos os textos dentro do arquivo. Caso a gente precise ler linha a linha, ou fazer algo com a informação dentro do arquivo, o recomendado é usar o pacote bufio, onde se encontra o método NewRead(), onde passamos no parâmetro o arquivo, que abrimos com o os.Open(arquivo.txt)

O método acima, tem algumas facilidades, para podermos ler o arquivo linha a linha.
Caso queira ler linha a linha, usamos um método chamado ReadString, onde passamos o comando de pulo de linha, que é o \n.

Exemplo do código:

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	isErro(err)
	leitor := bufio.NewReader(arquivo)
	for {
		linha, _ := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
	}
	return sites
}

O TrimSpace é usado, pois, no final de cada linha, existe um espaço e um \n. Assim, ele tira esse espaço e tabulações.
No código acima, o for é usado para ler todas as linhas do arquivo. Abaixo, uma implementação para validar se ele chegou na última linha. Caso sim, ele para a execução e retorna as informações do arquivo.

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

### Escrevendo arquivos

Caso seja necessário abrir um arquivo e escrever informações nele, como um log, por exemplo, podemos usar o openFile do pacote "os".

func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0006)
	isErro(err)
	fmt.Println(arquivo)
}

Acima, ele pede três parâmetros no OpenFile...o nome do arquivo, a flag e a permissão. A flag é o tipo de de comando que o OpenFile vai poder executar. No caso acima, por exemplo, o RDWR abre o arquivo de leitura e grava informações nele. O CREATE, cria um novo arquivo, caso não exista. E o último parâmetro é a permissão que ele vai ter.

https://pkg.go.dev/os - Documentação

Existe um método para escrever as linhas no arquivo que eu criei ou abrir, que é o WriteString
Passo no parâmetro as informações, em string, que eu desejo escrever e, no fim, concateno com um "\n". 

func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	isErro(err)

	arquivo.WriteString(site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

No método acima, transformamos o bool em string, usando o pacote strconv e o método FormatBool, passando o status, que é o bool.

O Go tem uma forma diferente de formatar sua data. Abaixo um exemplo.

arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

Documentação sobre esse padrão de formatação: https://go.dev/src/time/format.go



