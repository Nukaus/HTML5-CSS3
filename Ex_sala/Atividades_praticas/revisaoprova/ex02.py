import re 
texto = "Ame a ema"
limpo = re.sub(r'[^a-zA-Z]', '', texto.lower())
palindromo = limpo == limpo[::-1]
print(palindromo)