package service

import ("TrabalhoGo1/database"
		"fmt"
)

const insertComando string = "insert into usuario(nome_cliente, telefone_cliente, renda_cliente, nasc_cliente, cep_cliente) values(?, ?, ?, ?, ?)"

func InsereDoador(nome string, telefone string, renda string, nasc string, cep string) {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Exec(insertComando, nome, telefone, renda, nasc, cep)
	if erro != nil{
		fmt.Print("Erro ao inserir cliente", erro)
	}else{
		fmt.Print(resultado)
	}
}