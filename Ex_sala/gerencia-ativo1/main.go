package main

import (
	"fmt"
	"strconv"
	"html/template"
	"net/http"
	"log"
	"gerencia-ativo1/service"
)

func abreCadastro(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/cadastro.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func cadastraAtivo(resposta http.ResponseWriter, requisicao *http.Request){
	pagina, erro:=template.ParseFiles("template/cadastro.html")

	if erro != nil{
		fmt.Println("O erro foi", erro)
		return
	}

	if requisicao.Method == http.MethodPost{
		nomeUser := requisicao.FormValue("nome")		
		numero_serieUser := requisicao.FormValue("numero_serie")		
		categoriaUser := requisicao.FormValue("categoria")	
		localizacaocUser := requisicao.FormValue("localizacao")
		responsavelUser := requisicao.FormValue("responsavel")
		data_aquisicaoUser := requisicao.FormValue("data_aquisicao")
		statusAtivoUser := requisicao.FormValue("statusAtivo")
		criadoUser := requisicao.FormValue("criado_em")
		
		
		
		service.InsereDoador(nomeUser, numero_serieUser, categoriaUser, localizacaocUser, responsavelUser, data_aquisicaoUser, statusAtivoUser, criadoUser)
	}
	pagina.Execute(resposta, nil)
}
func listarAtivos(resposta http.ResponseWriter, requisicao *http.Request){
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

func pesquisarAtivo(resposta http.ResponseWriter, requisicao *http.Request) {
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
                dados.Erro = "Erro ao buscar ativo. Tente novamente."
                log.Println("Erro ao buscar ativo:", err)
            } else if usuario.Id == 0 { // Usuário não encontrado
                dados.Erro = "Ativo não encontrado."
            } else {
                dados.Usuario = usuario
            }
        }
    }

	if nomeStr != "" {
		usuario, err := service.BuscaUsuariosnome(nomeStr) // Chama a função corrigida
		if err != nil {
			dados.Erro = "Erro ao buscar ativo. Tente novamente."
			log.Println("Erro ao buscar ativo:", err)
		}else {
			dados.Usuario = usuario
		}
	}
	
    // Renderizar a página com os dados e erro (se houver)
    pagina.Execute(resposta, dados)
}

func excluirAtivo(resposta http.ResponseWriter, requisicao *http.Request) {
	// Recuperar o ID do usuário da URL
	idStr := requisicao.URL.Path[len("/deletar/"):]

	// Logando o ID para verificação
	log.Println("ID recebido para exclusão: ", idStr)

	// Converter o ID para inteiro
	id, erro := strconv.Atoi(idStr)
	if erro != nil {
		http.Error(resposta, "ID inválido", http.StatusBadRequest)
		return
	}

	// Chamar o serviço para excluir o usuário do banco
	err := service.DeletarUsuario(id)
	if err != nil {
		http.Error(resposta, "Erro ao excluir ativo", http.StatusInternalServerError)
		return
	}

	// Redirecionar de volta para a página de detalhes após a exclusão
	http.Redirect(resposta, requisicao, "/detalhes", http.StatusSeeOther)
}


func main(){

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/cadastro", abreCadastro)
	http.HandleFunc("/scadastro", cadastraAtivo)
	http.HandleFunc("/detalhes", listarAtivos)
	http.HandleFunc("/pesquisa", abrePesquisa)
	http.HandleFunc("/pesquisar", pesquisarAtivo)
	http.HandleFunc("/deletar/", excluirAtivo)

	erro := http.ListenAndServe("0.0.0.0:8080", nil)
	if erro != nil{
		fmt.Print("Servidor com Problemas.")
	}
	
}
