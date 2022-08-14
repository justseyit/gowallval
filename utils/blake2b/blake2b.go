package blake2b

import "encoding/binary"

const (
	// The size of a BLAKE2b hash in bytes.
	BLAKE2B_BLOCKBYTES  = 64
	BLAKE2B_DIGESTBYTES = 32
)

type Blake2b struct {
	h [8]uint64
	s [4]uint64
	t [2]uint64
	f [2]uint64
	g [8]uint64
	x [16]uint64
	b [BLAKE2B_BLOCKBYTES]byte
	n uint64
}

func New() *Blake2b {
	b := new(Blake2b)
	b.Reset()
	return b
}

func (b *Blake2b) Reset() {
	b.h = [8]uint64{
		0x6a09e667f3bcc908,
		0xbb67ae8584caa73b,
		0x3c6ef372fe94f82b,
		0xa54ff53a5f1d36f1,
		0x510e527fade682d1,
		0x9b05688c2b3e6c1f,
		0x1f83d9abfb41bd6b,
		0x5be0cd19137e2179,
	}
	b.s = [4]uint64{
		0x6a09e667f3bcc908,
		0xbb67ae8584caa73b,
		0x3c6ef372fe94f82b,
		0xa54ff53a5f1d36f1,
	}
	b.t = [2]uint64{
		0x510e527fade682d1,
		0x9b05688c2b3e6c1f,
	}
	b.f = [2]uint64{
		0x1f83d9abfb41bd6b,
		0x5be0cd19137e2179,
	}
	b.g = [8]uint64{
		0x6a09e667f3bcc908,
		0xbb67ae8584caa73b,
		0x3c6ef372fe94f82b,
		0xa54ff53a5f1d36f1,
		0x510e527fade682d1,
		0x9b05688c2b3e6c1f,
		0x1f83d9abfb41bd6b,
		0x5be0cd19137e2179,
	}
	b.x = [16]uint64{
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
	b.n = 0
}

func (b *Blake2b) Update(data []byte) {
	for i := 0; i < len(data); i++ {
		b.b[b.n] = data[i]
		b.n++
		if b.n == BLAKE2B_BLOCKBYTES {
			b.compress()
			b.n = 0
		}
	}
}

func (b *Blake2b) Final() []byte {
	b.Update([]byte{0x81})
	for i := b.n; i < BLAKE2B_BLOCKBYTES; i++ {
		b.b[i] = 0
	}
	b.compress()
	var out [32]byte
	for i := 0; i < 8; i++ {
		binary.LittleEndian.PutUint64(out[i*8:], b.h[i])
	}
	return out[:]
}

func (b *Blake2b) compress() {
	var v [16]uint64
	for i := 0; i < 16; i++ {
		v[i] = b.x[i]
	}
	for i := 0; i < 8; i++ {
		v[i] = v[i] ^ b.f[i%2]
	}
	for i := 0; i < 8; i++ {
		v[i+8] = v[i+8] ^ b.g[i]
	}
	for i := 0; i < 12; i++ {
		v[i] = v[i] ^ v[i+4]
	}
	for i := 0; i < 16; i++ {
		v[i] = b.rot64(v[i], uint64(i%8*8))
	}
	for i := 0; i < 16; i++ {
		v[i] = v[i] ^ v[i+8]
	}
	for i := 0; i < 16; i++ {
		v[i] = v[i] ^ b.h[i%8]
	}
	for i := 0; i < 8; i++ {
		b.h[i] = b.h[i] ^ v[i] ^ v[i+8]
	}
}

func (b *Blake2b) rot64(x uint64, n uint64) uint64 {
	return (x << n) | (x >> (64 - n))
}
