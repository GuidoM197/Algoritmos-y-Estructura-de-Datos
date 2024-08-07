from collections import deque, Counter
from grafo import Grafo
from random import choice
from pila import Pila

def bfs(grafo, origen, destino=None):
    visitados = set()
    orden = {origen: 0}
    padres = {origen: None}
    cola = deque()
    cola.append(origen)
    visitados.add(origen)

    while cola:
        actual = cola.popleft()
        if destino is not None and actual == destino:
            return padres, orden

        for ady in grafo.adyacentes(actual):
            if ady not in visitados:
                cola.append(ady)
                orden[ady] = orden[actual] + 1
                padres[ady] = actual
                visitados.add(ady)

    return padres, orden


def armar_ciclo(padres, ini, fin):
    v = fin
    res = []
    res.append(ini)
    while v != ini:
        res.append(v)
        v = padres[v]
    res.append(ini)

    return res[::-1]

def bfs_encontrar_ciclo(grafo, origen):
    cola = deque()
    padres = {}
    padres[origen] = None
    cola.append(origen)
    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in padres:
                padres[w] = v
                cola.append(w)
            elif origen == w:
                return armar_ciclo(padres, origen, v)

    return None


global orden_contador

def componentes_fuertemente_conexas(grafo, v, visitados, orden, mb, pila, apilados, todas_cfc, orden_contador):

    visitados.add(v)
    orden[v] = orden_contador
    mb[v] = orden[v]
    orden_contador += 1
    pila.apilar(v)
    apilados.add(v)


    for w in grafo.adyacentes(v):
        if w not in visitados:
            componentes_fuertemente_conexas(grafo, w, visitados, orden, mb, pila, apilados, todas_cfc, orden_contador)

        if w in apilados:
            mb[v] = min(mb[v], mb[w])

    if orden[v] == mb[v] and not pila.esta_vacia():
        nueva_cfc = []
        while True:
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break

        todas_cfc.append(nueva_cfc)

def inicializar_grafo(ruta):
    grafo = Grafo(es_dirigido=True)
    vertices_visitados = set()
    with open(f'./{ruta}') as f:
        for linea in f:
            mensaje = linea.split()
            id1, id2 = mensaje[0], mensaje[1]
            if id1 not in vertices_visitados:
                grafo.agregar_vertice(id1)
            if id2 not in vertices_visitados:
                grafo.agregar_vertice(id2)

            if not grafo.estan_unidos(id1, id2):
                peso_arista = grafo.peso_arista(id1, id2)
                grafo.agregar_arista(id1, id2, peso=1)
    return grafo

def imprimir_res(origen, destino, camino):
    actual = destino
    aux = Pila()
    aux.apilar(actual)

    while actual != origen:
        actual = camino[actual]
        aux.apilar(actual)

    while not aux.esta_vacia():
        elemento = aux.desapilar()
        if not aux.esta_vacia():
            print(f'{elemento} ->', end=" ")
        else:
            print(f'{elemento}')

CANTIDAD_DE_CAMINOS = 100 # Constante de random walk
LONGITUD = 500 # Constante de random walk

def random_walks(grafo):
    caminos = []

    for _ in range(CANTIDAD_DE_CAMINOS):
        vertice = grafo.vertice_aleatorio()
        camino = [vertice]

        for _ in range(LONGITUD):
            adyacentes = grafo.adyacentes(vertice)
            if adyacentes:
                elegido = choice(adyacentes)
                camino.append(elegido)
                vertice = elegido
        caminos.append(camino)

    return caminos

def page_rank(grafo):
    caminos = random_walks(grafo)
    vertices_centrales = []
    suma_de_caminos = []

    for camino in caminos:
        for vertice in camino:
            suma_de_caminos.append(vertice)

    mas_probables = Counter(suma_de_caminos).most_common()
    for i in range(len(mas_probables)):
        if i < len(grafo.obtener_vertices()):
            vertices_centrales.append(mas_probables[i][0])

    return vertices_centrales

def max_freq(label, adyacentes):
    labelAdyacentes = {}
    for w in adyacentes:
        if label[w] in labelAdyacentes:
            labelAdyacentes[label[w]] += 1
        else:
            labelAdyacentes[label[w]] = 1

    label_mas_repetido = None
    num_repeticiones_label = 0

    for label_w in labelAdyacentes:
        if labelAdyacentes[label_w] > num_repeticiones_label:
            label_mas_repetido = label_w
            num_repeticiones_label = labelAdyacentes[label_w]

    return label_mas_repetido

def label_propagation(grafo, n):
    labels = {}
    i = 1
    for v in grafo.obtener_vertices():
        labels[v] = i
        i += 1

    for i in range(n):
        visitados = set()
        while len(visitados) != len(grafo.obtener_vertices()):
            random_v = choice(grafo.obtener_vertices())
            if random_v not in visitados:
                visitados.add(random_v)
                labels[random_v] = max_freq(labels, grafo.adyacentes(random_v))

    comunidades = {}
    for v, comunidad in labels.items():
        if comunidad not in comunidades:
            comunidades[comunidad] = []
        comunidades[comunidad].append(v)

    return comunidades