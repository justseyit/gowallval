package base58

import "fmt"

const ALPHABET string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var ALPHABET_MAP map[string]int = make(map[string]int)

func Init() {
	for i, c := range ALPHABET {
		ALPHABET_MAP[string(c)] = i
	}
}

const BASE = len(ALPHABET)

func Decode(input string) ([]int, error) {

	if len(input) == 0 {
		return nil, fmt.Errorf("invalid input (empty string)")
	}

	bytes := make([]int, 0)

	for _, i := range input {
		//Check if the character is in the alphabet
		if _, ok := ALPHABET_MAP[string(i)]; !ok {
			return nil, fmt.Errorf("non-base58 character in input: %v", i)
		}

		for j := 0; j < len(bytes); j++ {
			bytes[j] *= BASE
		}

		bytes[0] += ALPHABET_MAP[string(i)]

		carry := 0
		for j := 0; j < len(bytes); j++ {
			bytes[j] += carry

			carry = bytes[j] >> 8
			bytes[j] &= 0xff
		}

		for carry > 0 {
			bytes = append(bytes, carry&0xff)

			carry >>= 8
		}

		//Remove leading 0s
		for bytes[0] == 0 && len(bytes) > 1 {
			bytes = bytes[1:]
		}

	}

	//return bytes in reverse order
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return bytes, nil
}
