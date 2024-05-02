package vuelos

import (
	Fecha "algueiza/fechas"
	"fmt"
	"strings"
)

type vuelo struct {
	codigo    string
	aerolinea string
	origen    string
	destino   string
	numero    string
	prioridad int
	fecha     Fecha.Fecha
	demora    int
	duracion  int
	cancelado int
}

func CrearVuelo(codigo, aerolinea, origen, destino, numero string, fecha Fecha.Fecha,
	prioridad, demora, duracion, cancelado int) Vuelo {
	vuelo := new(vuelo)
	vuelo.codigo = codigo
	vuelo.ActualizarDatos(aerolinea, origen, destino, numero, fecha, prioridad, demora, duracion, cancelado)
	return vuelo
}

func (v *vuelo) ActualizarDatos(aerolinea, origen, destino, numero string, fecha Fecha.Fecha,
	prioridad, demora, duracion, cancelado int) {
	v.aerolinea = aerolinea
	v.origen = origen
	v.destino = destino
	v.numero = numero
	v.fecha = fecha
	v.demora = demora
	v.duracion = duracion
	v.cancelado = cancelado
	v.prioridad = prioridad
}

func (v *vuelo) ObtenerInformacion() string {
	return fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d", v.codigo, v.aerolinea, v.origen, v.destino, v.numero, v.prioridad, v.fecha.ObtenerFecha(), v.demora, v.duracion, v.cancelado)
}

func (v *vuelo) VerPrioridad() int {
	return v.prioridad
}

func (v *vuelo) VerCodigo() string {
	return v.codigo
}

func (v *vuelo) VerFecha() Fecha.Fecha {
	return v.fecha
}

func (v *vuelo) VerOrigen() string {
	return v.origen
}

func (v *vuelo) VerDestino() string {
	return v.destino
}

func CmpCodigos(a, b Vuelo) int {
	return strings.Compare(b.VerCodigo(), a.VerCodigo())
}

func CmpVuelos(a, b Vuelo) int {
	comparacion := a.VerPrioridad() - b.VerPrioridad()

	if comparacion != 0 {
		return comparacion
	}

	return CmpCodigos(a, b)
}
