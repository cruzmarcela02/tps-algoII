package archivos

import (
	"bufio"
	"os"
	Ordenar "rerepolez/ordenamiento"
	Voto "rerepolez/votos"
	"strconv"
	"strings"
)

func ObtenerPartidos(ruta string) []Voto.Partido {
	lista_partidos := make([]Voto.Partido, 0)

	archivo, err := os.Open(ruta)
	if err != nil {
		return nil
	}

	defer archivo.Close()

	lista_partidos = append(lista_partidos, Voto.CrearVotosEnBlanco())
	linea := bufio.NewScanner(archivo)
	for linea.Scan() {
		candidatos := strings.Split(linea.Text(), ",")
		partido := Voto.CrearPartido(candidatos[0], [3]string{candidatos[1], candidatos[2], candidatos[3]})
		lista_partidos = append(lista_partidos, partido)
	}

	return lista_partidos
}

func ObtenerPadrones(ruta string) []Voto.Votante {
	lista_padrones := make([]Voto.Votante, 0)

	archivo, err := os.Open(ruta)
	if err != nil {
		return nil
	}

	defer archivo.Close()

	linea := bufio.NewScanner(archivo)
	for linea.Scan() {
		dni, _ := strconv.Atoi(linea.Text())
		votante := Voto.CrearVotante(dni)
		lista_padrones = append(lista_padrones, votante)
	}

	Ordenar.OrdenarPadronesRadix(lista_padrones)

	return lista_padrones
}
