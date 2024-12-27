function check(){
    const a = parseFloat(document.getElementById("number1").value);
    let mensagemElement = document.getElementById("mensage");
    if (a <= 1){
        mensagemElement.innerText = `O número não é primo`;
    }else if( a % a == 0){
        
    }
}