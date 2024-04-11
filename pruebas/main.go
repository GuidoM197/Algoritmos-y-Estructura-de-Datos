package main

import TDAPila "tdas/pila"

func main() {
	p := TDAPila.CrearPilaDinamica[int]()
	p.Apilar(1)
	p.Apilar(2)
	p.Apilar(3)
	p.Apilar(4)
	p.Apilar(5)
	p.VerPila()
}
