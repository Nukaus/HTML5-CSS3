function check(){
    const nota = parseFloat(document.getElementById("number").value);
    let mensagemElement = document.getElementById("mensage");
    if (nota >= 7){
        mensagemElement.innerText = "Você foi aprovado!!";
    }else{
        mensagemElement.innerText = "Você foi reprovado!!";
    }
}