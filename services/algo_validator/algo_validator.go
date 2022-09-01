package algo_validator

import (
	"bytes"
	"crypto/sha512"
	"encoding/base32"
	"fmt"

	goalgo "github.com/algorand/go-algorand-sdk/types"
)

const algorand_checksum_byte_length = 4
const algorand_address_length = 58
const public_key_length = 32

func IsValidAddress(address string) bool {
	_, err := goalgo.DecodeAddress(address)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
	//return verifyChecksum(address)
}

/*func correctPadding(address string) string {
	if len(address)%8 == 0 {
		println("len address: ", len(address))
		return strconv.Itoa(len(address))
	}
	return address + strings.Repeat("=", 8-len(address)%8)
}

func verifyChecksum(address string) bool {
	if len(address) != algorand_address_length {
		return false
	}

	// Decode base32 Address
	parsed := base32parse(correctPadding(address))
	println("parsed: ", parsed)

	decoded := bytes.NewBuffer(parsed)
	println("decoded: ", decoded)

	addr := bytes.NewBuffer(decoded.Bytes()[:len(decoded.Bytes())-algorand_checksum_byte_length])
	println("addr: ", addr)

	// Calculate Checksum
	checksum := bytes.NewBuffer(decoded.Bytes()[len(decoded.Bytes())-algorand_checksum_byte_length:])
	println("checksum: ", checksum)

	//Convert checksum to hex
	checksumHex := byteToHex(checksum.Bytes())

	//Hash Address and Checksum
	hash := sha512.Sum512(addr.Bytes())
	println("hash: ", hash)

	code := byteToHex(hash[:])
	code = code[:len(code)-8]
	println("code: ", code)
	println("checksumHex: ", checksumHex)

	//hash := sha512.Sum512(append(addr.Bytes(), base32parse(checksumHex)...))

	/*decoded := bytes.NewBuffer(base32parse(correctPadding(address)))
	println("decoded: ", decoded)

	addr := decoded.Bytes()[:len(decoded.Bytes())-algorand_checksum_byte_length]
	println("addr: ", addr)

	checksum := decoded.Bytes()[len(decoded.Bytes())-algorand_checksum_byte_length:]
	println("checksum: ", checksum)

	checksumHex := byteToHex(checksum)

	hash := sha512.New()
	hash.Write(addr)

	code := byteToHex(hash.Sum(nil))
	code = code[len(code)-8:]
	println("code: ", code)
	println("checksumHex: ", checksumHex)*/ /*
	return code == checksumHex
}

func byteToHex(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func println(vname string, s interface{}) {
	fmt.Println(vname, s)
}

func base32parse(s string) []uint8 {
	dst := make([]uint8, 0)
	base32.StdEncoding.Encode(dst, []byte(s))
	return dst
}*/

func decodeAddress(address string) ([]uint8, error) {

	addressBytes, _ := base32.StdEncoding.DecodeString(address)

	if len(addressBytes) != public_key_length+algorand_checksum_byte_length-1 {
		return nil, fmt.Errorf("Invalid address length: %d. Must be: %d", len(addressBytes), public_key_length+algorand_checksum_byte_length)
	}

	publicKey := addressBytes[:public_key_length]
	checksum := addressBytes[public_key_length : public_key_length+algorand_checksum_byte_length]

	// Compute the expected checksum
	computedChecksum := sha512.Sum512_256(publicKey)

	//Check if computed checksum equals the checksum as list
	if !bytes.Equal(computedChecksum[:algorand_checksum_byte_length], checksum) {
		return nil, fmt.Errorf("Invalid checksum, bytes are not equal")
	}

	return publicKey, nil
}
