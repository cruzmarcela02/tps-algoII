package votos

import "fmt"

type partidoImplementacion struct {
	nombre     string
	votos      [CANT_VOTACION]int
	candidatos [CANT_VOTACION]string
}

type partidoEnBlanco struct {
	votos           [CANT_VOTACION]int
	votosImpugnados int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.votos[tipo] == 1 {
		return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], partido.votos[tipo], VOTO)
	}

	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], partido.votos[tipo], VOTOS)
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	if tipo == IMPUGNADO {
		blanco.votosImpugnados++
	} else {
		blanco.votos[tipo]++
	}
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if tipo == IMPUGNADO {
		return blanco.ObtenerResultadoImpugnados()
	}

	if blanco.votos[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d %s", blanco.votos[tipo], VOTO)
	}

	return fmt.Sprintf("Votos en Blanco: %d %s", blanco.votos[tipo], VOTOS)
}

func (blanco partidoEnBlanco) ObtenerResultadoImpugnados() string {
	if blanco.votosImpugnados == 1 {
		return fmt.Sprintf("Votos Impugnados: %d %s", blanco.votosImpugnados, VOTO)
	}

	return fmt.Sprintf("Votos Impugnados: %d %s", blanco.votosImpugnados, VOTOS)
}
