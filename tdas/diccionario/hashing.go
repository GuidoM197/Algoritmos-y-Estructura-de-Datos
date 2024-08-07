package diccionario

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

// xxhash64 calculates the 64-bit hash of the input data using XXHash algorithm
func xxhash64(data []byte) uint64 {
	length := len(data)
	var (
		i      int
		h64    uint64
		prime1 uint64
		prime2 uint64
		prime3 uint64
		prime4 uint64
		prime5 uint64
	)
	prime1 = 11400714785074694791
	prime2 = 14029467366897019727
	prime3 = 1609587929392839161
	prime4 = 9650029242287828579
	prime5 = 2870177450012600261

	if length >= 32 {
		end := length - 32
		for i = 0; i <= end; i += 32 {
			v1 := binary.LittleEndian.Uint64(data[i:])
			v2 := binary.LittleEndian.Uint64(data[i+8:])
			v3 := binary.LittleEndian.Uint64(data[i+16:])
			v4 := binary.LittleEndian.Uint64(data[i+24:])
			h64 += v1 * prime2
			h64 = (bits.RotateLeft64(h64, 31)*prime1 + h64*prime4) + v2*prime2
			h64 = (bits.RotateLeft64(h64, 27)*prime1 + h64*prime4) + v3*prime2
			h64 = (bits.RotateLeft64(h64, 33)*prime1 + h64*prime4) + v4*prime2
		}
	}

	if i <= length-16 {
		v1 := binary.LittleEndian.Uint64(data[i:])
		v2 := binary.LittleEndian.Uint64(data[i+8:])
		h64 += v1 * prime2
		h64 = (bits.RotateLeft64(h64, 31) * prime1) + v2*prime2
		i += 16
	}

	if i <= length-8 {
		v := binary.LittleEndian.Uint64(data[i:])
		h64 ^= (bits.RotateLeft64(v*prime2, 37) * prime1)
		h64 = (h64*prime1 + prime4) * prime2
		i += 8
	}

	if i <= length-4 {
		v := binary.LittleEndian.Uint32(data[i:])
		h64 ^= uint64(v) * prime1
		h64 = (h64*prime2 + prime3)
		i += 4
	}

	for i < length {
		h64 ^= uint64(data[i]) * prime5
		h64 = (h64*prime2 + prime1)
		i++
	}

	h64 ^= uint64(length)

	h64 ^= h64 >> 33
	h64 *= prime2
	h64 ^= h64 >> 29
	h64 *= prime3
	h64 ^= h64 >> 32

	// Asegurarse de que el resultado sea no negativo
	return h64 & ((1 << 63) - 1)
}

// convertirABytes convierte una clave en un slice de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
