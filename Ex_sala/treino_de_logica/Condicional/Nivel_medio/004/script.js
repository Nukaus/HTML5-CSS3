function check(){
    const nota = parseFloat(document.getElementById("number").value);
    let mensagemElement = document.getElementById("mensage");
    if (nota >= 8 ){
        mensagemElement.innerText = `Nota A`;
    }else if ( nota >= 6){
        mensagemElement.innerText = `Nota B`;
    }else if ( nota == 5){
        mensagemElement.innerText = `Nota C`;
    }else{
        mensagemElement.innerText = `Nota D`;
    }
}