package main

import "fmt"

func Comparar(vector1 []int, vector2 []int) int {
  // Verifico que ninguno sea un areglo vacio para meter un early return en caso de ser necesario para no ejecutar codigo de mas.
  if len(vector1) == 0 && len(vector2) != 0 {
    return -1
  } else if len(vector1) != 0 && len(vector2) == 0 {
    return 1
  } else if len(vector1) == 0 && len(vector2) == 0 {
    return 0
  }

  // Determino si hay algun arreglo con menos longitud de antemano y lo utilizo para comparar los elementos en ese rango, tambien me guardo el numero del mas largo.
  longitud_max := 0
  vector_mas_largo := 0

  if len(vector1) > len(vector2) {
    longitud_max = len(vector2)
    vector_mas_largo = 1
  } else if len(vector1) < len(vector2) {
    longitud_max = len(vector1)
    vector_mas_largo = 2
  } else {
    // Ya que ambos vectores tendrian el mismo tamaño elijo cualquiera de los dos.
    longitud_max = len(vector1) 
  }

  // En este punto es indiferente que arreglo usar como referencia de len ya que son del mismo tamaño.
  for i := 0; i < longitud_max; i++ { 

    elemento_v1, elemento_v2 := vector1[i], vector2[i]

    if elemento_v1 > elemento_v2 {
        return 1
    } else if elemento_v1 < elemento_v2 {
        return -1
    }
	
  }

  // Si no tienen elementos mas grandes determino el mayor arreglo segun su tamaño.
  if vector_mas_largo == 1 {
    return 1
  } else if vector_mas_largo == 2 {
    return -1
  }

  // Serian iguales. 
  return 0
  
}

func main() {

  a := []int{10}
  b := []int{10}
  // Comparar(a, b)
  fmt.Println(Comparar(a, b))
  fmt.Println(Comparar(b, a))
}
