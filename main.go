package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

// type Funcionario struct {
	// pis  string
	// nome string
// }

type Registro struct {
	data string
	hora string
	pis  string
}

func main() {
	registro_afd, err := os.Open("./afd_20260922.txt")

	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}

	defer registro_afd.Close()

	scanner := bufio.NewScanner(registro_afd)

	for scanner.Scan() {
		registro := scanner.Text()
		if len(registro) < 38 || registro[9] != '3' {
			continue
		}

		ponto := Registro{
			data: registro[10:18],
			hora: registro[18:22],
			pis: registro[23:34],
		}
		
		fmt.Printf("data: %s\nhora: %s\npis: %s\n", ponto.data, ponto.hora, ponto.pis)
	}

	// usuario_afd, err := os.ReadFile("./usuarios")
	// usuarios := strings.Split(string(usuario_afd), "\n")
	// for _, usuario := range usuarios {

	// if strings.Contains(usuario, "pis") {
	// continue
	// }

	// infos := strings.Split(usuario, ";")

	// if len(infos) >= 2 {
	// staff := Funcionario{infos[0], infos[1]}
	// fmt.Printf("pis: %s\nnome: %s\n", staff.pis, staff.nome)
	// }
	// }
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro crítico durante a leitura do arquivo: %v", err)
	}
}
