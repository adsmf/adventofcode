package fnv

const (
	fnvp32 uint32 = 0x01000193
	fnvo32 uint32 = 0x811c9dc5
)

// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1a_hash

// algorithm fnv-1a is
//     hash := FNV_offset_basis

//     for each byte_of_data to be hashed do
//         hash := hash XOR byte_of_data
//         hash := hash × FNV_prime

//     return hash

func FNV1a_HashBytes32(input []byte) uint32 {
	hash := fnvo32

	// for each byte_of_data to be hashed do
	//   hash := hash XOR byte_of_data
	//   hash := hash × FNV_prime

	for _, b := range input {
		hash ^= uint32(b)
		hash *= fnvp32
	}

	return hash
}

type Hasher struct {
	value uint32
}

func (h *Hasher) Reset() {
	h.value = fnvo32
}
func (h *Hasher) AddByte(b byte) {
	h.value ^= uint32(b)
	h.value *= fnvp32
}
func (h Hasher) Sum() uint32 {
	return h.value
}
