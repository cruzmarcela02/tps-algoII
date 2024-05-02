def cargar_aeropuertos(ruta, aeropuertos_ciudad, aeropuertos_latitud, grafo):
    with open(ruta) as f:
        for linea in f:
            ciudad, codigo, latitud, longitud = linea.rstrip('\n').split(',')         
            aeropuertos_ciudad[ciudad] = aeropuertos_ciudad.get(ciudad, []) + [codigo]
            aeropuertos_latitud[codigo] = (float(latitud), float(longitud))
            grafo.agregar_vertice(codigo)

def cargar_vuelos(ruta, grafo):
    with open(ruta) as f:
        for linea in f:
            origen, destino, tiempo, precio, cant_vuelos = linea.rstrip('\n').split(',')
            grafo.agregar_arista(origen, destino, (int(tiempo), int(precio), int(cant_vuelos)))

def cargar_itinerario(ruta, grafo):
    with open(ruta) as f:
        for linea in f:
            ciudades = linea.rstrip('\n').split(',') 
            if len(ciudades) != 2:
                for ciudad in ciudades:
                    grafo.agregar_vertice(ciudad)
            else:
                grafo.agregar_arista(ciudades[0], ciudades[1])
