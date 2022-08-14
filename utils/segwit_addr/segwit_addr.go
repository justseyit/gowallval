package segwitaddr

import (
	"fmt"
	bech32 "gowallval/utils/bech32"
	"strings"
)

//convert this JavaScript code to Go code:
/*
function convertbits (data, frombits, tobits, pad) {
  var acc = 0
  var bits = 0
  var ret = []
  var maxv = (1 << tobits) - 1
  for (var p = 0; p < data.length; ++p) {
    var value = data[p]
    if (value < 0 || (value >> frombits) !== 0) {
      return null
    }
    acc = (acc << frombits) | value
    bits += frombits
    while (bits >= tobits) {
      bits -= tobits
      ret.push((acc >> bits) & maxv)
    }
  }
  if (pad) {
    if (bits > 0) {
      ret.push((acc << (tobits - bits)) & maxv)
    }
  } else if (bits >= frombits || ((acc << (tobits - bits)) & maxv)) {
    return null
  }
  return ret
}
*/

func convertbits(data []uint8, frombits uint8, tobits uint8, pad bool) ([]uint8, error) {
	acc := 0
	bits := 0
	ret := []uint8{}
	maxv := (1 << tobits) - 1
	for _, value := range data {
		if value < 0 || (value>>frombits) != 0 {
			return nil, fmt.Errorf("invalid input")
		}
		acc = int(acc<<frombits) | int(value)
		bits += int(frombits)
		for bits >= int(tobits) {
			bits -= int(tobits)
			ret = append(ret, uint8(acc>>bits)&uint8(maxv))
		}
	}
	if pad {
		if bits > 0 {
			ret = append(ret, (uint8(acc)<<(tobits-uint8(bits)))&uint8(maxv))
		}
	} else if uint8(bits) >= frombits || ((acc<<(tobits-uint8(bits)))&maxv) != 0 {
		return nil, fmt.Errorf("overflow")
	}
	return ret, nil
}

//convert this JavaScript code to Go code:
/*
function decode (hrp, addr) {
  var dec = bech32.decode(addr)
  if (dec === null || dec.hrp !== hrp || dec.data.length < 1 || dec.data[0] > 16) {
    return null
  }
  var res = convertbits(dec.data.slice(1), 5, 8, false)
  if (res === null || res.length < 2 || res.length > 40) {
    return null
  }
  if (dec.data[0] === 0 && res.length !== 20 && res.length !== 32) {
    return null
  }
  return { version: dec.data[0], program: res }
}
*/

func Decode(hrp string, addr string) (int, []uint8, error) {
	_, slc, dhrp, err := bech32.Decode(addr)
	if err != nil {
		return 0, nil, err
	}
	if dhrp != hrp || len(slc) < 1 || slc[0] > 16 {
		return 0, nil, fmt.Errorf("invalid input")
	}
	res, err := convertbits(intArrayToUint8Array(slc[1:]), 5, 8, false)
	if err != nil {
		return 0, nil, err
	}
	if slc[0] == 0 && len(res) != 20 && len(res) != 32 {
		return 0, nil, fmt.Errorf("invalid input")
	}
	return slc[0], res, nil
}

// int array to uint8 array
func intArrayToUint8Array(arr []int) []uint8 {
	var ret []uint8
	for _, v := range arr {
		ret = append(ret, uint8(v))
	}
	return ret
}

//convert this JavaScript code to Go code:
/*
function encode (hrp, version, program) {
  var ret = bech32.encode(hrp, [version].concat(convertbits(program, 8, 5, true)))
  if (decode(hrp, ret) === null) {
    return null
  }
  return ret
}
*/

func Encode(hrp string, version int, program []uint8) (string, error) {
	cb, _ := convertbits([]uint8{uint8(version)}, 8, 5, true)
	ret := bech32.Encode(hrp, append([]int{version}, uint8ArrayToInt(cb)))
	if _, d, _ := Decode(hrp, ret); d == nil {
		return "", fmt.Errorf("invalid input")
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

//convert this JavaScript code to Go code:
/*
function isValidAddress (address) {
  var hrp = 'bc'
  var ret = decode(hrp, address)

  if (ret === null) {
    hrp = 'tb'
    ret = decode(hrp, address)
  }

  if (ret === null) {
    return false
  }

  var recreate = encode(hrp, ret.version, ret.program)
  return recreate === address.toLowerCase()
}
*/

func IsValidAddress(address string) bool {
	hrp := "bc"
	ver, ret, _ := Decode(hrp, address)
	if ret == nil {
		hrp = "tb"
		_, ret, _ = Decode(hrp, address)
	}

	if ret == nil {
		return false
	}

	recreate, _ := Encode(hrp, ver, ret)
	return recreate == strings.ToLower(address)
}
