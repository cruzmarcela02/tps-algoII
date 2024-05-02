#!/usr/bin/python3
import sys
from tdas.grafo import Grafo
import archivos.lectura_de_archivos as a
import libreria.comandos as comandos

CAMINO_MAS = 'camino_mas'
CAMINO_ESCALAS = 'camino_escalas'
CENTRALIDAD = 'centralidad'
NUEVA_AEROLINEA = 'nueva_aerolinea'
ITINERARIO = 'itinerario'
EXPORTAR = 'exportar_kml'
CERRAR = ''

def main():
    params = sys.argv[1:]

    if len(params) != 2:
        print('Error: faltan parámetros')
        return

    aeropuertos_ciudad = {}
    aeropuertos_coordenadas = {}
    camino = []
    grafo_aeropuerto = Grafo()

    a.cargar_aeropuertos(params[0], aeropuertos_ciudad, aeropuertos_coordenadas, grafo_aeropuerto)
    a.cargar_vuelos(params[1], grafo_aeropuerto)

    for linea in sys.stdin:
        if linea == CERRAR:
            break

        comando = linea.rstrip().split(' ')
        
        if comando[0] == CAMINO_MAS:
            comandos.camino_mas(' '.join(comando[1:]), camino, aeropuertos_ciudad, grafo_aeropuerto)
        elif comando[0] == CAMINO_ESCALAS:
            comandos.camino_escalas(' '.join(comando[1:]), camino, aeropuertos_ciudad, grafo_aeropuerto)
        elif comando[0] == CENTRALIDAD:
            comandos.centralidad(comando[1:], grafo_aeropuerto)
        elif comando[0] == NUEVA_AEROLINEA:
            comandos.nueva_aerolinea(comando[1:], grafo_aeropuerto)
        elif comando[0] == ITINERARIO:
            comandos.itinerario(comando[1:], aeropuertos_ciudad, grafo_aeropuerto)
        elif comando[0] == EXPORTAR:
            comandos.exportar(comando[1:], camino, aeropuertos_coordenadas)
        else:
            print('Error: comando incorrecto')

main()
