import random

class Grafo:
    def __init__(self, dirigido = False):
        '''Inicializa una nuevo grafo, vacío.'''
        self.dirigido = dirigido
        self.vertices = {}
    
    def agregar_vertice(self, v):
        '''Agrega un nuevo vértice al grafo.'''
        if v in self.vertices:
            return
        
        self.vertices[v] = {}

    def borrar_vertice(self, v):
        '''Borra un vértice del grafo.'''
        if v not in self.vertices:
            return
        
        self.vertices.pop(v)
        for clave, adyacentes in self.vertices.items():
            if v not in adyacentes:
                continue

            self.vertices[clave].pop(v)

    def agregar_arista(self, v, w, peso = 1):
        '''Agrega una nueva arista al grafo entre vértices recibidos.'''
        if v not in self.vertices:
            return
        
        self.vertices[v][w] = peso

        if not self.dirigido:
            self.vertices[w][v] = peso

    def borrar_arista(self, v, w):
        '''Borra la arista del grafo entre vértices recibidos.'''
        if v not in self.vertices or w not in self.vertices[v]:
            return
        
        self.vertices[v].pop(w)

        if not self.dirigido:
            self.vertices[w].pop(v)

    def estan_unidos(self, v, w):
        '''Devuelve True o False según si existe una arista entre los vértices recibidos.'''
        return w in self.vertices[v]
    
    def peso_arista(self, v, w):
        '''Devuelve el peso de la arista que hay entre los vértices recibidos.'''
        if v not in self.vertices or w not in self.vertices[v]:
            return None
        
        return self.vertices[v][w]

    def obtener_vertices(self):
        '''Devuelve una lista con todos los vértices del grafo.'''
        return list(self.vertices)
    
    def vertice_aleatorio(self):
        '''Devuelve un vértice aleatorio del grafo.'''
        if len(self.vertices) == 0: 
            return
        
        return list(self.vertices)[random.randint(0, len(self.vertices)-1)]
    
    def adyacentes(self, v):
        '''Devuelve una lista con los vértices adyacentes al vértice recibido.'''
        if v not in self.vertices: 
            return
        
        return list(self.vertices[v])
