function check(){
    const ano = parseFloat(document.getElementById("number").value);
    let mensagemElement = document.getElementById("mensage");
    if (ano % 4 == 0){
        if (ano % 100 == 0){
            if(ano % 400 == 0){ 
                mensagemElement.innerText = "O ano é bissexto";
            }else{
                mensagemElement.innerText = "O ano não é bissexto";
            }
        }else{
            mensagemElement.innerText = "O ano é bissexto";
        }
    }else{
        mensagemElement.innerText = "O ano não é bissexto";
    }
}