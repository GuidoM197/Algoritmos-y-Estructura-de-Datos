"Esta la tengo de Algo1 asi que me la traje."

class Pila:
    def __init__(self):
        self.tope = None

    def apilar(self, dato):
        nodo = _Nodo(dato, self.tope)
        self.tope = nodo

    def desapilar(self):
        if self.esta_vacia():
            raise ValueError("pila vacía")
        dato = self.tope.dato
        self.tope = self.tope.prox
        return dato

    def ver_tope(self):
        if self.esta_vacia():
            raise ValueError("pila vacía")
        return self.tope.dato

    def esta_vacia(self):
        return self.tope is None

    def __str__(self):
        res = "tope <| "
        act = self.tope
        while act:
            res += str(act.dato)
            if act.prox:
                res += " <- "
            act = act.prox
        return res + " <| fondo"

class _Nodo:
    def __init__(self, dato, prox=None):
        self.dato = dato
        self.prox = prox