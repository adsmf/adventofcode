package murmur3

import (
	"unsafe"
)

// https://en.wikipedia.org/wiki/MurmurHash#Algorithm
//
// algorithm Murmur3_32 is
//     // Note: In this version, all arithmetic is performed with unsigned 32-bit integers.
//     //       In the case of overflow, the result is reduced modulo 232.
//     input: key, len, seed
//
//     c1 ← 0xcc9e2d51
//     c2 ← 0x1b873593
//     r1 ← 15
//     r2 ← 13
//     m ← 5
//     n ← 0xe6546b64
//
//     hash ← seed
//
//     for each fourByteChunk of key do
//         k ← fourByteChunk
//
//         k ← k × c1
//         k ← k ROL r1
//         k ← k × c2
//
//         hash ← hash XOR k
//         hash ← hash ROL r2
//         hash ← (hash × m) + n
//
//     with any remainingBytesInKey do
//         remainingBytes ← SwapToLittleEndian(remainingBytesInKey)
//         // Note: Endian swapping is only necessary on big-endian machines.
//         //       The purpose is to place the meaningful digits towards the low end of the value,
//         //       so that these digits have the greatest potential to affect the low range digits
//         //       in the subsequent multiplication.  Consider that locating the meaningful digits
//         //       in the high range would produce a greater effect upon the high digits of the
//         //       multiplication, and notably, that such high digits are likely to be discarded
//         //       by the modulo arithmetic under overflow.  We don't want that.
//
//         remainingBytes ← remainingBytes × c1
//         remainingBytes ← remainingBytes ROL r1
//         remainingBytes ← remainingBytes × c2
//
//         hash ← hash XOR remainingBytes
//
//     hash ← hash XOR len
//
//     hash ← hash XOR (hash >> 16)
//     hash ← hash × 0x85ebca6b
//     hash ← hash XOR (hash >> 13)
//     hash ← hash × 0xc2b2ae35
//     hash ← hash XOR (hash >> 16)

type Murmur3_32 struct {
	seed uint32
}

func NewMurmer3_32(seed uint32) Murmur3_32 {
	return Murmur3_32{
		seed: seed,
	}
}

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
	r1        = 15
	r2        = 13
	m  uint32 = 5
	n  uint32 = 0xe6546b64
)

func (m32 Murmur3_32) HashBytes(key []byte) uint32 {
	hash := m32.seed

	for i := 0; i <= len(key)-4; i += 4 {
		k := *(*uint32)(unsafe.Pointer(&key[i]))
		k *= c1
		k = k<<r1 | k>>(32-r1)
		k *= c2
		hash ^= k
		hash = hash<<r2 | hash>>(32-r2)
		hash = hash*m + n
	}

	var k uint32
	rem := key[(len(key)>>2)<<2:]
	switch len(rem) {
	case 3:
		k ^= uint32(rem[2]) << 16
		fallthrough
	case 2:
		k ^= uint32(rem[1]) << 8
		fallthrough
	case 1:
		k ^= uint32(rem[0])
		k *= c1
		k = k<<r1 | k>>(32-r1)
		k *= c2
		hash ^= k
	}

	hash ^= uint32(len(key))

	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16

	return hash
}
