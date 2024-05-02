package archivos

import (
	Aeropuerto "algueiza/aeropuerto"
	Conexion "algueiza/conexion"
	Fecha "algueiza/fechas"
	Vuelo "algueiza/vuelos"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CargarVuelo(ruta string, aeropuerto Aeropuerto.Aeropuerto) string {
	archivo, err := os.Open(ruta)
	if err != nil {
		return ""
	}

	defer archivo.Close()

	linea := bufio.NewScanner(archivo)
	for linea.Scan() {
		vuelo := strings.Split(linea.Text(), ",")
		fecha := Fecha.CrearFecha(vuelo[6])
		prioridad, _ := strconv.Atoi(vuelo[5])
		demora, _ := strconv.Atoi(vuelo[7])
		duracion, _ := strconv.Atoi(vuelo[8])
		cancelado, _ := strconv.Atoi(vuelo[9])
		vueloCreado := Vuelo.CrearVuelo(vuelo[0], vuelo[1], vuelo[2], vuelo[3], vuelo[4], fecha, prioridad, demora, duracion, cancelado)
		conexion := Conexion.CrearConexion(vuelo[2], vuelo[3])
		aeropuerto.CargarVuelo(vueloCreado)
		aeropuerto.AgregarConexion(conexion, fecha, vueloCreado)
	}

	return "OK"
}
