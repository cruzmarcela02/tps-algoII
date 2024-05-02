package comandos

import (
	"fmt"
	Busqueda "rerepolez/busqueda"
	Cola "rerepolez/cola"
	Error "rerepolez/errores"
	Recurso "rerepolez/recursos"
	Voto "rerepolez/votos"
	"strconv"
)

func Ingresar(dniIngresado string, padronesHabilitados []Voto.Votante, colaVotantes Cola.Cola[Voto.Votante]) {
	dni, errDni := strconv.Atoi(dniIngresado)
	if len(dniIngresado) > Voto.DNI_LONGITUD || errDni != nil || dni <= 0 {
		fmt.Println(Error.DNIError{}.Error())
		return
	}

	posicion := Busqueda.BuscarVotante(dni, padronesHabilitados)
	if posicion < 0 {
		fmt.Println(Error.DNIFueraPadron{}.Error())
		return
	}

	colaVotantes.Encolar(padronesHabilitados[posicion])
	fmt.Println("OK")
}

func Votar(categoria string, nroLista string, colaVotantes Cola.Cola[Voto.Votante], partidos []Voto.Partido) {
	if colaVotantes.EstaVacia() {
		fmt.Println(Error.FilaVacia{}.Error())
		return
	}

	tipo, errTipoVoto := Recurso.ConvertirATipoVoto(categoria)
	if errTipoVoto != nil {
		fmt.Println(errTipoVoto.Error())
		return
	}

	nroPartido, errAlternativa := strconv.Atoi(nroLista)
	if errAlternativa != nil || nroPartido < 0 || nroPartido >= len(partidos) {
		fmt.Println(Error.ErrorAlternativaInvalida{}.Error())
		return
	}

	votante := colaVotantes.VerPrimero()
	errFraudulento := votante.Votar(tipo, nroPartido)
	if errFraudulento != nil {
		colaVotantes.Desencolar()
		fmt.Println(errFraudulento.Error())
		return
	}

	fmt.Println("OK")
}

func Deshacer(colaVotantes Cola.Cola[Voto.Votante]) {
	if colaVotantes.EstaVacia() {
		fmt.Println(Error.FilaVacia{}.Error())
		return
	}

	votante := colaVotantes.VerPrimero()
	errFraudulento, errNoDeshacer := votante.Deshacer()

	switch {
	case errFraudulento != nil:
		colaVotantes.Desencolar()
		fmt.Println(errFraudulento.Error())
		return
	case errNoDeshacer != nil:
		fmt.Println(errNoDeshacer.Error())
		return
	}

	fmt.Println("OK")
}

func FinVotar(colaVotantes Cola.Cola[Voto.Votante], partidos []Voto.Partido) {
	if colaVotantes.EstaVacia() {
		fmt.Println(Error.FilaVacia{}.Error())
		return
	}

	votante := colaVotantes.Desencolar()
	votoFinal, errFraudulento := votante.FinVoto()

	if errFraudulento != nil {
		fmt.Println(errFraudulento.Error())
		return
	}

	if votoFinal.Impugnado {
		partidos[0].VotadoPara(Voto.IMPUGNADO)
	} else {
		for categoria, nroLista := range votoFinal.VotoPorTipo {
			partido := partidos[nroLista]
			partido.VotadoPara(Voto.TipoVoto(categoria))
		}
	}

	fmt.Println("OK")
}

func ImprimirResultados(colaVotantes Cola.Cola[Voto.Votante], partidos []Voto.Partido) {
	if !colaVotantes.EstaVacia() {
		fmt.Println(Error.ErrorCiudadanosSinVotar{}.Error())
	}

	for i := 0; i < int(Voto.CANT_VOTACION); i++ {
		categoria := Recurso.ConvertirACadena(Voto.TipoVoto(i))
		fmt.Printf("%s:\n", categoria)
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(Voto.TipoVoto(i)))
		}
		fmt.Println()
	}

	fmt.Println(partidos[0].ObtenerResultado(Voto.IMPUGNADO))
}
