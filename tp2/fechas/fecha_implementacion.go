package fechas

import (
	"strconv"
	"strings"
)

type Fecha struct {
	anio     int
	mes      int
	dia      int
	hora     int
	minutos  int
	segundos int
	fechastr string
}

func CrearFecha(fechaCadena string) Fecha {
	listaFecha := strings.Split(fechaCadena, "T")
	fechaAMD := strings.Split(listaFecha[0], "-")
	horarioHMS := strings.Split(listaFecha[1], ":")

	anio, _ := strconv.Atoi(fechaAMD[0])
	mes, _ := strconv.Atoi(fechaAMD[1])
	dia, _ := strconv.Atoi(fechaAMD[2])
	hora, _ := strconv.Atoi(horarioHMS[0])
	minutos, _ := strconv.Atoi(horarioHMS[1])
	segundos, _ := strconv.Atoi(horarioHMS[2])

	fecha := new(Fecha)
	fecha.anio = anio
	fecha.mes = mes
	fecha.dia = dia
	fecha.hora = hora
	fecha.minutos = minutos
	fecha.segundos = segundos
	fecha.fechastr = fechaCadena
	return *fecha
}

func (f Fecha) ObtenerFecha() string {
	return f.fechastr
}

func (f Fecha) CompararFechas(fecha2 Fecha) int {
	if f.anio-fecha2.anio != 0 {
		return f.anio - fecha2.anio
	}

	if f.mes-fecha2.mes != 0 {
		return f.mes - fecha2.mes
	}

	if f.dia-fecha2.dia != 0 {
		return f.dia - fecha2.dia
	}

	if f.hora-fecha2.hora != 0 {
		return f.hora - fecha2.hora
	}

	if f.minutos-fecha2.minutos != 0 {
		return f.minutos - fecha2.minutos
	}

	return f.segundos - fecha2.segundos
}

func CmpFechas(a, b Fecha) int {
	return a.CompararFechas(b)
}

func CmpFechasInversa(a, b Fecha) int {
	return b.CompararFechas(a)
}
