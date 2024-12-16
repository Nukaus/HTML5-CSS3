function somar() {
    const num1 = parseFloat(document.getElementById("numero1").value);
    const num2 = parseFloat(document.getElementById("numero2").value);
    const resultado = num1 + num2;
    document.getElementById("resultado").innerText = `Resultado: ${resultado}`;
}

function subtrair() {
    const num1 = parseFloat(document.getElementById("numero1").value);
    const num2 = parseFloat(document.getElementById("numero2").value);
    const resultado = num1 - num2;
    document.getElementById("resultado").innerText = `Resultado: ${resultado}`;
}

function multiplicar() {
    const num1 = parseFloat(document.getElementById("numero1").value);
    const num2 = parseFloat(document.getElementById("numero2").value);
    const resultado = num1 * num2;
    document.getElementById("resultado").innerText = `Resultado: ${resultado}`;
}

function dividir() {
    const num1 = parseFloat(document.getElementById("numero1").value);
    const num2 = parseFloat(document.getElementById("numero2").value);
    if (num2 === 0) {
        document.getElementById("resultado").innerText = `Erro: Divisão por zero não permitida.`;
    } else {
        const resultado = num1 / num2;
        document.getElementById("resultado").innerText = `Resultado: ${resultado}`;
    }
}

function raizquadrada() {
    const num1 = parseFloat(document.getElementById("numero1").value);
    const resultado = Math.sqrt(num1);
    document.getElementById("resultado").innerText = `Resultado: ${resultado}`;
}