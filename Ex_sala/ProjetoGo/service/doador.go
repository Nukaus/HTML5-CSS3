package service

import ("projeto/database"
		"fmt"
)

const insertComando string = "insert into doador(nome_doador, telefone_doador, email_doador, tipo) values(?, ?, ?, ?)"

func InsereDoador(nome string, telefone string, email string, tipo string) {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Exec(insertComando, nome, telefone, email, tipo)
	if erro != nil{
		fmt.Print("Erro ao inserir doaodr", erro)
	}else{
		fmt.Print(resultado)
	}
}