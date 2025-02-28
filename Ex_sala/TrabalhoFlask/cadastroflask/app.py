from flask import Flask, render_template, request, redirect, url_for
import mysql.connector

app = Flask(__name__)

def conectar_banco():
    return mysql.connector.connect(
        host="127.0.0.1",
        user="root",
        password="1346",
        database="db_infousers"
    )

@app.route("/")
def home():
    return render_template("index.html")

@app.route('/adicionar', methods=['POST'])
def adicionar_usuario():
    if request.method == 'POST':
        nome = request.form['nome']
        tel = request.form['tel']
        renda = request.form['renda']
        dia = request.form['dia']
        cep = request.form['cep']

        conexao = conectar_banco()
        cursor = conexao.cursor()

        sql = "INSERT INTO usuario (nome_cliente, telefone_cliente, renda_cliente, nasc_cliente, cep_cliente) VALUES (%s, %s, %s, %s, %s)"
        valores = (nome, tel, renda, dia, cep)
        cursor.execute(sql, valores)

        conexao.commit()
        cursor.close()
        conexao.close()

        return redirect('/')
    
if __name__ == "__main__":
    app.run(debug=True)