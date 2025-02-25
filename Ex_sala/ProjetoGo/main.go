package main

import (
	"fmt"
	"html/template"
	"net/http"
	"projeto/database"
)

func abreIndex(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/index.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func abreCadastro(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/cadastro.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func abreLogin(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/login.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func salvadoador(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/cadastro.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}

	if requisicao.Method == http.MethodPost{
		nomeDoador := requisicao.FormValue("nome")
		fmt.Println("O nome do usuario é", nomeDoador)
		nomeTel := requisicao.FormValue("tel")
		fmt.Println("O telefone do usuario é", nomeTel)
		nomeEmail := requisicao.FormValue("email")
		fmt.Println("O email do usuario é", nomeEmail)
		nomeTsan := requisicao.FormValue("tiposanguineo")
		fmt.Println("O Tipo Sanguineo do usuario é", nomeTsan)
		database.ConectaBanco()
	}
	pagina.Execute(resposta, nil)
}

func fazlogin(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/login.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}

	if requisicao.Method == http.MethodPost{
		emailDoador := requisicao.FormValue("email")
		fmt.Println("O e-mail do usuario é", emailDoador)
		senhaDoador := requisicao.FormValue("senha")
		fmt.Println("A senha do usuario é", senhaDoador)
	}
	pagina.Execute(resposta, nil)
}

func main(){
	// Servindo arquivos estáticos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	//Criando um EndPoint
	http.HandleFunc("/inicial", abreIndex)
	http.HandleFunc("/cadastro", abreCadastro)
	http.HandleFunc("/login", abreLogin)
	http.HandleFunc("/scadastro", salvadoador)
	http.HandleFunc("/fazerlogin", fazlogin)

	//Subindo Servidor
	erro := http.ListenAndServe("0.0.0.0:8080", nil)
	if erro != nil{
		fmt.Print("Servidor com Problemas.")
	}
	
}