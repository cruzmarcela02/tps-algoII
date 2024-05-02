package vuelos

import (
	Fecha "algueiza/fechas"
)

type Vuelo interface {

	// Actualizar Datos, actualiza los datos del Vuelo ingresado por archivo
	ActualizarDatos(string, string, string, string, Fecha.Fecha, int, int, int, int)

	// ObtenerInformacion devuelve una cadena con toda la informacion del Vuelo
	ObtenerInformacion() string

	// Nos da la prioridad del Vuelo
	VerPrioridad() int

	// Nos da el codigo del Vuelo
	VerCodigo() string

	// Nos da la fecha del Vuelo
	VerFecha() Fecha.Fecha

	// Nos da el origen del Vuelo
	VerOrigen() string

	// Nos da el destino del Vuelo
	VerDestino() string
}
