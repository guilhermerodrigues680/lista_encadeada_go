package tela

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/guilhermerodrigues680/lista_encadeada_go/cmd/pessoa"
)

var listPessoa *list.List = list.New()

func addInitialPersonToList() {

	for i := 0; i < 32767; i++ {
		p1 := pessoa.NewPessoa("Emanuel", "Oliveira", 22)
		p2 := pessoa.NewPessoa("Maria", "Araujo", 25)
		listPessoa.PushBack(p1)
		listPessoa.PushBack(p2)
	}

}

func horizontalRuleCLI() string {
	nHyphen := 26
	hr := strings.Repeat("-", nHyphen)
	return "\u001B[33m" + hr + "\u001B[0m"
}

func formatTextCLI(text string) string {
	return fmt.Sprintf("\u001B[33m|\u001B[0m %-22s \u001B[33m|\u001B[0m", text)
}

func menuOptions() string {
	menu := ""
	menu += horizontalRuleCLI() + "\n"
	menu += formatTextCLI("1 - Inserir no inicio") + "\n"
	menu += formatTextCLI("2 - Inserir no final") + "\n"
	menu += formatTextCLI("3 - Localizar nó") + "\n"
	menu += formatTextCLI("4 - Excluir nó") + "\n"
	menu += formatTextCLI("5 - Exibir lista") + "\n"
	menu += formatTextCLI("6 - Tamanho da lista") + "\n"
	menu += formatTextCLI("7 - Sair") + "\n"
	menu += horizontalRuleCLI() + "\n"

	return menu
}

func readPessoaCLI() pessoa.Pessoaer {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nome -> ")
	scanner.Scan()
	nome := scanner.Text()

	fmt.Print("Sobrenome -> ")
	scanner.Scan()
	sobrenome := scanner.Text()

	fmt.Print("Idade -> ")
	scanner.Scan()
	idade, _ := strconv.Atoi(scanner.Text())

	return pessoa.NewPessoa(nome, sobrenome, idade)
}

// Start inicia o loop do menu
func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	addInitialPersonToList()

	for {
		fmt.Println(menuOptions())
		fmt.Print("Digite a opcao (1-7)\n-> ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "1":
			listPessoa.PushFront(readPessoaCLI())

		case "2":
			listPessoa.PushBack(readPessoaCLI())

		case "3":
			fmt.Print("Nome -> ")
			scanner.Scan()
			searchPessoa := scanner.Text()

			var researchedPerson pessoa.Pessoaer = nil

			// Loop over container list.
			for temp := listPessoa.Front(); temp != nil; temp = temp.Next() {
				pTemp := temp.Value.(pessoa.Pessoaer) // Type assertions

				if pTemp.GetNome() == searchPessoa {
					researchedPerson = pTemp
					fmt.Println(researchedPerson)
					break
				}
			}

			if researchedPerson == nil {
				fmt.Println("Nenhuma pessoa encontrada")
			}

		case "4":
			fmt.Print("Nome -> ")
			scanner.Scan()
			searchPessoa := scanner.Text()

			var personRemoved pessoa.Pessoaer = nil

			// Loop over container list.
			for temp := listPessoa.Front(); temp != nil; temp = temp.Next() {
				pTemp := temp.Value.(pessoa.Pessoaer) // Type assertions

				if pTemp.GetNome() == searchPessoa {
					personRemoved = listPessoa.Remove(temp).(pessoa.Pessoaer)
					fmt.Println(personRemoved)
					break
				}
			}

			if personRemoved == nil {
				fmt.Println("Nenhuma pessoa removida")
			}

		case "5":
			// Loop over container list.
			for temp := listPessoa.Front(); temp != nil; temp = temp.Next() {
				pTemp := temp.Value.(pessoa.Pessoaer) // Type assertions
				// fmt.Print(pTemp, "\t")
				fmt.Println(pTemp, &pTemp)
			}
			fmt.Println()

		case "6":
			fmt.Println("Tamanho: ", listPessoa.Len())

		case "7":
			os.Exit(0)

		default:
			fmt.Println("Comando inválido")
		}

	}

}
