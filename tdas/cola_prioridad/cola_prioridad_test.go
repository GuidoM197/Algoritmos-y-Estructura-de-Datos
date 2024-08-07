package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"tdas/cola_prioridad"
	"testing"
)

const (
	_TEST_VOLUMEN_ENCOLAR = 10000
)

func funcionCmpIntMax(a, b int) int {
	return a - b
}

func funcionCmpIntMin(a, b int) int {
	return b - a
}

func funcionCmpString(a, b string) int {
	return strings.Compare(a, b)
}

func TestCrearHeapVacio(t *testing.T) {
	heap := cola_prioridad.CrearHeap[int](funcionCmpIntMax)

	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")
}

func TestCrearHeapArrVacio(t *testing.T) {
	arreglo := []int{}
	heap := cola_prioridad.CrearHeapArr[int](arreglo, funcionCmpIntMax)

	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")

}

func TestCrearHeapArrMax(t *testing.T) {
	arreglo := []int{1, 5, 4, 3, 2}
	heap := cola_prioridad.CrearHeapArr[int](arreglo, funcionCmpIntMax)

	require.EqualValues(t, 5, heap.Cantidad(), "No tiene la cantidad correcta de elementos.")
	require.EqualValues(t, 5, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	for i := 5; i > 0; i-- {
		require.EqualValues(t, i, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	}
	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")
}

func TestCrearHeapArrMin(t *testing.T) {
	arreglo := []int{23, 13, 7, 2, 44, 37, 29, 9}
	heap := cola_prioridad.CrearHeapArr[int](arreglo, funcionCmpIntMin)

	require.EqualValues(t, 8, heap.Cantidad(), "No tiene la cantidad correcta de elementos.")
	require.EqualValues(t, 2, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	outPuts := []int{2, 7, 9, 13, 23, 29, 37, 44}
	for _, v := range outPuts {
		require.EqualValues(t, v, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	}
	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")
}

func TestHeapEncolar(t *testing.T) {
	heap := cola_prioridad.CrearHeap[int](funcionCmpIntMax)

	heap.Encolar(1)
	require.EqualValues(t, 1, heap.VerMax(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 1, heap.Cantidad(), "No retorna el elemento prioritario correcto.")

	heap.Encolar(5)
	require.EqualValues(t, 5, heap.VerMax(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 2, heap.Cantidad(), "No retorna el elemento prioritario correcto.")

	heap.Encolar(4)
	require.EqualValues(t, 5, heap.VerMax(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 3, heap.Cantidad(), "No retorna el elemento prioritario correcto.")

	heap.Encolar(8)
	require.EqualValues(t, 8, heap.VerMax(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 4, heap.Cantidad(), "No retorna el elemento prioritario correcto.")

	heap.Encolar(9)
	require.EqualValues(t, 9, heap.VerMax(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 5, heap.Cantidad(), "No retorna el elemento prioritario correcto.")

}

func TestHeapDesencolar(t *testing.T) {
	heap := cola_prioridad.CrearHeap[int](funcionCmpIntMax)

	heap.Encolar(1)
	heap.Encolar(5)
	heap.Encolar(4)
	heap.Encolar(8)
	heap.Encolar(9)
	require.EqualValues(t, 5, heap.Cantidad(), "No retorna el elemento prioritario correcto.")

	require.EqualValues(t, 9, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	require.EqualValues(t, 9, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 8, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	require.EqualValues(t, 8, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 5, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	require.EqualValues(t, 5, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 4, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	require.EqualValues(t, 4, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	require.EqualValues(t, 1, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	require.EqualValues(t, 1, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")
}

func TestHeapArrString(t *testing.T) {
	arreglo := []string{"a", "e", "d", "c"}
	heap := cola_prioridad.CrearHeapArr[string](arreglo, funcionCmpString)

	require.EqualValues(t, 4, heap.Cantidad(), "No tiene la cantidad correcta de elementos.")
	require.EqualValues(t, "e", heap.Desencolar(), "No retorna el elemento prioritario correcto.")

	heap.Encolar("f")
	require.EqualValues(t, "f", heap.VerMax(), "No retorna el elemento prioritario correcto.")
	heap.Encolar("e")
	require.EqualValues(t, "f", heap.VerMax(), "No retorna el elemento prioritario correcto.")
	heap.Encolar("b")

	outPuts := []string{"f", "e", "d", "c", "b", "a"} // como deveria verse lo que retorno por cada desencolar
	for _, char := range outPuts {
		require.EqualValues(t, char, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	}

	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")
}

func TestHeapArrInt(t *testing.T) {
	arreglo := []int{1, 5, 3, 2}
	heap := cola_prioridad.CrearHeapArr[int](arreglo, funcionCmpIntMin)

	require.EqualValues(t, 4, heap.Cantidad(), "No tiene la cantidad correcta de elementos.")
	require.EqualValues(t, 1, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	heap.Encolar(9)
	heap.Encolar(4)
	heap.Encolar(6)
	require.EqualValues(t, 1, heap.VerMax(), "No retorna el elemento prioritario correcto.")
	heap.Encolar(0)
	require.EqualValues(t, 0, heap.VerMax(), "No retorna el elemento prioritario correcto.")

	outPuts := []int{0, 1, 2, 3, 4, 5, 6, 9} // como deveria verse lo que retorno por cada desencolar
	for _, num := range outPuts {
		require.EqualValues(t, num, heap.Desencolar(), "No retorna el elemento prioritario correcto.")
	}

	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")
}

func TestVolumenHeap(t *testing.T) {
	heap := cola_prioridad.CrearHeap[int](funcionCmpIntMax)

	for i := range _TEST_VOLUMEN_ENCOLAR {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax(), "Al encolar no actualiza el maximo")
	}

	for i := heap.Cantidad() - 1; i > 0; i-- {
		heap.Desencolar()
		require.EqualValues(t, i-1, heap.VerMax(), "Al desencolar no actualiza el maximo")
	}
	heap.Desencolar()
	require.True(t, heap.EstaVacia(), "No arroja true cuando esta vacia.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No arroja panic al desencolar heap vacio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No arroja panic al VerMax heap vacio.")
	require.EqualValues(t, 0, heap.Cantidad(), "No tiene valor inicial 0.")

}

func TestHeapSort(t *testing.T) {
	arreglo := []int{23, 13, 7, 2, 44, 37, 29, 9}

	cola_prioridad.HeapSort[int](arreglo, funcionCmpIntMax)

	outPuts := []int{2, 7, 9, 13, 23, 29, 37, 44}
	for i, v := range outPuts {
		require.EqualValues(t, v, arreglo[i], "No retorna el elemento prioritario correcto.")
	}
}
