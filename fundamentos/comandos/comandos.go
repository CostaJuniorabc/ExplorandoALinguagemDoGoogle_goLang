package main

import "fmt"

func main() {

	fmt.Println("Outro programa em %s!", "Go")
}

/* comando para terminal

ls = comando para listar o que tem na pasta

go =  mostra varias possibilidades de uso do go

go help comando  = ele mostro o que o comando faz ex: go help get

go version = para verificar a versão do go

godoc -http:6060 = serve para abilitar a documentação offline do go ( abrir o navegador localhost:6060)

go env = da varias informaçoes do ambiente

go doc cmd/vet =  mostra a documentação desse comando ( faz verificação no codigo e reporta ) ex go vet comandos.go

go build comandos.go = ele gera um arquivo .go e pode executar ele dai vc coloca ./comandos.go ele executa

go run comandos.go = ele compila e executa o programa em um unuco comando

ls /home/ancarlosda/go/src = mostra o que esta instalado no go

ls /home/ancarlosda/go/src/github.com/ =

go get -u github.com/go-sql-driver/mysql = comando para baixar dependencias



go install github.com/go-sql-driver/mysql@latest

*/
