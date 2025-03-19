package main

import (
	"TrabalhoGo1/service"
	"fmt"
	"strconv"
	"html/template"
	"net/http"
	"log"
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
func listarUsuarios(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro := template.ParseFiles("template/detalhes.html")
	if erro != nil {
		fmt.Println("Erro ao carregar template:", erro)
		http.Error(resposta, "Erro interno", http.StatusInternalServerError)
		return
	}

	// Buscar usuários do banco
	usuarios, err := service.BuscaTodosUsuarios()
	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		http.Error(resposta, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	// Passar a lista de usuários para o template
	pagina.Execute(resposta, usuarios)
}

func abrePesquisa(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/pesquisa.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func pesquisarUsuario(resposta http.ResponseWriter, requisicao *http.Request) {
	pagina, erro := template.ParseFiles("template/pesquisa.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}
	// Extrair os dados do formulário
	idStr := requisicao.URL.Query().Get("id")
	nomeStr := requisicao.URL.Query().Get("nome")

	// Criar um objeto vazio de usuário
	dados := struct {
        Usuario service.Usuario
        Erro    string
    }{}

	if idStr != "" {
        id, convErr := strconv.Atoi(idStr) // Converte string para inteiro
        if convErr != nil {
            dados.Erro = "ID inválido. Insira um número válido."
        } else {
            usuario, err := service.BuscaUsuariosid(id) // Chama a função corrigida
            if err != nil {
                dados.Erro = "Erro ao buscar usuário. Tente novamente."
                log.Println("Erro ao buscar usuário:", err)
            } else if usuario.Id == 0 { // Usuário não encontrado
                dados.Erro = "Usuário não encontrado."
            } else {
                dados.Usuario = usuario
            }
        }
    }

	if nomeStr != "" {
		usuario, err := service.BuscaUsuariosnome(nomeStr) // Chama a função corrigida
		if err != nil {
			dados.Erro = "Erro ao buscar usuário. Tente novamente."
			log.Println("Erro ao buscar usuário:", err)
		}else {
			dados.Usuario = usuario
		}
	}
	
    // Renderizar a página com os dados e erro (se houver)
    pagina.Execute(resposta, dados)
}

func main(){

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/cadastro", abreCadastro)
	http.HandleFunc("/scadastro", salvadoador)
	http.HandleFunc("/detalhes", listarUsuarios)
	http.HandleFunc("/pesquisa", abrePesquisa)
	http.HandleFunc("/pesquisar", pesquisarUsuario)

	erro := http.ListenAndServe("0.0.0.0:8080", nil)
	if erro != nil{
		fmt.Print("Servidor com Problemas.")
	}
	
}
