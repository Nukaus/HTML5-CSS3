alguem = int(input("Digite um valor:"))
lista = [10, 20, 55]
existe = alguem in lista
if existe == True:
    print("Existe")
else:
    print("Não existe")