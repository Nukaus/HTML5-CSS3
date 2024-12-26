function check(){
    const vota = parseFloat(document.getElementById("number").value);
    let mensagemElement = document.getElementById("mensage");
    if (vota >= 18){
        mensagemElement.innerText = "Você já pode votar!!";
    }else{
        mensagemElement.innerText = "Você ainda não pode votar!!";
    }
}