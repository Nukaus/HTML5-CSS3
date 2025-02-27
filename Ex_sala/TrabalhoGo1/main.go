package main

import (
	"TrabalhoGo1/service"
	"fmt"
	"html/template"
	"net/http"
)


func abreCadastro(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/cadastro.html")

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
		nomeUser := requisicao.FormValue("nome")		
		telefoneUser := requisicao.FormValue("tel")		
		rendaUser := requisicao.FormValue("renda")	
		nascUser := requisicao.FormValue("dia")
		cepUser := requisicao.FormValue("cep")
		
		
		service.InsereDoador(nomeUser, telefoneUser, rendaUser, nascUser, cepUser)
	}
	pagina.Execute(resposta, nil)
}

func main(){
	// Servindo arquivos estáticos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	//Criando um EndPoint
	// http.HandleFunc("/inicial", abreIndex)
	http.HandleFunc("/cadastro", abreCadastro)
	// http.HandleFunc("/login", abreLogin)
	http.HandleFunc("/scadastro", salvadoador)
	// http.HandleFunc("/fazerlogin", fazlogin)

	//Subindo Servidor
	erro := http.ListenAndServe("0.0.0.0:8080", nil)
	if erro != nil{
		fmt.Print("Servidor com Problemas.")
	}
	
}