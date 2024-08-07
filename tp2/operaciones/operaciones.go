package operaciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"time"
)

const (
	_INICIALIZAR_CANTIDAD     = 1
	_INICIALIZAR_REPETICIONES = 1
	_POSICION_DE_BYTE         = 4
	_FACTOR_ES_DOS            = 5
	_LAYOUT                   = "2006-01-02T15:04:05-07:00"
)

type DiccOrdenado diccionario.DiccionarioOrdenado[string, Datos]
type DiccDicc diccionario.Diccionario[string, DiccOrdenado]

func AgregarArchivo(nombreArchivo string, visitantes DiccDicc, urlsMasVisitadas *[]Par) error {
	logs, err := procesadorDeArchivos(nombreArchivo)
	if err != nil {
		return err
	}

	visitados := diccionario.CrearHash[string, string]()
	DoSs := diccionario.CrearABB[string, bool](funcComparacionIp)
	for _, peticion := range logs {

		if !visitantes.Pertenece(peticion.url) {
			visitantes.Guardar(peticion.url, diccionario.CrearABB[string, Datos](funcComparacionIp))
		}

		url := visitantes.Obtener(peticion.url)

		if !url.Pertenece(peticion.ip) {
			encontroIp := false
			if visitados.Pertenece(peticion.ip) {
				urlExistente := visitados.Obtener(peticion.ip)
				dicIp := visitantes.Obtener(urlExistente)
				encontroIp = true

				dato := dicIp.Obtener(peticion.ip)
				url.Guardar(peticion.ip, Datos{ultimaConexion: dato.ultimaConexion, cantidadEn5seg: dato.cantidadEn5seg, cantidadTotal: _INICIALIZAR_CANTIDAD, DoS: dato.DoS})
				verificarConexiones(peticion, url, DoSs)

			}
			if !encontroIp {
				inicializarRepe := _INICIALIZAR_REPETICIONES
				dos := false
				url.Guardar(peticion.ip, Datos{ultimaConexion: &peticion.fecha, cantidadEn5seg: &inicializarRepe, cantidadTotal: _INICIALIZAR_CANTIDAD, DoS: &dos})
				visitados.Guardar(peticion.ip, peticion.url)
			}

		} else {
			verificarConexiones(peticion, url, DoSs)

		}

	}
	for iter := DoSs.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, _ := iter.VerActual()
		fmt.Printf("DoS: %s\n", ip)
	}
	*urlsMasVisitadas = crearArregloDePares(visitantes)
	restablecerDos(visitantes)
	fmt.Println("OK")

	return nil
}

func VerMasVisitados(cantidadUrlPedidas string, urlsMasVisitadas []Par) {
	k, _ := strconv.Atoi(cantidadUrlPedidas)
	if k > len(urlsMasVisitadas) {
		k = len(urlsMasVisitadas)
	}
	heap := cola_prioridad.CrearHeapArr(urlsMasVisitadas[:k], cmpMaxPar)

	for _, pares := range urlsMasVisitadas[k:] {
		if pares.cantTolal < heap.VerMax().cantTolal {
			continue
		}
		heap.Encolar(pares)
		heap.Desencolar()
	}

	urlsPedidas := make([]Par, k)
	for i := range k {
		urlsPedidas[i] = heap.Desencolar()
	}
	fmt.Println("Sitios más visitados:")
	for _, pares := range urlsPedidas {
		fmt.Printf("\t%s - %d\n", pares.url, pares.cantTolal)
	}
	fmt.Println("OK")

}

func VerVisitantes(urls DiccDicc, desde string, hasta string) {
	ABBIps := diccionario.CrearABB[string, int](funcComparacionIp)

	for iter := urls.Iterador(); iter.HaySiguiente(); iter.Siguiente() { //	O(u)
		_, ips := iter.VerActual()
		for iterIps := ips.IteradorRango(&desde, &hasta); iterIps.HaySiguiente(); iterIps.Siguiente() { //O(log v) <<< O(v) y en peor de los caso O(v)
			ip, _ := iterIps.VerActual()
			ABBIps.Guardar(ip, 0)
		}
	}
	fmt.Println("Visitantes:")
	for iterAbb := ABBIps.Iterador(); iterAbb.HaySiguiente(); iterAbb.Siguiente() { // O(w) siendo w la canidad guardada de ips
		ipsValidas, _ := iterAbb.VerActual() // en el mejor caso w <<< v y en el peor w == v
		fmt.Printf("\t%s\n", ipsValidas)
	}
	fmt.Println("OK")

}

// ========================= Aux ========================= //

type Datos struct {
	ultimaConexion *string
	cantidadEn5seg *int
	DoS            *bool
	cantidadTotal  int
}

type log struct {
	ip    string
	fecha string
	url   string
}

type Par struct {
	url       string
	cantTolal int
}

func procesadorDeArchivos(nombreArchivo string) ([]log, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("Error en comando agregar_archivo")
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	var logs []log

	for s.Scan() {
		line := s.Text()
		logActual := strings.Fields(line)
		logs = append(logs, log{ip: logActual[0], fecha: logActual[1], url: logActual[3]})
	}
	err = s.Err()
	if err != nil {
		return nil, fmt.Errorf("Error en comando agregar_archivo")
	}

	return logs, nil
}

func verificarConexiones(peticion log, url DiccOrdenado, DoSs diccionario.DiccionarioOrdenado[string, bool]) {
	ip := url.Obtener(peticion.ip)
	ultimaConexion, nuevaConexion := ip.ultimaConexion, peticion.fecha

	if CompararTiempos(*ultimaConexion, nuevaConexion) {
		*ip.cantidadEn5seg += 1
		ipActualizada := url.Obtener(peticion.ip)

		if esDoS(ipActualizada) && !*ip.DoS {
			*ip.DoS = true
			DoSs.Guardar(peticion.ip, true)
		}

	} else {
		*ip.cantidadEn5seg = 1
		*ultimaConexion = nuevaConexion
	}
	url.Guardar(peticion.ip, Datos{ultimaConexion: ultimaConexion, cantidadEn5seg: ip.cantidadEn5seg, cantidadTotal: ip.cantidadTotal + 1, DoS: ip.DoS})

}

func CompararTiempos(ultimaConexion string, nuevaConexion string) bool {
	start, _ := time.Parse(_LAYOUT, ultimaConexion)
	actual, _ := time.Parse(_LAYOUT, nuevaConexion)
	elapsed := actual.Sub(start)
	return elapsed < 2*time.Second
}

func esDoS(d Datos) bool {
	return *d.cantidadEn5seg == _FACTOR_ES_DOS
}

func funcComparacionIp(ip1 string, ip2 string) int {
	vetorIp1 := strings.Split(ip1, ".")
	vetorIp2 := strings.Split(ip2, ".")

	for i := 0; i < _POSICION_DE_BYTE; i++ {
		num1, _ := strconv.Atoi(vetorIp1[i])
		num2, _ := strconv.Atoi(vetorIp2[i])

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	return 0
}

func cmpMaxPar(url1, url2 Par) int {
	return url1.cantTolal - url2.cantTolal
}

func contarAccesosTotales(dict DiccOrdenado) int {
	cantidadTotal := 0
	for iter := dict.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, datos := iter.VerActual()
		cantidadTotal += datos.cantidadTotal
	}
	return cantidadTotal
}

func crearArregloDePares(visitantes DiccDicc) []Par {
	arregloPares := []Par{}
	for iter := visitantes.Iterador(); iter.HaySiguiente(); iter.Siguiente() { //O(u) siendo u las urls
		urlActual, dictIps := iter.VerActual()
		cantidadTotalIPs := contarAccesosTotales(dictIps) // O(i) siendo i las ips que son
		nuevoPar := Par{url: urlActual, cantTolal: cantidadTotalIPs}
		arregloPares = append(arregloPares, nuevoPar)
	}
	// O(u+i) === O(n)
	return arregloPares
}

func restablecerDos(visitantes DiccDicc) {

	for iterVisitantes := visitantes.Iterador(); iterVisitantes.HaySiguiente(); iterVisitantes.Siguiente() {
		_, ipAcutal := iterVisitantes.VerActual()
		for iterIp := ipAcutal.Iterador(); iterIp.HaySiguiente(); iterIp.Siguiente() {
			nombreIp, dato := iterIp.VerActual()
			dos := false
			*dato.cantidadEn5seg = _INICIALIZAR_CANTIDAD
			*dato.ultimaConexion = ""
			ipAcutal.Guardar(nombreIp, Datos{ultimaConexion: dato.ultimaConexion, cantidadEn5seg: dato.cantidadEn5seg, cantidadTotal: dato.cantidadTotal, DoS: &dos})
		}
	}
}
