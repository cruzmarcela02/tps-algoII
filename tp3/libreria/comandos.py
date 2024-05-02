import libreria.biblioteca as b
import archivos.exportar_archivos as e
import archivos.lectura_de_archivos as a
from tdas.grafo import Grafo

RAPIDO = 'rapido'
TIEMPO = 0
PRECIO = 1

def camino_mas(comandos, camino, ciudades, grafo):
    comandos = comandos.split(',')

    if len(comandos) != 3:
        print('Error: faltan parámetros')
        return
    
    modo, origen, destino = comandos
    if origen not in ciudades or destino not in ciudades:
        print('Error: no hay aeropuertos en las ciudades ingresadas')
        return
    
    if modo == RAPIDO:
        camino_obtenido = camino_mas_barato_rapido(ciudades[origen], ciudades[destino], grafo, TIEMPO)
    else:
        camino_obtenido = camino_mas_barato_rapido(ciudades[origen], ciudades[destino], grafo, PRECIO)
    
    camino.clear()
    for c in camino_obtenido:
        camino.append(c)

    imprimir_camino(camino_obtenido)

def camino_mas_barato_rapido(aeropuertos_origen, aeropuertos_destino, grafo, variante):
    camino_minimo = float('inf')
    camino = []

    for origen in aeropuertos_origen:
        for destino in aeropuertos_destino:
            padres, distancia = b.camino_minimo_dijkstra(grafo, origen, variante, destino)
            
            if distancia[destino] < camino_minimo:
                camino_minimo = distancia[destino]
                camino = b.reconstruir_camino(padres, destino)

    return camino

def camino_escalas(comandos, camino, ciudades, grafo):
    comandos = comandos.split(',')

    if len(comandos) != 2:
        print('Error: faltan parámetros')
        return
    
    origen, destino = comandos
    if origen not in ciudades or destino not in ciudades:
        print('Error: no hay aeropuertos en las ciudades ingresadas')
        return
    
    aeropuertos_origen = ciudades[origen]
    aeropuertos_destinos = ciudades[destino]
    camino_obtenido = camino_menor_escalas(aeropuertos_origen, aeropuertos_destinos, grafo, camino)
    
    camino.clear()
    for c in camino_obtenido:
        camino.append(c)
    
    imprimir_camino(camino_obtenido)

def camino_menor_escalas(aeropuertos_origen, aeropuertos_destinos, grafo, camino = None):
    camino_minimo = float('inf')
    for origen in aeropuertos_origen:
        padre, distancia = b.bfs(grafo, origen)
        
        for destino in aeropuertos_destinos:
            if distancia[destino] < camino_minimo:
                camino = b.reconstruir_camino(padre, destino)
                camino_minimo = distancia[destino]

    return camino

def centralidad(comandos, grafo):
    if len(comandos) != 1 or not comandos[0].isdigit():
        print('Error: faltan parámetros o no ingresó un dígito')
        return
    
    numero = int(comandos[0])
    aeropuertos = b.mas_importantes(grafo)
    imprimir_recorrido(aeropuertos[:numero])

def nueva_aerolinea(comandos, grafo):
    if len(comandos) != 1:
        print('Error: faltan parámetros')
        return
    
    rutas_minimas = b.mst_prim(grafo, PRECIO)
    e.exportar_nueva_aerolinea(comandos[0], rutas_minimas)
    print('OK')

def itinerario(comandos, aeropuertos_ciudad, grafo):
    if len(comandos) != 1:
        print('Error: faltan parámetros')
        return
    
    grafo_itinerario = Grafo(dirigido=True)
    a.cargar_itinerario(comandos[0], grafo_itinerario)
    orden_recorrido = b.topologico_dfs(grafo_itinerario)
    imprimir_recorrido(orden_recorrido)

    for i in range(len(orden_recorrido) - 1):
        aeropuertos_origen = aeropuertos_ciudad[orden_recorrido[i]]
        aeropuertos_destino = aeropuertos_ciudad[orden_recorrido[i + 1]]
        camino = camino_menor_escalas(aeropuertos_origen, aeropuertos_destino, grafo)
        imprimir_camino(camino)

def exportar(comandos, camino, coordenadas):
    if len(comandos) != 1:
        print('Error: faltan parámetros')
        return

    e.exportar_kml(comandos[0], camino, coordenadas)
    imprimir_camino(camino)
    print('OK')

def imprimir_camino(camino):
    print(' -> '.join(camino))

def imprimir_recorrido(camino):
    print(', '.join(camino))
