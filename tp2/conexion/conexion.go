package conexion

import (
	Fecha "algueiza/fechas"
	Vuelo "algueiza/vuelos"
)

type Conexion interface {
	// Devuelve el proximo vuelo a la fecha, en caso de no existir retorna nil
	SiguienteVuelo(fecha Fecha.Fecha) Vuelo.Vuelo

	// AgregarConexion agrega un vuelo en la fecha indicada en Infoconexion
	AgregarConexion(fecha Fecha.Fecha, vuelo Vuelo.Vuelo)

	// BorrarConexion borra un vuelo de la InfoConexion si es que el codigo pertenece.
	BorrarConexion(fecha Fecha.Fecha, codigo string)

	// Devuelde la cantidad de vuelos dentro de Infoconexion
	CantidadVuelos() int
}
