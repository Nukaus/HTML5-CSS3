const elementos = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
const maxTentativas = 3;
let tentativas = 0; // Contador global de tentativas

function check() {
    const check = parseFloat(document.getElementById("num1").value);
    const mensagemElement = document.getElementById("mensagem");

    // Verificar se ainda há tentativas disponíveis
    if (tentativas >= maxTentativas) {
        mensagemElement.innerText = "Acabaram as tentativas!";
        return;
    }

    // Verificar se o número está na lista
    if (elementos.includes(check)) {
        mensagemElement.innerText = "SIM";
        return;
    } else {
        tentativas++; // Incrementa as tentativas
        if (tentativas < maxTentativas) {
            mensagemElement.innerText = `NÃO. Você ainda tem ${maxTentativas - tentativas} tentativa(s).`;
        } else {
            mensagemElement.innerText = "Acabaram as tentativas!";
        }
    }
}