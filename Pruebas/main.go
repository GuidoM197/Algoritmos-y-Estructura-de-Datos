package main

import (
  "fmt"
)

func Swap(x *int, y *int) {
  *x, *y = *y, *x
}

func Maximo(vector []int) int {
  if len(vector) == 0 {
    return -1
  }

  maximo, posicionMax := vector[0], 0

  for i := 0; i < len(vector); i++ {
    
    if vector[i] > maximo {
      posicionMax = i
      maximo = vector[i]
    }
  }

  return posicionMax

}

func Seleccion(vector []int) {

  for i := (len(vector) - 1); i > 0; i-- {
    posicionMax := Maximo(vector[:i+1])
    Swap(&vector[i], &vector[posicionMax])

  }

}

func main() {

  arr := []int{1,2,3,4,5}
  Seleccion(arr)
  fmt.Println(arr)


}
