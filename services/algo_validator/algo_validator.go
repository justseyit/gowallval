//NOT COMPLETED

package algo_validator

import (
	"bytes"
	"crypto/sha512"
	"encoding/base32"
	"fmt"
	"strconv"
	"strings"
)

const algorand_checksum_byte_length = 4
const algorand_address_length = 58

func IsValidAddress(address string) bool {
	return verifyChecksum(address)
}

func correctPadding(address string) string {
	if len(address)%8 == 0 {
		println("len address: ", len(address))
		return strconv.Itoa(len(address))
	}
	println("address + (8 - len(address)%8): ", address+strings.Repeat("=", 8-len(address)%8))
	return address + strings.Repeat("=", 8-len(address)%8)
}

func verifyChecksum(address string) bool {
	if len(address) != algorand_address_length {
		return false
	}
	//Decode base32 address
	//dcd, _ := base32.StdEncoding.DecodeString(correctPadding(address))
	decoded := bytes.NewBuffer([]byte(decode(correctPadding(address))))
	println("decoded: ", decoded)
	//println("decoded.Bytes(): ", decoded.Bytes())

	//Assign addr with the 0 to decoded.length - 4 of variable decoded
	addr := decoded.Bytes()[:len(decoded.Bytes())-algorand_checksum_byte_length]
	println("addr: ", addr)

	//Assign checksum with the last 4 of variable decoded
	checksum := decoded.Bytes()[len(decoded.Bytes())-algorand_checksum_byte_length:]
	println("checksum: ", checksum)

	//Convert checksum to hex
	checksumHex := byteToHex(checksum)
	//println(checksumHex)

	hash := sha512.New()
	hash.Write(addr)

	//Convert hash to hex

	code := byteToHex(hash.Sum(nil))
	//Get last 8 characters of code
	code = code[len(code)-8:]
	println("code: ", code)
	println("checksumHex: ", checksumHex)
	return code == checksumHex
}

// convert byte array to hex string
func byteToHex(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func println(vname string, s interface{}) {
	fmt.Println(vname, s)
}

// Decode string as base32
func decode(s string) string {
	dcd, _ := base32.StdEncoding.DecodeString(correctPadding(s))
	println("dcd: ", dcd)
	return string(dcd)
}
