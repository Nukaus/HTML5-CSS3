function check(){
    const num = parseFloat(document.getElementById("number").value);
    let mensagemElement = document.getElementById("mensage");
    if (num == 0){
        mensagemElement.innerText = "O número é 0.";
    }else if (num > 0){
        mensagemElement.innerText = "O número é positivo.";
    }else{
        mensagemElement.innerText = "O número é negativo.";
    }
}