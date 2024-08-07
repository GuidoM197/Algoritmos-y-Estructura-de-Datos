#!/usr/bin/env python3
import biblioteca, funciones_grafo
import sys

def comandos_eleccion(parametros, grafo_delincuentes):
    comando = parametros[0]
    mas_importantes = []

    if comando == "min_seguimientos":
        if len(parametros) < 3:
            raise "Error en comandos"
        origen = parametros[1]
        destino = parametros[2]
        biblioteca.minimo_seguimientos(origen, destino, grafo_delincuentes)

    elif comando == "mas_imp":
        if len(parametros) < 2:
            raise "Error en comandos"
        cantidad = int(parametros[1])

        if not mas_importantes:
            mas_importantes = funciones_grafo.page_rank(grafo_delincuentes)
        biblioteca.delincuentres_mas_importantes(cantidad, mas_importantes)

    elif comando == "persecucion":
        if len(parametros) < 3:
            raise "Error en comandos"
        agentes_encubierto = parametros[1].split(',')
        k_mas_importantes = int(parametros[2])
        if not mas_importantes:
            mas_importantes = funciones_grafo.page_rank(grafo_delincuentes)
        camino = biblioteca.persecucion(grafo_delincuentes, agentes_encubierto, k_mas_importantes, mas_importantes)

    elif comando == "comunidades":
        if len(parametros) < 2:
            raise "Error en comandos"
        integrantes = int(parametros[1])
        biblioteca.comunidades(grafo_delincuentes, integrantes)

    elif comando == "divulgar":
        if len(parametros) < 3:
            raise "Error en comandos"
        delincuente = parametros[1]
        dist_max = int(parametros[2])
        biblioteca.divulgar(delincuente, dist_max, grafo_delincuentes)

    elif comando == "divulgar_ciclo":
        if len(parametros) < 2:
            raise "Error en comandos"
        delincuente = parametros[1]
        biblioteca.divulgar_ciclo(delincuente, grafo_delincuentes)

    elif comando == "cfc":
        if len(parametros) < 1:
            raise "Error en comandos"

        biblioteca.cfc(grafo_delincuentes)



def main(datos):
    if len(datos) != 1:
        raise ("Error en parametros de entrada")
    ruta = datos[0]
    contenido_archivo2 = sys.stdin.read().split()

    grafo_delincuentes = funciones_grafo.inicializar_grafo(ruta)
    comandos_eleccion(contenido_archivo2, grafo_delincuentes)


if __name__ == "__main__":
    main(sys.argv[1:])
