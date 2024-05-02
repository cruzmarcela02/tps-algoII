package recursos

import (
	Error "rerepolez/errores"
	Voto "rerepolez/votos"
)

const (
	PRESIDENTE = "Presidente"
	GOBERNADOR = "Gobernador"
	INTENDENTE = "Intendente"
)

func ConvertirATipoVoto(categoria string) (Voto.TipoVoto, error) {
	var votoA Voto.TipoVoto

	switch categoria {
	case PRESIDENTE:
		votoA = Voto.PRESIDENTE
	case GOBERNADOR:
		votoA = Voto.GOBERNADOR
	case INTENDENTE:
		votoA = Voto.INTENDENTE
	default:
		return votoA, Error.ErrorTipoVoto{}
	}

	return votoA, nil
}

func ConvertirACadena(categoria Voto.TipoVoto) string {
	var cadena string

	switch categoria {
	case Voto.PRESIDENTE:
		cadena = PRESIDENTE
	case Voto.GOBERNADOR:
		cadena = GOBERNADOR
	case Voto.INTENDENTE:
		cadena = INTENDENTE
	}

	return cadena
}
