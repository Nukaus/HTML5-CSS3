function check(){
    let verificar = document.getElementById("quantidade").value;
    if (verificar <= 12){
        const preco1 = (verificar * 0.30)
        document.getElementById("mensagem").innerText = `O valor da sua compra foi de ${preco1}`;
    }
    else{
        const preco2 = (verificar * 0.25)
        document.getElementById("mensagem").innerText = `O valor da sua compra foi de ${preco2}`;
    }
}