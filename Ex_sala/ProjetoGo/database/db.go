package database

import("database/sql"
		"fmt"
		_ "github.com/go-sql-driver/mysql")

func ConectaBanco(){

	//config. da conexão 
	conexao := "root:1346@tcp(127.0.0.1:3306)/db_hospital"

	//abrir conexão - informe o driver e a string de conexão
	db, err := sql.Open("mysql", conexao)
	
	if err != nil{
		fmt.Print("Erro ao conectar ao banco", err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Print("Erro ao buscar o servidor", err)
	}else{
		fmt.Print("Conexão realizada com sucesso")
	}
}  
