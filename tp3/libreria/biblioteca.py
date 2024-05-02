from tdas.cola import Cola
from tdas.pila import Pila
import heapq

CANT_VUELOS = 2

def bfs(grafo, origen):
    visitados = set()
    padres, distancia = {}, {}
    q = Cola()
    padres[origen] = None
    distancia[origen] = 0
    visitados.add(origen)
    q.encolar(origen)
    
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                distancia[w] = distancia[v] + 1
                visitados.add(w)
                q.encolar(w)

    return padres, distancia

def reconstruir_camino(padres, destino):
    camino = []
    
    while destino is not None:
        camino.append(destino)
        destino = padres[destino]

    camino.reverse()
    return camino

def camino_minimo_dijkstra(grafo, origen, variante, destino = None):
    # VARIANTE: (tiempo, precio, cant_vuelos)
    distancias, padres = {}, {}

    for v in grafo.obtener_vertices():
        distancias[v] = float('inf')

    distancias[origen] = 0
    padres[origen] = None
    heap = []
    heapq.heappush(heap, (0, origen))

    while heap:
        _, v = heapq.heappop(heap)
        
        if v == destino:
            return padres, distancias
        
        for w in grafo.adyacentes(v):
            if variante == CANT_VUELOS:
                frecuencia = 1 / grafo.peso_arista(v,w)[variante]
                dist_actual = distancias[v] + frecuencia
            else:
                dist_actual = distancias[v] + grafo.peso_arista(v, w)[variante]
            
            if dist_actual < distancias[w]:
                distancias[w] = dist_actual
                padres[w] = v
                heapq.heappush(heap, (distancias[w], w))

    return padres, distancias

def topologico_dfs(grafo):
    pila = Pila()
    visitados = set()

    for v in grafo.obtener_vertices():
        if v not in visitados:
            visitados.add(v)
            dfs(grafo, v, visitados, pila)

    return pila_a_lista(pila)

def dfs(grafo, v, visitados, pila):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            dfs(grafo, w, visitados, pila)
    
    pila.apilar(v)

def pila_a_lista(pila):
    res = []
    while not pila.esta_vacia():
        res.append(pila.desapilar())

    return res

def mst_prim(grafo):
    v = grafo.vertice_aleatorio()
    visitados = set()
    arbol = []
    heap = []
    suma = 0

    visitados.add(v)
    for w in grafo.adyacentes(v):
        peso = grafo.peso_arista(v, w)
        heapq.heappush(heap, (peso, v, w, peso))

    while heap:
        _, v, w, peso = heapq.heappop(heap)
        if w in visitados:
            continue

        visitados.add(w)
        arbol.append((v, w, peso))
        suma += peso
        
        for x in grafo.adyacentes(w):
            if x not in visitados:
                peso = grafo.peso_arista(w, x)
                heapq.heappush(heap, (peso, w, x, peso))

    return arbol, suma

def centralidad(grafo):
    cent = {}
    for v in grafo.obtener_vertices():
        cent[v] = 0
   
    for v in grafo.obtener_vertices():
        padres, distancias = camino_minimo_dijkstra(grafo, v, CANT_VUELOS)
        cent_aux = {}
        for w in grafo.obtener_vertices():
            cent_aux[w] = 0
        # Aca filtramos (de ser necesario) los vertices a distancia infinita,
        # y ordenamos de mayor a menor
        vertices_ordenados = ordenar_vertices(distancias)
        for w in vertices_ordenados:
            if v == w or padres[w] is None:
                continue
            cent_aux[padres[w]] += 1 + cent_aux[w]
        # Le sumamos 1 a la centralidad de todos los vertices que se encuentren en
        # el medio del camino
        for w in grafo.obtener_vertices():
            if w == v:
                continue
            cent[w] += cent_aux[w]

    return cent

def mas_importantes(grafo):
    aeropuertos = centralidad(grafo)
    return sorted(aeropuertos, key=lambda v: aeropuertos[v], reverse=True)

def ordenar_vertices(distancia):
    for codigo_ap in distancia:
        if distancia[codigo_ap] == float('inf'):
            distancia.pop(codigo_ap)
    
    return sorted(distancia, key=lambda v: distancia[v], reverse=True)
