import funciones_grafo
from pila import Pila

def minimo_seguimientos(origen, destino, grafo_delincuentes):
    padres, _ = funciones_grafo.bfs(grafo_delincuentes, origen, destino)
    if not destino in padres:
        print("Seguimiento imposible")
    else:
        camino = armar_camino(origen, destino, padres)
        imprimir_res(camino, " ->")

def delincuentres_mas_importantes(cantidad, mas_imp):
    contador = 1

    for id in mas_imp:
        if contador < cantidad:
            print(id, end=", ")
        else:
            print(id)
            break
        contador += 1

def persecucion(grafo_delincuentes, agentes_encubierto, k, mas_importantes):
    camino_res = {}
    agente_importante = None
    delincuentre_importante = None
    orden_mas_importante = 0
    k_mas_importantes = []

    for i in range(k):
        k_mas_importantes.append(mas_importantes[i])

    for agente in agentes_encubierto:
        for importante in k_mas_importantes:
            camino_minimo, orden = funciones_grafo.bfs(grafo_delincuentes, agente, importante)

            if len(camino_res) == 0: # Unicamente en la primera iteración
                agente_importante = agente
                delincuentre_importante = importante
                orden_mas_importante = orden[importante]
                camino_res = camino_minimo

            elif orden_mas_importante > orden[importante]:
                agente_importante = agente
                delincuentre_importante = importante
                orden_mas_importante = orden[importante]
                camino_res = camino_minimo

    camino = armar_camino(agente_importante, delincuentre_importante, camino_res)
    imprimir_res(camino, " ->")

def comunidades(grafo_delincuentes, n_integrantes):
    comunidades = funciones_grafo.label_propagation(grafo_delincuentes, n_integrantes)
    num_comunidad = 1

    for comunidades, integrantes in comunidades.items():
        if len(integrantes) >= n_integrantes:
            print("Comunidad", str(num_comunidad), end=": ")
            imprimir_res(integrantes, ',')
            num_comunidad += 1

def divulgar(delincuente, dist_max, grafo_delincuentes):
    res = []
    _, orden = funciones_grafo.bfs(grafo_delincuentes, delincuente)

    for key, values in orden.items():
        if int(values) <= dist_max and key != delincuente:
            res.append(key)

    imprimir_res(res, ',')

def divulgar_ciclo(delincuente, grafo_delincuentes):
    ciclo = funciones_grafo.bfs_encontrar_ciclo(grafo_delincuentes, delincuente)
    if ciclo is None:
        print("No se encontro recorrido")
    else:
        imprimir_res(ciclo, " ->")

def cfc(grafo_delincuentes):
    cfcs = []
    funciones_grafo.componentes_fuertemente_conexas(grafo_delincuentes, grafo_delincuentes.vertice_aleatorio(), set(), {}, {}, Pila(), set(), cfcs, 0)
    for i, cfc in enumerate(cfcs):
        print(f'CFC {i+1}: {cfc}')

# ------------------------- Aux ------------------------- #

def armar_camino(origen, destino, camino):
    actual = destino
    res = [actual]

    while actual != origen:
        actual = camino[actual]
        res.append(actual)

    return res[::-1]

def imprimir_res(camino, algo):
    for i in range(len(camino)):
        if i != len(camino) - 1:
            print(f'{camino[i]}{algo}', end=" ")
        else:
            print(f'{camino[i]}')