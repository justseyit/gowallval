package segwitaddr

import (
	"fmt"
	"strings"

	bech32 "github.com/btcsuite/btcutil/bech32"
)

func convertbits(data []uint8, frombits uint8, tobits uint8, pad bool) ([]uint8, error) {
	acc := uint8(0)
	bits := uint8(0)
	ret := make([]uint8, 0)
	maxv := uint8((1 << tobits) - 1)
	for p := 0; p < len(data); p++ {
		value := data[p]
		if value < 0 || (value>>frombits) != 0 {
			return nil, fmt.Errorf("invalid value")
		}

		acc = (acc << frombits) | value
		bits += frombits

		for bits >= tobits {
			bits -= tobits
			ret = append(ret, (acc>>bits)&maxv)
		}

		if pad {
			if bits > 0 {
				ret = append(ret, (acc<<(tobits-bits))&maxv)
			} else if bits >= frombits || ((acc<<(tobits-bits))&maxv) != 0 {
				return nil, fmt.Errorf("invalid value")
			}
		}

	}
	return ret, nil
}

func decode(hrp string, addr string) (int, bool, []uint8, error) {
	dhrp, slc, err := bech32.Decode(addr)
	if err != nil {
		return 0, false, nil, err
	}
	if dhrp != "" || dhrp != hrp || len(slc) < 1 || slc[0] > 16 {
		return 0, false, nil, fmt.Errorf("invalid input (exception at line 47)")
	}
	res, err := convertbits(slc[1:], 5, 8, false)
	if err != nil {
		return 0, false, nil, err
	}
	if slc[0] == 0 && len(res) != 20 && len(res) != 32 {
		return 0, false, nil, fmt.Errorf("invalid input (exception at line 54)")
	}
	return int(slc[0]), true, res, nil
}

// int array to uint8 array
func intArrayToUint8Array(arr []int) []uint8 {
	var ret []uint8
	for _, v := range arr {
		ret = append(ret, uint8(v))
	}
	return ret
}

func encode(hrp string, version int, program []uint8) (string, error) {
	cb, err := convertbits([]uint8{uint8(version)}, 8, 5, true)
	if err != nil {
		return "", err
	}
	ls := append([]int{version}, uint8ArrayToInt(cb))
	ret, err1 := bech32.Encode(hrp, intArrayToUint8Array(ls))
	if err1 != nil {
		return "", err1
	}
	_, succ, _, err := decode(hrp, ret)
	if !succ {
		return "", err
	}
	return ret, nil
}

// uint8 array to int
func uint8ArrayToInt(arr []uint8) int {
	var ret int
	for _, v := range arr {
		ret = ret<<8 + int(v)
	}
	return ret
}

func IsValidAddress(address string) bool {
	fmt.Println("segwit.IsValidAddress:", address)
	hrp := "bc"
	ver, succ, ret, _ := decode(hrp, address)
	if ret == nil {
		hrp = "tb"
		_, _, ret, _ = decode(hrp, address)
	}

	if ret == nil {
		return false
	}

	recreate, err := encode(hrp, ver, ret)
	if err != nil {
		fmt.Println("Error at IsValidAddress: ", err)
	}
	fmt.Println("recreate:", recreate)
	fmt.Println("address:", address)
	fmt.Println("hrp:", hrp)
	fmt.Println("ver:", ver)
	fmt.Println("ret:", ret)
	fmt.Println("succ:", succ)
	return recreate == strings.ToLower(address)
}
