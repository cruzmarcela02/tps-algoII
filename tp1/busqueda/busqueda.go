package busqueda

import Voto "rerepolez/votos"

func BuscarVotante(dni int, padronesHabilitados []Voto.Votante) int {
	return busquedaBinaria(padronesHabilitados, 0, len(padronesHabilitados), dni)
}

func busquedaBinaria(padronesHabilitados []Voto.Votante, inicio, fin, dni int) int {
	if inicio > fin {
		return -1
	}

	medio := (inicio + fin) / 2
	if padronesHabilitados[medio].LeerDNI() == dni {
		return medio
	}

	if padronesHabilitados[medio].LeerDNI() > dni {
		return busquedaBinaria(padronesHabilitados, inicio, medio-1, dni)
	} else {
		return busquedaBinaria(padronesHabilitados, medio+1, fin, dni)
	}
}
