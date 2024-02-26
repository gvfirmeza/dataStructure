package main

import "fmt"

type Contato struct {
    nome  string
    email string
}

func (ctt *Contato) AlterarEmail(emailNovo string) {
    ctt.email = emailNovo
}

func adicionarContato(ctt Contato, arrayContatos [5]*Contato) {
    for i, c := range arrayContatos {
        if c == nil {
            arrayContatos[i] = &ctt
            fmt.Println("Contato adicionado!")
            return
        }
    }
    fmt.Println("Não foi possível adicionar o contato pois já foram adicionados 5 contatos.")
}

func excluirContato(arrayContatos [5]*Contato) {
    for i := 4; i >= 0; i-- {
        if arrayContatos[i] != nil {
            arrayContatos[i] = nil
            fmt.Println("O último contato foi removido com sucesso.")
            return
        }
    }

    fmt.Println("Nenhum contato foi cadastrado.")
}

func main() {
    var arrayContatos [5]*Contato

    for {
        var opcao int
        fmt.Println("\nDigite:")
        fmt.Println("1 → Adicionar Contato")
        fmt.Println("2 → Excluir o Último Contato")
        fmt.Println("3 → Sair")
        fmt.Scanln(&opcao)

        switch opcao {
        case 1:
            var nome, email string
            fmt.Println("\nDigite:")
            fmt.Print("Nome do contato: ")
            fmt.Scanln(&nome)
            fmt.Print("E-mail do contato: ")
            fmt.Scanln(&email)
            novoContato := Contato{nome: nome, email: email}
            adicionarContato(novoContato, arrayContatos)
        case 2:
            excluirContato(arrayContatos)
        case 3:
            fmt.Println("Programa Encerrado.")
            return
        default:
            fmt.Println("Opção inválida.")
        }
    }
}