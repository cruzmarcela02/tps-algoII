package aeropuerto

import (
	Cola "algueiza/cola"

	Conexion "algueiza/conexion"
	Fecha "algueiza/fechas"
	Vuelo "algueiza/vuelos"
)

const (
	ASCENDENTE  = "asc"
	DESCENDENTE = "desc"
)

type Aeropuerto interface {

	// CargarVuelo, carga los datos de un vuelo en el Diccionario de Vuelos
	CargarVuelo(Vuelo.Vuelo)

	// ObtenerVuelosPrioritarios retorna un slice con los n vuelos con mayor prioridad
	ObtenerVuelosPrioritarios(n int) []Vuelo.Vuelo

	// Muestra la informacion del vuelo indicado por el codigo ingresado por parámetro
	MostrarInformacion(codigo string) string

	// Borra los vuelos realizados entra las fechas recividas (e incluidas) del aeropuerto
	Borrar(desde, hasta Fecha.Fecha) Cola.Cola[Vuelo.Vuelo]

	// Muestra los vuelos realizados entre las fechas recividas(e incluidas) del aeropuerto
	VerTablero(cantidad int, modo string, desde, hasta Fecha.Fecha) Cola.Cola[Vuelo.Vuelo]

	// SiguienteVuelo retorna el siguiente Vuelo con origen, destino y fecha ingresados por parámetro
	SiguienteVuelo(fecha Fecha.Fecha, conexion Conexion.ConexionOD) Vuelo.Vuelo

	// AgregarConexion agrega una conexion con el origen, destino y fecha de un vuelo
	AgregarConexion(conexion Conexion.ConexionOD, fecha Fecha.Fecha, vuelo Vuelo.Vuelo)
}
