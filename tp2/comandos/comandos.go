package comandos

import (
	Aeropuerto "algueiza/aeropuerto"
	Conexion "algueiza/conexion"
	Fecha "algueiza/fechas"
	"fmt"

	Archivo "algueiza/archivos"
	Error "algueiza/errores"
	"os"
	"strconv"
)

func AgregarArchivo(ruta string, aeropuerto Aeropuerto.Aeropuerto) {
	respuesta := Archivo.CargarVuelo(ruta, aeropuerto)

	if respuesta != "" {
		fmt.Println("OK")
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: "agregar_archivo"}.Error())
	}
}

func VerTablero(comandos []string, aeropuerto Aeropuerto.Aeropuerto) {
	if len(comandos) != 5 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comandos[0]}.Error())
		return
	}

	cantidad, err := strconv.Atoi(comandos[1])
	if err != nil || cantidad <= 0 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comandos[0]}.Error())
		return
	}

	modo := comandos[2]
	if modo != Aeropuerto.ASCENDENTE && modo != Aeropuerto.DESCENDENTE {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comandos[0]}.Error())
		return
	}

	fDesde := Fecha.CrearFecha(comandos[3])
	fHasta := Fecha.CrearFecha(comandos[4])

	if fDesde.CompararFechas(fHasta) > 0 {
		fmt.Println("OK")
		return
	}

	vuelosTablero := aeropuerto.VerTablero(cantidad, modo, fDesde, fHasta)

	for !vuelosTablero.EstaVacia() {
		vuelo := vuelosTablero.Desencolar()
		fmt.Printf("%s - %s\n", vuelo.VerFecha().ObtenerFecha(), vuelo.VerCodigo())
	}

	fmt.Println("OK")
}

func InfoVuelo(codigo string, aeropuerto Aeropuerto.Aeropuerto) {
	respuesta := aeropuerto.MostrarInformacion(codigo)

	if respuesta != "" {
		fmt.Println(respuesta)
		fmt.Println("OK")
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: "info_vuelo"}.Error())
	}
}

func PrioridadVuelos(cantidad string, aeropuerto Aeropuerto.Aeropuerto) {
	cantidadInt, err := strconv.Atoi(cantidad)

	if err != nil || cantidadInt <= 0 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: "prioridad_vuelos"}.Error())
		return
	}

	prioritarios := aeropuerto.ObtenerVuelosPrioritarios(cantidadInt)

	for _, vuelo := range prioritarios {
		fmt.Printf("%d - %s\n", vuelo.VerPrioridad(), vuelo.VerCodigo())
	}

	fmt.Println("OK")
}

func SiguienteVuelo(comandos []string, aeropuerto Aeropuerto.Aeropuerto) {
	if len(comandos) != 4 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comandos[0]}.Error())
		return
	}

	conexion := Conexion.CrearConexion(comandos[1], comandos[2])
	fecha := Fecha.CrearFecha(comandos[3])
	vueloSiguiente := aeropuerto.SiguienteVuelo(fecha, conexion)

	if vueloSiguiente != nil {
		fmt.Println(vueloSiguiente.ObtenerInformacion())
	} else {
		fmt.Println(Error.ErrorSiguienteVuelo{Origen: comandos[1], Destino: comandos[2], Fecha: comandos[3]})
	}

	fmt.Println("OK")
}

func Borrar(comandos []string, aeropuerto Aeropuerto.Aeropuerto) {
	if len(comandos) != 3 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comandos[0]}.Error())
		return
	}

	fDesde := Fecha.CrearFecha(comandos[1])
	fHasta := Fecha.CrearFecha(comandos[2])

	if fDesde.CompararFechas(fHasta) > 0 {
		fmt.Println("OK")
		return
	}

	vuelosBorrados := aeropuerto.Borrar(fDesde, fHasta)
	for !vuelosBorrados.EstaVacia() {
		vuelo := vuelosBorrados.Desencolar()
		fmt.Println(vuelo.ObtenerInformacion())
	}

	fmt.Println("OK")
}

func ComandoDefault(comando string) {
	fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorComando{Comando: comando}.Error())
}
