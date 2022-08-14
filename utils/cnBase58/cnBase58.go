package cnbase58

import (
	"fmt"
	"math"
	"math/big"
)

const ALPHABET_STR string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var ALPHABET = make([]uint8, 0)

func init() {
	for _, c := range ALPHABET_STR {
		ALPHABET = append(ALPHABET, uint8(c))
	}
}

var encoded_block_sizes = []int{0, 2, 3, 5, 6, 7, 9, 10, 11}
var alphabet_size = len(ALPHABET)
var full_block_size = 8
var full_encoded_block_size = 11

var UINT64_MAX = big.NewInt(2).Exp(big.NewInt(2), big.NewInt(64), nil)

func hextobin(hex string) ([]uint8, error) {
	if len(hex)%2 != 0 {
		return nil, fmt.Errorf("invalid hex string (length must be even)")
	}
	res := make([]uint8, len(hex)/2)
	for i := 0; i < len(hex); i += 2 {
		a, err := hexToByte(hex[i : i+2])
		if err != nil {
			return nil, err
		}
		res[i/2] = a
	}
	return res, nil
}

func hexToByte(hex string) (uint8, error) {
	if len(hex) != 2 {
		return 0, fmt.Errorf("invalid hex string (must be 2 characters)")
	}
	a, err := hexToInt(rune(hex[0]))
	if err != nil {
		return 0, err
	}
	b, err := hexToInt(rune(hex[1]))
	if err != nil {
		return 0, err
	}
	return uint8(a*16 + b), nil
}

func hexToInt(c rune) (int, error) {
	if c >= '0' && c <= '9' {
		return int(c - '0'), nil
	} else if c >= 'a' && c <= 'f' {
		return int(c - 'a' + 10), nil
	} else if c >= 'A' && c <= 'F' {
		return int(c - 'A' + 10), nil
	} else {
		return 0, fmt.Errorf("invalid hex character: %v", c)
	}
}

func bintohex(bin []uint8) string {
	res := make([]rune, 0)
	for _, b := range bin {
		res = append(res, intToHex(int(b>>4)))
		res = append(res, intToHex(int(b&0x0f)))
	}
	return string(res)
}

func intToHex(i int) rune {
	if i < 10 {
		return rune(i + '0')
	} else {
		return rune(i + 'a' - 10)
	}
}

func strtobin(s string) ([]uint8, error) {
	res := make([]uint8, len(s))
	for i, c := range s {
		res[i] = uint8(c)
	}
	return res, nil
}

func bintostr(bin []uint8) string {
	res := make([]rune, 0)
	for _, b := range bin {
		res = append(res, rune(b))
	}
	return string(res)
}

func uint8_be_to_64(bin []uint8) (*big.Int, error) {

	if len(bin) < 1 || len(bin) > 8 {
		return nil, fmt.Errorf("invalid input length (must be 1-8 bytes)")
	}

	var res big.Int = *big.NewInt(0)

	var twopow8 big.Int = *big.NewInt(2).Exp(big.NewInt(2), big.NewInt(8), nil)

	var i int = 0

	switch 9 - len(bin) {
	case 1:
		res = *res.Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 2:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 3:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 4:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 5:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 6:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 7:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	case 8:
		res = *res.Mul(&res, &twopow8).Add(&res, big.NewInt(int64(bin[i])))
		i++
	default:
		return nil, fmt.Errorf("impossible condition")
	}
	return &res, nil
}

func uint64_to_8be(i *big.Int, size int) ([]uint8, error) {
	if size < 1 || size > 8 {
		return nil, fmt.Errorf("invalid input length (must be 1-8 bytes)")
	}
	var res []uint8 = make([]uint8, size)
	var twopow8 big.Int = *big.NewInt(2).Exp(big.NewInt(2), big.NewInt(8), nil)

	for j := size - 1; j >= 0; j-- {
		res[j] = uint8(i.Mod(i, &twopow8).Int64())
		i = i.Div(i, &twopow8)
	}
	return res, nil
}

func encode_block(data []uint8, buf []uint8, index int) ([]uint8, error) {
	if len(data) < 1 || len(data) > full_encoded_block_size {
		return nil, fmt.Errorf("invalid input length (must be 1-%d bytes)", full_encoded_block_size)
	}

	num, _ := uint8_be_to_64(data)

	i := encoded_block_sizes[len(data)] - 1

	//while num > 0
	for num.Cmp(big.NewInt(0)) == 1 {
		buf[index+i] = uint8(num.Mod(num, big.NewInt(256)).Int64())
		num = num.Div(num, big.NewInt(256))
		i--
	}
	return buf, nil
}

func encode(hex string) ([]uint8, error) {
	if len(hex)%2 != 0 {
		return nil, fmt.Errorf("invalid hex string (must be even length)")
	}
	data, err := hextobin(hex)
	if err != nil {
		return nil, err
	}
	buf := make([]uint8, 0)
	for i := 0; i < len(data); i += encoded_block_sizes[len(data[i:])] {
		buf, err = encode_block(data[i:], buf, len(buf))
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func indexOf(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func indexOfUint8(arr []uint8, val uint8) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func decode_block(data []uint8, buf []uint8, index int) ([]uint8, error) {
	if len(data) < 1 || len(data) > full_encoded_block_size {
		return nil, fmt.Errorf("invalid input length (must be 1-%d bytes)", full_encoded_block_size)
	}
	// res_size = index of encoded_block_sizes[len(data)]

	res_size := indexOf(encoded_block_sizes, len(data))

	if res_size <= 0 {
		return nil, fmt.Errorf("invalid block size (must be 1-%d bytes)", full_encoded_block_size)
	}

	res_num := big.NewInt(0)
	order := big.NewInt(1)

	for i := len(data) - 1; i >= 0; i-- {
		digit := indexOfUint8(ALPHABET, data[i])
		if digit < 0 {
			return nil, fmt.Errorf("invalid symbol")
		}
		product := order.Mul(order, big.NewInt(int64(digit)))
		product = product.Add(product, res_num)

		if product.Cmp(UINT64_MAX) == 1 {
			return nil, fmt.Errorf("overflow")
		}

		res_num = product
		order = order.Mul(order, big.NewInt(int64(alphabet_size)))
	}

	numb := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(int64(res_size*8)), nil)

	if res_size < full_block_size && numb.Cmp(res_num) <= 0 {
		return nil, fmt.Errorf("overflow 2")
	}

	// Convert JavaScript code to Go code: buf.set(uint64_to_8be(res_num, res_size), index)
	buf, err := uint64_to_8be(res_num, res_size)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func decode(enc string) (string, error) {
	encbin, _ := strtobin(enc)
	if len(encbin) == 0 {
		return "", nil
	}

	full_block_count := math.Floor(float64(len(encbin) / full_encoded_block_size))
	last_block_size := len(encbin) % full_encoded_block_size
	last_block_decoded_size := indexOf(encoded_block_sizes, last_block_size)
	if last_block_decoded_size <= 0 {
		return "", fmt.Errorf("invalid encoded length")
	}

	data_size := full_block_count*float64(full_block_size) + float64(last_block_decoded_size)
	data := make([]uint8, int(data_size))
	for i := 0; i < int(full_block_count); i++ {
		data, _ = decode_block(subarray(encbin, i*full_encoded_block_size, i*full_encoded_block_size+full_encoded_block_size), data, i*full_block_size)
	}

	if last_block_size > 0 {
		data, _ = decode_block(subarray(encbin, int(full_block_count)*full_encoded_block_size, int(full_block_count)*full_encoded_block_size+last_block_size), data, int(full_block_count)*full_block_size)
	}

	return bintohex(data), nil
}

func subarray(arr []uint8, start int, end int) []uint8 {
	return arr[start:end]
}
