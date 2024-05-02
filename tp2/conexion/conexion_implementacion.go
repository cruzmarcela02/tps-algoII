package conexion

import (
	Dicc "algueiza/diccionario"
	Fecha "algueiza/fechas"
	Vuelo "algueiza/vuelos"
)

type InfoConexion struct {
	conexion       ConexionOD
	vuelosPorFecha Dicc.DiccionarioOrdenado[Fecha.Fecha, Dicc.Diccionario[string, Vuelo.Vuelo]]
	cantVuelos     int
}

type ConexionOD struct {
	origen  string
	destino string
}

func CrearConexion(origen, destino string) ConexionOD {
	conexion := new(ConexionOD)
	conexion.origen = origen
	conexion.destino = destino
	return *conexion
}

func CrearInfoConexion(conexion ConexionOD) Conexion {
	infoConexion := new(InfoConexion)
	infoConexion.conexion = conexion
	infoConexion.vuelosPorFecha = Dicc.CrearABB[Fecha.Fecha, Dicc.Diccionario[string, Vuelo.Vuelo]](Fecha.CmpFechas)
	return infoConexion
}

func (infoConexion *InfoConexion) AgregarConexion(fecha Fecha.Fecha, vuelo Vuelo.Vuelo) {
	var dicc Dicc.Diccionario[string, Vuelo.Vuelo]

	if infoConexion.vuelosPorFecha.Pertenece(fecha) {
		dicc = infoConexion.vuelosPorFecha.Obtener(fecha)
	} else {
		dicc = Dicc.CrearHash[string, Vuelo.Vuelo]()
	}

	dicc.Guardar(vuelo.VerCodigo(), vuelo)
	infoConexion.vuelosPorFecha.Guardar(fecha, dicc)
	infoConexion.cantVuelos++
}

func (infoConexion *InfoConexion) SiguienteVuelo(fecha Fecha.Fecha) Vuelo.Vuelo {
	iterCon := infoConexion.vuelosPorFecha.IteradorRango(&fecha, nil)

	if infoConexion.vuelosPorFecha.Pertenece(fecha) {
		iterCon.Siguiente()
	}

	if iterCon.HaySiguiente() {
		_, diccionario := iterCon.VerActual()
		for iterDicc := diccionario.Iterador(); iterDicc.HaySiguiente(); iterDicc.Siguiente() {
			_, vuelo := iterDicc.VerActual()
			return vuelo
		}
	}

	return nil
}

func (infoConexion *InfoConexion) BorrarConexion(fecha Fecha.Fecha, codigo string) {
	dicc := infoConexion.vuelosPorFecha.Obtener(fecha)
	if dicc.Pertenece(codigo) {
		dicc.Borrar(codigo)
		infoConexion.cantVuelos--
	}

	if dicc.Cantidad() == 0 {
		infoConexion.vuelosPorFecha.Borrar(fecha)
	}
}

func (infoConexion *InfoConexion) CantidadVuelos() int {
	return infoConexion.cantVuelos
}
