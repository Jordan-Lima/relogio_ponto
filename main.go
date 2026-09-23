package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"time"
	"strings"
)

type Funcionario struct {
	pis  string
	nome string
	registros []Registro
}

type Registro struct {
	data time.Time
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
		
		dataMarcacao, err := time.Parse("020120061504", registro[10:22])

		if err != nil {
			log.Printf("Erro de parse na linha: %s | Erro: %v", registro, err)
			continue
		}
		
		ponto := Registro{
			data: dataMarcacao,
			pis: registro[22:34],
		}
		
		fmt.Printf("data: %s\npis: %s\n",
			ponto.data,
			ponto.pis)
	}

	usuario_afd, err := os.Open("./usuarios")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo: %v", err)
	}

	defer usuario_afd.Close()
	
	scanner = bufio.NewScanner(usuario_afd)
	for scanner.Scan() {
		usuario := scanner.Text()
		if strings.Contains(usuario, "pis") {
			continue
		}

		infos := strings.Split(usuario, ";")

		if len(infos) >= 2 {
			staff := Funcionario{
				pis: infos[0],
				nome: infos[1],
				
			}
			fmt.Printf("pis: %s\nnome: %s\n", staff.pis, staff.nome)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro crítico durante a leitura do arquivo: %v", err)
	}
}
