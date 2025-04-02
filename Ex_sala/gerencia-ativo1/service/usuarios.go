package service

import ("gerencia-ativo1/database"
		"database/sql"
		"fmt"
		"log"
		
)

const insertComando string = "insert into ativos(nome, numero_serie, categoria, localizacao, responsavel,  statusAtivo, data_aquisicao, criado_em) values(?, ?, ?, ?, ?, ?, ?, ?)"

func InsereDoador(nome string, numero_serie string, categoria string, localizacao string, responsavel string,data_aquisicao string, statusAtivo string, criado_em string) {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Exec(insertComando, nome, numero_serie, categoria, localizacao, responsavel,  statusAtivo, data_aquisicao, criado_em)
	

	fmt.Println("Valor recebido para status: ", statusAtivo)
	
	if erro != nil{
		fmt.Print("Erro ao inserir produto ", erro)
	}else{
		fmt.Print(resultado)
	}
}

const selectComando string = "SELECT id, nome, numero_serie, categoria, localizacao, responsavel,  statusAtivo, data_aquisicao, criado_em FROM ativos"

type Usuario struct {
	Id             int
	Nome           string
	NumeroSerie    string
	Categoria      string
	Localizacao    string
	Responsavel    string
	StatusAtivo    string
	DataAquisicao  string
	CriadoEm       string
}


func BuscaUsuariosid(id int) (Usuario, error){
	conexao := database.ConectaBanco()
	defer conexao.Close()

	selectComando := "SELECT id, nome, numero_serie, categoria, localizacao, responsavel, data_aquisicao,  statusAtivo, criado_em  FROM ativos WHERE id = ?"

	var usuario Usuario
	// Usando QueryRow para buscar um único usuário
	err := conexao.QueryRow(selectComando, id).Scan(&usuario.Id, &usuario.Nome, &usuario.NumeroSerie, &usuario.Categoria, &usuario.Localizacao, &usuario.Responsavel, &usuario.DataAquisicao, &usuario.StatusAtivo, &usuario.CriadoEm)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Ativo não encontrado")
			return Usuario{}, err
		}
	}
	return usuario, nil
}

func BuscaUsuariosnome(nome string) (Usuario, error){
	conexao := database.ConectaBanco()
	defer conexao.Close()

	selectComando := "SELECT id, nome, numero_serie, categoria, localizacao, responsavel,  data_aquisicao, statusAtivo, criado_em  FROM ativos WHERE nome LIKE ?"

	var usuario Usuario
	// Usando QueryRow para buscar um único usuário
	nomeBusca := "%" + nome + "%"
	err := conexao.QueryRow(selectComando, nomeBusca).Scan(&usuario.Id, &usuario.Nome, &usuario.NumeroSerie, &usuario.Categoria, &usuario.Localizacao, &usuario.Responsavel, &usuario.DataAquisicao, &usuario.StatusAtivo, &usuario.CriadoEm)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Ativo não encontrado")
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
		log.Println("Erro ao buscar ativos:", err)
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario

	
	for rows.Next() {
		var user Usuario
		err := rows.Scan(&user.Id, &user.Nome, &user.NumeroSerie, &user.Categoria, &user.Localizacao, &user.Responsavel, &user.DataAquisicao, &user.StatusAtivo, &user.CriadoEm)
		if err != nil {
			log.Println("Erro ao ler ativos:", err)
			continue
		}
		usuarios = append(usuarios, user)
	}

	return usuarios, nil
}

func DeletarUsuario(id int) error {
	conexao := database.ConectaBanco()
	defer conexao.Close()

	// Comando SQL para excluir o usuário com o ID fornecido
	deleteComando := "DELETE FROM ativos WHERE id = ?"

	// Executar o comando
	_, err := conexao.Exec(deleteComando, id)
	if err != nil {
		log.Println("Erro ao excluir o ativo:", err)
		return err
	}

	return nil
}
