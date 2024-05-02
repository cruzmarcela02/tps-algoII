from tdas.grafo import Grafo
import libreria.biblioteca as b

def obtener_fila(grafo): # otra opcion es usar un heap(?)
    res = []
    visitados = vertice_equipo(grafo) # diccionario con todos los vertices y su equipo
    random = grafo.vertice_aleatorio()
    res.append(random)

    for v in grafo.obtener_vertices():
        if visitados[res[-1:]] != visitados[v]: # chequeo si el ultimo elemento de la lista es distinto que el vertice que estoy analizando
            res.append(v)
        else:
            pass # insertarlo en otra parte si se puede? caso contrario retorno lista vacia

            
    
def vertice_equipo(grafo):
    visitados = {}
    cont = 0
    for v in grafo.obtener_vertices():
        if v not in visitados:
            visitados[v] = cont
            dfs(grafo, v, visitados, cont)
            cont+=1

    return visitados

def dfs(grafo, v, visitados, cont):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados[w] = cont
            dfs(grafo, w, visitados, cont)


def main():
    grafo = Grafo()

    grafo.agregar_vertice('A')
    grafo.agregar_vertice('B')
    grafo.agregar_vertice('C')
    grafo.agregar_vertice('D')
    grafo.agregar_vertice('E')
    grafo.agregar_vertice('F')
    grafo.agregar_vertice('G')

    grafo.agregar_arista('A', 'B', 3)
    grafo.agregar_arista('B', 'D', 3)
    grafo.agregar_arista('B', 'E', 3)
    grafo.agregar_arista('F', 'D', 2)
    grafo.agregar_arista('C', 'D', 2)
    grafo.agregar_arista('G', 'D', 1)
    grafo.agregar_arista('A', 'C', 4)
    grafo.agregar_arista('D', 'E', 4)
    grafo.agregar_arista('C', 'B', 5)
    grafo.agregar_arista('C', 'G', 6)
    grafo.agregar_arista('E', 'F', 6)

    

    
main()