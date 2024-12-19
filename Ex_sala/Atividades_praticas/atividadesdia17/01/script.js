function check(){
    const check1 = parseFloat(document.getElementById("num1").value);
    const check2 =  parseFloat(document.getElementById("num2").value);
    
    if (check1 > check2) {
        document.getElementById("maior").innerText = `O maior número é: ${check1}`;
    }
    else {
        document.getElementById("maior").innerText = `O maior número é: ${check2}`
    }
}
