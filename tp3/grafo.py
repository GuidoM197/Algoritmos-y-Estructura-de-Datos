import random


class Grafo:
    def __init__(self, es_dirigido=False, vertices_init=[]):
        self.es_dirigido = es_dirigido
        self.vertices = {}
        for vertice in vertices_init:
            self.agregar_vertice(vertice)

    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}

    def borrar_vertice(self, v):
        if v in self.vertices:
            del self.vertices[v]
            for vert in self.vertices:
                if v in self.vertices[vert]:
                    del self.vertices[vert][v]

    def agregar_arista(self, v, w, peso=1):
        if v not in self.vertices:
            self.agregar_vertice(v)
        if w not in self.vertices:
            self.agregar_vertice(w)
        self.vertices[v][w] = peso
        if not self.es_dirigido:
            self.vertices[w][v] = peso

    def borrar_arista(self, v, w):
        if v in self.vertices and w in self.vertices[v]:
            del self.vertices[v][w]
        if not self.es_dirigido and w in self.vertices and v in self.vertices[w]:
            del self.vertices[w][v]

    def estan_unidos(self, v, w):
        return v in self.vertices and w in self.vertices[v]

    def peso_arista(self, v, w):
        if self.estan_unidos(v, w):
            return self.vertices[v][w]
        return None

    def obtener_vertices(self):
        return list(self.vertices.keys())

    def vertice_aleatorio(self):
        if not self.vertices:
            return None
        return random.choice(list(self.vertices.keys()))

    def adyacentes(self, v):
        if v in self.vertices:
            return list(self.vertices[v].keys())
        return []

    def __str__(self):
        grafo_str = ""
        for vertice in self.vertices:
            for adyacente in self.vertices[vertice]:
                grafo_str += f"{vertice} --({self.vertices[vertice][adyacente]})--> {adyacente}\n"
        return grafo_str