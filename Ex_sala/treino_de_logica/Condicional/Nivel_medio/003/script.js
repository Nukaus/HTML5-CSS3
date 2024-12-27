function check(){
    const num1 = parseFloat(document.getElementById("number1").value);
    const num2 = parseFloat(document.getElementById("number2").value);
    let mensagemElement = document.getElementById("mensage");
    if (num1 > num2){
        mensagemElement.innerText = `O maior número é ${num1}`;
    }else if ( num1 == num2){
        mensagemElement.innerText = `Os números são iguais`;
    }else{
        mensagemElement.innerText = `O maior número é ${num2}`;
    }
}