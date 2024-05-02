package votos

import (
	Error "rerepolez/errores"
	pila "rerepolez/pila"
)

type votanteImplementacion struct {
	dni        int
	voto       Voto
	pila_votos pila.Pila[Voto]
	finalizo   bool
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.pila_votos = pila.CrearPilaDinamica[Voto]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.finalizo {
		return Error.ErrorVotanteFraudulento{Dni: votante.dni}
	}

	if votante.pila_votos.EstaVacia() {
		votante.pila_votos.Apilar(votante.voto)
	}

	if alternativa == LISTA_IMPUGNA {
		votante.voto.Impugnado = true
	}

	votante.voto.VotoPorTipo[int(tipo)] = alternativa
	votante.pila_votos.Apilar(votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() (error, error) {
	if votante.finalizo {
		return Error.ErrorVotanteFraudulento{Dni: votante.dni}, nil
	}

	if votante.pila_votos.EstaVacia() {
		votante.pila_votos.Apilar(votante.voto)
		return nil, Error.ErrorNoHayVotosAnteriores{}
	}

	votoB := votante.pila_votos.Desapilar()
	if votante.pila_votos.EstaVacia() {
		votante.pila_votos.Apilar(votoB)
		return nil, Error.ErrorNoHayVotosAnteriores{}
	} else {
		votante.voto = votante.pila_votos.VerTope()
	}

	return nil, nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.finalizo {
		return votante.voto, Error.ErrorVotanteFraudulento{Dni: votante.dni}
	}

	votante.finalizo = true
	return votante.voto, nil
}
