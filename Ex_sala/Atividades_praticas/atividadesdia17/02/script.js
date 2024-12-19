function check(){
    let verificar = document.getElementById("senha").value;
    if (verificar == "1234"){
        document.getElementById("mensagem").innerText = `ACESSO	PERMITIDO`;
    }
    else{
        document.getElementById("mensagem").innerText = `ACESSO	NEGADO`;
    }
}