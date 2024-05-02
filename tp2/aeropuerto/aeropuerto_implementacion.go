package aeropuerto

import (
	Cola "algueiza/cola"
	Heap "algueiza/cola_prioridad"
	Conexion "algueiza/conexion"
	Dicc "algueiza/diccionario"
	Fecha "algueiza/fechas"
	Lista "algueiza/lista"
	"strings"

	Vuelo "algueiza/vuelos"
)

type aeropuerto struct {
	vuelos       Dicc.Diccionario[string, Vuelo.Vuelo]
	ascendentes  Dicc.DiccionarioOrdenado[Fecha.Fecha, Lista.Lista[Vuelo.Vuelo]]
	descendentes Dicc.DiccionarioOrdenado[Fecha.Fecha, Lista.Lista[Vuelo.Vuelo]]
	conexiones   Dicc.Diccionario[Conexion.ConexionOD, Conexion.Conexion]
}

func CrearAeropuerto() Aeropuerto {
	ap := new(aeropuerto)
	ap.vuelos = Dicc.CrearHash[string, Vuelo.Vuelo]()
	ap.ascendentes = Dicc.CrearABB[Fecha.Fecha, Lista.Lista[Vuelo.Vuelo]](Fecha.CmpFechas)
	ap.descendentes = Dicc.CrearABB[Fecha.Fecha, Lista.Lista[Vuelo.Vuelo]](Fecha.CmpFechasInversa)
	ap.conexiones = Dicc.CrearHash[Conexion.ConexionOD, Conexion.Conexion]()
	return ap
}

func (ap *aeropuerto) CargarVuelo(vuelo Vuelo.Vuelo) {
	var listaVuelosAsc Lista.Lista[Vuelo.Vuelo]
	var listaVuelosDesc Lista.Lista[Vuelo.Vuelo]

	ap.borrarDatosVuelo(vuelo.VerCodigo())
	ap.vuelos.Guardar(vuelo.VerCodigo(), vuelo)

	if ap.ascendentes.Pertenece(vuelo.VerFecha()) {
		listaVuelosAsc = ap.ascendentes.Obtener(vuelo.VerFecha())
		listaVuelosDesc = ap.descendentes.Obtener(vuelo.VerFecha())
	} else {
		listaVuelosAsc = Lista.CrearListaEnlazada[Vuelo.Vuelo]()
		listaVuelosDesc = Lista.CrearListaEnlazada[Vuelo.Vuelo]()
	}

	actualizarListaVuelos(1, listaVuelosAsc, vuelo, ap.ascendentes)
	actualizarListaVuelos(-1, listaVuelosDesc, vuelo, ap.descendentes)
}

func actualizarListaVuelos(multiplo int, lista Lista.Lista[Vuelo.Vuelo], vuelo Vuelo.Vuelo, diccOrd Dicc.DiccionarioOrdenado[Fecha.Fecha, Lista.Lista[Vuelo.Vuelo]]) {
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		vueloActual := iter.VerActual()
		if strings.Compare(vuelo.VerCodigo(), vueloActual.VerCodigo())*multiplo < 0 {
			iter.Insertar(vuelo)
			diccOrd.Guardar(vuelo.VerFecha(), lista)
			return
		}
	}

	lista.InsertarUltimo(vuelo)
	diccOrd.Guardar(vuelo.VerFecha(), lista)
}

func (ap *aeropuerto) borrarDatosVuelo(codigo string) {
	if ap.vuelos.Pertenece(codigo) {
		vueloViejo := ap.vuelos.Obtener(codigo)
		listaVuelosAsc := ap.ascendentes.Obtener(vueloViejo.VerFecha())
		listaVuelosDesc := ap.descendentes.Obtener(vueloViejo.VerFecha())
		infoConexion := ap.conexiones.Obtener(Conexion.CrearConexion(vueloViejo.VerOrigen(), vueloViejo.VerDestino()))
		infoConexion.BorrarConexion(vueloViejo.VerFecha(), vueloViejo.VerCodigo())
		borrarVuelo(listaVuelosAsc, vueloViejo.VerCodigo())
		borrarVuelo(listaVuelosDesc, vueloViejo.VerCodigo())
	}
}

func borrarVuelo(lista Lista.Lista[Vuelo.Vuelo], codigo string) {
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		vueloActual := iter.VerActual()
		if vueloActual.VerCodigo() == codigo {
			iter.Borrar()
			return
		}
	}
}

func (ap *aeropuerto) ObtenerVuelosPrioritarios(cantidad int) []Vuelo.Vuelo {
	prioritarios := Heap.CrearHeap(Vuelo.CmpVuelos)
	vuelosPrioritarios := make([]Vuelo.Vuelo, 0)

	for iter := ap.vuelos.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vueloActual := iter.VerActual()
		prioritarios.Encolar(vueloActual)
	}

	for !prioritarios.EstaVacia() && cantidad > 0 {
		vuelosPrioritarios = append(vuelosPrioritarios, prioritarios.Desencolar())
		cantidad--
	}

	return vuelosPrioritarios
}

func (ap *aeropuerto) MostrarInformacion(codigo string) string {
	if !ap.vuelos.Pertenece(codigo) {
		return ""
	}

	vuelo := ap.vuelos.Obtener(codigo)
	return vuelo.ObtenerInformacion()
}

func (ap *aeropuerto) Borrar(desde, hasta Fecha.Fecha) Cola.Cola[Vuelo.Vuelo] {
	fechasABorrar := Cola.CrearColaEnlazada[Fecha.Fecha]()
	vuelosBorrados := Cola.CrearColaEnlazada[Vuelo.Vuelo]()
	iter := ap.ascendentes.IteradorRango(&desde, &hasta)

	for iter.HaySiguiente() {
		fecha, listaVuelos := iter.VerActual()
		fechasABorrar.Encolar(fecha)

		for iterLista := listaVuelos.Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
			vueloActual := iterLista.VerActual()
			ap.vuelos.Borrar(vueloActual.VerCodigo())
			vuelosBorrados.Encolar(vueloActual)
			conexion := Conexion.CrearConexion(vueloActual.VerOrigen(), vueloActual.VerDestino())
			infoConexion := ap.conexiones.Obtener(conexion)
			infoConexion.BorrarConexion(vueloActual.VerFecha(), vueloActual.VerCodigo())

			if infoConexion.CantidadVuelos() == 0 {
				ap.conexiones.Borrar(conexion)
			}
		}

		iter.Siguiente()
	}

	for !fechasABorrar.EstaVacia() {
		fechas := fechasABorrar.Desencolar()
		ap.ascendentes.Borrar(fechas)
		ap.descendentes.Borrar(fechas)
	}

	return vuelosBorrados
}

func (ap *aeropuerto) VerTablero(cantidad int, modo string, desde, hasta Fecha.Fecha) Cola.Cola[Vuelo.Vuelo] {
	var iter Dicc.IterDiccionario[Fecha.Fecha, Lista.Lista[Vuelo.Vuelo]]
	vuelosTablero := Cola.CrearColaEnlazada[Vuelo.Vuelo]()

	if modo == ASCENDENTE {
		iter = ap.ascendentes.IteradorRango(&desde, &hasta)
	} else {
		iter = ap.descendentes.IteradorRango(&hasta, &desde)
	}

	for iter.HaySiguiente() {
		_, listaVuelos := iter.VerActual()
		iterLista := listaVuelos.Iterador()

		for iterLista.HaySiguiente() && cantidad > 0 {
			vuelosTablero.Encolar(iterLista.VerActual())
			cantidad--
			iterLista.Siguiente()
		}

		if cantidad == 0 {
			return vuelosTablero
		}

		iter.Siguiente()
	}

	return vuelosTablero
}

func (ap *aeropuerto) SiguienteVuelo(fecha Fecha.Fecha, conexion Conexion.ConexionOD) Vuelo.Vuelo {
	if !ap.conexiones.Pertenece(conexion) {
		return nil
	}

	conexionObtenida := ap.conexiones.Obtener(conexion)

	if conexionObtenida.CantidadVuelos() == 0 {
		return nil
	}

	return conexionObtenida.SiguienteVuelo(fecha)
}

func (ap *aeropuerto) AgregarConexion(conexion Conexion.ConexionOD, fecha Fecha.Fecha, vuelo Vuelo.Vuelo) {
	var infoConexion Conexion.Conexion

	if ap.conexiones.Pertenece(conexion) {
		infoConexion = ap.conexiones.Obtener(conexion)
	} else {
		infoConexion = Conexion.CrearInfoConexion(conexion)
	}

	infoConexion.AgregarConexion(fecha, vuelo)
	ap.conexiones.Guardar(conexion, infoConexion)
}
