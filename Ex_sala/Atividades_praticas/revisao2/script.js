function check(){
    const check = parseFloat(document.getElementById("checkano").value);
    const resultado = 2024 - check;
    document.getElementById("idade").innerText = `Idade: ${resultado}`;
    const container = document.getElementById("imagens");

    if (resultado >= 18){
        document.getElementById("pode").innerText = `Você já pode tirar habilitação!!!`;
        container.innerHTML = '<img src="imagens/car200.jpg" alt="Carro">';
    }
    else{
        document.getElementById("pode").innerText = `Você NÂO pode tirar habilitação!!!`;
        container.innerHTML = '<img src="imagens/bike200.jpg" alt="Bicicleta">';
    }
}