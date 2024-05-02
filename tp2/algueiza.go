package main

import (
	Aeropuerto "algueiza/aeropuerto"
	Comando "algueiza/comandos"
	Error "algueiza/errores"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	aeropuerto := Aeropuerto.CrearAeropuerto()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		comandos := strings.Split(s.Text(), " ")

		if len(comandos) == 1 {
			fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comandos[0]}.Error())
			continue
		}

		switch comandos[0] {
		case "agregar_archivo":
			Comando.AgregarArchivo(comandos[1], aeropuerto)
		case "ver_tablero":
			Comando.VerTablero(comandos, aeropuerto)
		case "info_vuelo":
			Comando.InfoVuelo(comandos[1], aeropuerto)
		case "prioridad_vuelos":
			Comando.PrioridadVuelos(comandos[1], aeropuerto)
		case "siguiente_vuelo":
			Comando.SiguienteVuelo(comandos, aeropuerto)
		case "borrar":
			Comando.Borrar(comandos, aeropuerto)
		default:
			Comando.ComandoDefault(comandos[0])
		}
	}

}
