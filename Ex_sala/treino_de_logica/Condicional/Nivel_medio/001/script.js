function check(){
    const number = parseFloat(document.getElementById("num1").value);
    let mensagemElement = document.getElementById("mensage");
    if (number % 2 == 0){
        mensagemElement.innerText = "O número digitado é Par";
    }else{
        mensagemElement.innerText = "O número digitado é Ímpar";
    }
}