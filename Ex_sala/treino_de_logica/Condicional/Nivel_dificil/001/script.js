function check(){
    const a = parseFloat(document.getElementById("number1").value);
    const b = parseFloat(document.getElementById("number2").value);
    const c = parseFloat(document.getElementById("number3").value);
    let mensagemElement = document.getElementById("mensage");
    if (a + b > c && a + c > b && b + c > a){
        mensagemElement.innerText = `Com essas três retas você pode formar um triângulo`;
    }else{
        mensagemElement.innerText = `Você não pode formar um triângulo com essas retas`;
    }
}