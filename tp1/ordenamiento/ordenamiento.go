package ordenamiento

import (
	"fmt"
	Cola "rerepolez/cola"
	Voto "rerepolez/votos"
	"strconv"
)

func OrdenarPadronesRadix(padrones []Voto.Votante) {
	for i := Voto.DNI_LONGITUD; i > 0; i-- {
		ordenarPorDigitoCounting(padrones, 10, i-1, i)
	}
}

func ordenarPorDigitoCounting(padrones []Voto.Votante, largo int, indiceIni int, indiceFin int) {
	colas := make([]Cola.Cola[Voto.Votante], largo)

	for i := range colas {
		colas[i] = Cola.CrearColaEnlazada[Voto.Votante]()
	}

	for _, padron := range padrones {
		dniCompleto := agregarCeros(padron.LeerDNI())
		digito, _ := strconv.Atoi(dniCompleto[indiceIni:indiceFin])
		colas[digito].Encolar(padron)
	}

	indice := 0
	for _, cola := range colas {
		for !cola.EstaVacia() {
			padrones[indice] = cola.Desencolar()
			indice++
		}
	}
}

func agregarCeros(valor int) string {
	return fmt.Sprintf("%08d", valor)
}
