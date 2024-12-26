function check(){
    const idade1 = parseFloat(document.getElementById("number").value);
    let mensagemElement = document.getElementById("mensage");
    if (idade1 >= 18){
        mensagemElement.innerText = "Você é maior de idade!!";
    }else{
        mensagemElement.innerText = "Você é menor de idade!!";
    }
}