package blake256

import (
	"encoding/binary"
)

const (
	BLAKE256_BLOCKBYTES  = 64
	BLAKE256_DIGESTBYTES = 32
)

//Define a Blake256 struct

type Blake256 struct {
	h [8]uint32
	s [4]uint32
	t [2]uint32
	f [2]uint32
	g [8]uint32
	b [BLAKE256_BLOCKBYTES]byte
	c [2]uint32
	n uint32
	x [BLAKE256_BLOCKBYTES]byte
}

// Define a New() function that returns a pointer to a Blake256 struct.
func New() *Blake256 {
	h := [8]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}
	s := [4]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
	}
	t := [2]uint32{
		0, 0,
	}
	f := [2]uint32{
		0, 0,
	}
	b := [BLAKE256_BLOCKBYTES]byte{}
	c := [2]uint32{}
	n := uint32(0)
	x := [BLAKE256_BLOCKBYTES]byte{}
	g := [8]uint32{}
	return &Blake256{h, s, t, f, g, b, c, n, x}
}

func (b *Blake256) Update(data []byte) {
	b.n += uint32(len(data))
	if b.n > BLAKE256_BLOCKBYTES {
		panic("blake256: message too long")
	}
	copy(b.b[b.n:], data)
	if b.n == BLAKE256_BLOCKBYTES {
		b.compress()
	}
}

func (b *Blake256) Final() []byte {
	b.Update([]byte{0x81})
	for i := b.n; i < BLAKE256_BLOCKBYTES; i++ {
		b.b[i] = 0
	}
	binary.LittleEndian.PutUint64(b.b[BLAKE256_BLOCKBYTES-8:], uint64(b.n*8))
	binary.LittleEndian.PutUint64(b.b[BLAKE256_BLOCKBYTES-16:], uint64(b.n*8))
	b.compress()
	var digest [BLAKE256_DIGESTBYTES]byte
	for i := 0; i < 8; i++ {
		binary.LittleEndian.PutUint32(digest[i*4:], b.h[i])
	}
	return digest[:]
}

func (b *Blake256) compress() {
	for i := 0; i < 16; i++ {
		b.x[i] = b.b[i]
	}
	for i := 16; i < BLAKE256_BLOCKBYTES; i++ {
		b.x[i] = b.b[i] ^ b.x[i-16]
	}
	for i := 0; i < 8; i++ {
		b.h[i] = b.s[i]
	}
	for i := 0; i < 8; i++ {
		b.f[0] = b.h[0] ^ b.h[1] ^ b.h[2] ^ b.h[3] ^ b.h[4] ^ b.h[5] ^ b.h[6] ^ b.h[7] ^ uint32(i)
		b.f[1] = b.f[0] ^ b.h[0]
		b.g[0] = (b.f[0] + b.h[0]) ^ b.h[1]
		b.g[1] = (b.f[1] + b.h[1]) ^ b.h[2]
		b.g[2] = (b.g[0] + b.h[2]) ^ b.h[3]
		b.g[3] = (b.g[1] + b.h[3]) ^ b.h[4]
		b.g[4] = (b.g[2] + b.h[4]) ^ b.h[5]
		b.g[5] = (b.g[3] + b.h[5]) ^ b.h[6]
		b.g[6] = (b.
			g[4] + b.h[6]) ^ b.h[7]
		b.g[7] = (b.g[5] + b.h[7]) ^ b.h[0]
		b.h[0] = b.g[0]
		b.h[1] = b.g[1]
		b.h[2] = b.g[2]
		b.h[3] = b.g[3]
		b.h[4] = b.g[4]
		b.h[5] = b.g[5]
		b.h[6] = b.g[6]
		b.h[7] = b.g[7]
	}
	for i := 0; i < 8; i++ {
		//convert b.h[i] to bytes
		binary.LittleEndian.PutUint32(b.b[i*4:], b.h[i])
		//b.h[i] ^= b.x[i]
	}
}

func (b *Blake256) Reset() {
	b.h = [8]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}
	b.s = [4]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
	}
	b.t = [2]uint32{
		0, 0,
	}
	b.f = [2]uint32{
		0, 0,
	}
	b.b = [BLAKE256_BLOCKBYTES]byte{}
	b.c = [2]uint32{}
	b.n = uint32(0)
	b.x = [BLAKE256_BLOCKBYTES]byte{}
}

func (b *Blake256) Size() int {
	return BLAKE256_DIGESTBYTES
}

func (b *Blake256) BlockSize() int {
	return BLAKE256_BLOCKBYTES
}

func (b *Blake256) Sum(data []byte) []byte {
	t := b.Final()
	return append(data, t...)
}

func (b *Blake256) ResetWithIV(iv []byte) {
	b.Reset()
	copy(b.b[:], iv)
}

func (b *Blake256) SumWithIV(data, iv []byte) []byte {
	b.ResetWithIV(iv)
	b.Update(data)

	return b.Final()
}

func (b *Blake256) Sum256(data []byte) []byte {
	return b.Sum(data)
}

func (b *Blake256) Sum256WithIV(data, iv []byte) []byte {
	return b.SumWithIV(data, iv)
}
