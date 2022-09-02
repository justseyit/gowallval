package lumen_validator

import (
	"bytes"
	"regexp"

	buffer "github.com/arloliu/jsonpack/buffer"
	basex "github.com/hailongz/golang/basex"
)

const allowedChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

const regEx string = "^G[" + allowedChars + "]{55}$"

func IsValidAddress(address string) bool {
	match, err := regexp.MatchString(regEx, address)
	if err != nil {
		return false
	}
	return verifyChecksum(address) && match
}

func verifyChecksum(address string) bool {
	codec, _ := basex.NewEncoding(allowedChars)
	decoded, _ := codec.Decode(address)
	versionByte := decoded[0]
	checksum := decoded[len(decoded)-2:]
	payload := decoded[:len(decoded)-2]

	if address != codec.Encode(decoded) {
		return false
	}

	if versionByte != 6<<3 {
		return false
	}

	calculatedChecksum := buffer.Create(2)

	calculatedChecksum.WriteUint16LE(crc16xmodem(payload))

	return bytes.Equal(calculatedChecksum.Bytes(), checksum)
}

func crc16xmodem(data []byte) uint16 {
	var crc uint16 = 0x0000
	for _, v := range data {
		x := ((crc >> 8) ^ uint16(v)) & 0xff
		x ^= x >> 4
		crc = (crc << 8) ^ (x << 12) ^ (x << 5) ^ x
	}
	return crc
}
