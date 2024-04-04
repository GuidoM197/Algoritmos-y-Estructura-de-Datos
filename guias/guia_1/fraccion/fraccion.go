package fraccion

import "fmt"

type Fraccion struct {
	numerador   int
	denominador int
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func simplificar(num *int, den *int) {
	if *num == 0 {
		return
	}
	// Si el denominador es negativo, entonces invierto.
	// Entonces si el numerador era negativo, ambos quedan positivos,
	// y si no me queda el negativo en el numerador.
	if *den < 0 {
		*num *= -1
		*den *= -1
	}

	var maxPosibleDivisor int
	if abs(*num) < abs(*den) {
		maxPosibleDivisor = abs(*num)
	} else {
		maxPosibleDivisor = abs(*den)
	}
	for i := 2; i <= maxPosibleDivisor; i++ {
		for *num%i == 0 && *den%i == 0 {
			*num /= i
			*den /= i
		}
	}
}

// CrearFraccion crea una fraccion con el numerador y denominador indicados.
// Si el denominador es 0, entra en panico.
func CrearFraccion(numerador, denominador int) Fraccion {
	if denominador == 0 {
		panic("El denomindor es 0")
	}
	simplificar(&numerador, &denominador)
	return Fraccion{numerador: numerador, denominador: denominador}
}

// Sumar crea una nueva fraccion, con el resultante de hacer la suma de las fracciones originales
func (f Fraccion) Sumar(otra Fraccion) Fraccion {
	denominador := f.denominador * otra.denominador
	numerador := (f.numerador * otra.denominador) + (f.denominador * otra.denominador)
	return CrearFraccion(numerador, denominador)
}

// Multiplicar crea una nueva fraccion con el resultante de multiplicar ambas fracciones originales
func (f Fraccion) Multiplicar(otra Fraccion) Fraccion {
	numerador := f.numerador * otra.numerador
	denominador := f.denominador * otra.denominador
	return CrearFraccion(numerador, denominador)
}

// ParteEntera devuelve la parte entera del numero representado por la fracción.
// Por ejemplo, para "7/2" = 3.5 debe devolver 3.
func (f Fraccion) ParteEntera() int {
	return (f.numerador / f.denominador)
}

// Representacion devuelve una representación en cadena de la fraccion simplificada (por ejemplo, no puede devolverse
// "10/8" sino que debe ser "5/4"). Considerar que si se trata de un número entero, debe mostrarse como tal.
// Considerar tambien el caso que se trate de un número negativo.
func (f Fraccion) Representacion() string {
	if f.denominador == 1 {
		return fmt.Sprintf("%d", f.numerador)
	} else {
		return fmt.Sprintf("d%/%d", f.numerador, f.denominador)
	}
}
