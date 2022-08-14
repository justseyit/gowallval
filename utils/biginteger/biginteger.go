package biginteger

/*import "strconv"

//Write a big integer library in Go.
//
//The library should support the following operations:
//
//Addition
//Subtraction
//Multiplication
//Division
//Modulo
//Exponentiation
//
//The library should support the following types:
//
//uint8, uint16, uint32, uint64
//int8, int16, int32, int64
//uint, int
//


type BigInteger struct {
	value []uint64
	sign  int
}

func NewBigInteger(value []uint64, sign int) *BigInteger {
	return &BigInteger{value, sign}
}

func (bi *BigInteger) GetSign() int {
	return bi.sign
}

func (bi *BigInteger) GetValue() []uint64 {
	return bi.value
}

func (bi *BigInteger) GetBit(index int) uint64 {
	return bi.value[index/64] >> (uint64(index%64)) & 1
}

func (bi *BigInteger) GetBitLength() int {
	return len(bi.value) * 64
}

func (bi *BigInteger) Add(bigint BigInteger) int {
	if bi.sign == bigint.sign {
		bi.value = add(bi.value, bigint.value)
		bi.sign = 1
	} else {
		bi.value = subtract(bi.value, bigint.value)
		bi.sign = -1
	}
	return bi.sign
}

func (bi *BigInteger) Subtract(bigint BigInteger) int {
	if bi.sign == bigint.sign {
		bi.value = subtract(bi.value, bigint.value)
		bi.sign = 1
	} else {
		bi.value = add(bi.value, bigint.value)
		bi.sign = -1
	}
	return bi.sign
}

func (bi *BigInteger) Multiply(bigint BigInteger) int {
	bi.value = multiply(bi.value, bigint.value)
	bi.sign = bi.sign * bigint.sign
	return bi.sign
}

func (bi *BigInteger) Divide(bigint BigInteger) int {
	bi.value = divide(bi.value, bigint.value)
	bi.sign = bi.sign * bigint.sign
	return bi.sign
}

func (bi *BigInteger) Modulo(bigint BigInteger) int {
	bi.value = modulo(bi.value, bigint.value)
	bi.sign = bi.sign * bigint.sign
	return bi.sign
}

func (bi *BigInteger) Exponentiation(bigint BigInteger) int {
	bi.value = exponentiation(bi.value, bigint.value)
	bi.sign = bi.sign * bigint.sign
	return bi.sign
}

func (bi *BigInteger) ToString() string {
	return toString(bi.value)
}

func add(a []uint64, b []uint64) []uint64 {
	if len(a) > len(b) {
		return add(b, a)
	}
	c := make([]uint64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	for i := len(a); i < len(b); i++ {
		c[i] = b[i]
	}
	for i := 0; i < len(c)-1; i++ {
		if c[i] >= 1<<64 {
			c[i+1]++
			c[i] -= 1<<64
		}
	}
	if c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

func subtract(a []uint64, b []uint64) []uint64 {
	if len(a) > len(b) {
		return subtract(b, a)
	}
	c := make([]uint64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] - b[i]
	}
	for i := len(a); i < len(b); i++ {
		c[i] = 0 - b[i]
	}
	for i := 0; i < len(c)-1; i++ {
		if c[i] < 0 {
			c[i+1]--
			c[i] += 1<<64
		}
	}
	if c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

func multiply(a []uint64, b []uint64) []uint64 {
	c := make([]uint64, len(a)+len(b))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i+j] += a[i] * b[j]
		}
	}
	for i := 0; i < len(c)-1; i++ {
		if c[i] >= 1<<64 {
			c[i+1]++
			c[i] -= 1<<64
		}
	}
	if c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

func divide(a []uint64, b []uint64) []uint64 {
	c := make([]uint64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	for i := 0; i < len(c); i++ {
		c[i] = c[i] / b[0]
	}
	return c
}

func modulo(a []uint64, b []uint64) []uint64 {
	c := make([]uint64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	for i := 0; i < len(c); i++ {
		c[i] = c[i] % b[0]
	}
	return c
}

func exponentiation(a []uint64, b []uint64) []uint64 {
	c := make([]uint64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		c = multiply(c, c)
	}
	return c
}


func toString(a []uint64) string {
	s := ""
	if a[0] == 0 {
		s = "0"
	} else {
		for i := len(a) - 1; i >= 0; i-- {
			s += strconv.FormatUint(a[i], 10)
		}
	}
	return s
}
*/
