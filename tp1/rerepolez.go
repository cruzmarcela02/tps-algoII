package main

import (
	"bufio"
	"fmt"
	"os"
	Archivo "rerepolez/archivos"
	Cola "rerepolez/cola"
	Comando "rerepolez/comandos"
	Error "rerepolez/errores"
	Voto "rerepolez/votos"
	"strings"
)

func main() {
	params := os.Args[1:]

	if len(params) != 2 {
		fmt.Println(Error.ErrorParametros{}.Error())
		return
	}

	partidos := Archivo.ObtenerPartidos(params[0])
	padrones := Archivo.ObtenerPadrones(params[1])

	if partidos == nil || padrones == nil {
		fmt.Println(Error.ErrorLeerArchivo{}.Error())
		return
	}

	votantes := Cola.CrearColaEnlazada[Voto.Votante]()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		comando := strings.Split(s.Text(), " ")

		switch comando[0] {
		case "ingresar":
			Comando.Ingresar(comando[1], padrones, votantes)
		case "votar":
			Comando.Votar(comando[1], comando[2], votantes, partidos)
		case "deshacer":
			Comando.Deshacer(votantes)
		case "fin-votar":
			Comando.FinVotar(votantes, partidos)
		}
	}

	Comando.ImprimirResultados(votantes, partidos)
}
