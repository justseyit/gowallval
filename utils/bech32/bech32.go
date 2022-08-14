package bech32

import (
	"fmt"
	"strings"
)

const CHARSET string = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

var GENERATOR []int = make([]int, len(CHARSET))

func init() {
	GENERATOR = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
}

func Polymod(values []int) int {
	chk := 1
	for _, value := range values {
		b := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ value
		for i := 0; i < 5; i++ {
			if (b>>i)&1 == 1 {
				chk ^= GENERATOR[i]
			}
		}
	}
	return chk
}

func HrpExpand(hrp string) []int {
	ret := make([]int, 0)
	for _, c := range hrp {
		ret = append(ret, int(c)>>5)
	}
	return ret
}

func VerifyChecksum(hrp string, data []int) bool {
	return Polymod(append(HrpExpand(hrp), data...)) == 1
}

func CreateChecksum(hrp string, data []int) []int {

	values := make([]int, 0)
	values = append(HrpExpand(hrp), data...)
	values = append(values, []int{0, 0, 0, 0, 0, 0}...)

	polymod := Polymod(values) ^ 1
	ret := make([]int, 0)
	for i := 0; i < 6; i++ {
		ret = append(ret, (polymod>>uint(5*(5-i)))&31)
	}

	return ret

}

func Encode(hrp string, data []int) string {
	combined := make([]int, 0)
	combined = append(data, CreateChecksum(hrp, data)...)

	ret := hrp + "1"
	for _, i := range combined {
		ret += string(CHARSET[combined[i]])
	}

	return ret
}

func Decode(input string) ([]int, []int, string, error) {

	var hasLower = false
	var hasUpper = false
	for _, c := range input {
		if int(c) < 33 || int(c) > 126 {
			return nil, nil, "", fmt.Errorf("invalid character in input: %v", c)
		}
		if int(c) >= 97 && int(c) <= 122 {
			hasLower = true
		}
		if int(c) >= 65 && int(c) <= 90 {
			hasUpper = true
		}
	}

	if hasLower && hasUpper {
		return nil, nil, "", fmt.Errorf("mixed case in input")
	}

	if !hasLower && !hasUpper {
		return nil, nil, "", fmt.Errorf("no case in input")
	}

	//convert input to lowercase
	input = strings.ToLower(input)

	//get last index of "1"
	pos := strings.LastIndex(input, "1")

	if pos < 1 || pos+7 > len(input) || len(input) > 90 {
		return nil, nil, "", fmt.Errorf("invalid input")
	}

	hrp := input[:pos]

	data := make([]int, 0)

	for _, c := range input[pos+1:] {
		d := strings.Index(CHARSET, string(CHARSET[input[c]]))
		if d < 0 {
			return nil, nil, "", fmt.Errorf("invalid input")
		}
		data = append(data, d)
	}
	if !VerifyChecksum(hrp, data) {
		return nil, nil, "", fmt.Errorf("invalid checksum")
	}
	slice := make([]int, len(data)-6)
	copy(slice, data[:len(data)-6])
	return data, slice, hrp, nil
}
