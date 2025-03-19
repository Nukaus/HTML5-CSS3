package service

import ("TrabalhoGo1/database"
		"database/sql"
		"fmt"
		"log"
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

const selectComando string = "SELECT id_cliente, nome_cliente, telefone_cliente, renda_cliente, nasc_cliente, cep_cliente FROM usuario"

type Usuario struct {
	Id		 int
	Nome     string
	Telefone string
	Renda    float64
	Dia      string
	Cep      string
}

func BuscaUsuariosid(id int) (Usuario, error){
	conexao := database.ConectaBanco()
	defer conexao.Close()

	selectComando := "SELECT id_cliente, nome_cliente, telefone_cliente, renda_cliente, nasc_cliente, cep_cliente  FROM usuario WHERE id_cliente = ?"

	var usuario Usuario
	// Usando QueryRow para buscar um único usuário
	err := conexao.QueryRow(selectComando, id).Scan(&usuario.Id, &usuario.Nome, &usuario.Telefone, &usuario.Renda, &usuario.Dia, &usuario.Cep)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuário não encontrado")
			return Usuario{}, err
		}
	}
	return usuario, nil
}

func BuscaUsuariosnome(nome string) (Usuario, error){
	conexao := database.ConectaBanco()
	defer conexao.Close()

	selectComando := "SELECT id_cliente, nome_cliente, telefone_cliente, renda_cliente, nasc_cliente, cep_cliente  FROM usuario WHERE nome_cliente LIKE ?"

	var usuario Usuario
	// Usando QueryRow para buscar um único usuário
	nomeBusca := "%" + nome + "%"
	err := conexao.QueryRow(selectComando, nomeBusca).Scan(&usuario.Id, &usuario.Nome, &usuario.Telefone, &usuario.Renda, &usuario.Dia, &usuario.Cep)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuário não encontrado")
			return Usuario{}, err
		}
	}
	return usuario, nil
}

func BuscaTodosUsuarios() ([]Usuario, error) {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	
	rows, err := conexao.Query(selectComando)
	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario

	
	for rows.Next() {
		var user Usuario
		err := rows.Scan(&user.Id, &user.Nome, &user.Telefone, &user.Renda, &user.Dia, &user.Cep)
		if err != nil {
			log.Println("Erro ao ler usuário:", err)
			continue
		}
		usuarios = append(usuarios, user)
	}

	return usuarios, nil
}