package xtz_validator

import (
	"github.com/MixinNetwork/mixin/domains/tezos"
	//sha256p "github.com/btwhite/go-btw-photon/crypto/sha256"
	//base58 "github.com/btcsuite/btcutil/base58"
)

func IsValidAddress(address string) bool {
	return tezos.VerifyAddress(address) == nil
}

/*func IsValidAddress(address string) bool {
	decoded := base58.Decode(address)
	payload := decodeRaw(decoded)
	if payload == nil {
		return false
	}
	x := payload[3:]
	fmt.Println(x)

	return decoded != nil && len(decoded) == 27
}

func decodeRaw(buffer []byte) []byte {
	payload := buffer[:len(buffer)-4]
	checksum := buffer[len(buffer)-4:]
	newChecksum := hexStr2byteArray(sha256x2(byteArray2hexStr(payload)))

	if checksum[0]^newChecksum[0]|checksum[1]^newChecksum[1]|checksum[2]^newChecksum[2]|checksum[3]^newChecksum[3] == 1 {
		return nil
	}

	return payload

}

/*
function decodeRaw(buffer) {
    let payload = buffer.slice(0, -4);
    let checksum = buffer.slice(-4);
    let newChecksum = cryptoUtils.hexStr2byteArray(
        cryptoUtils.sha256x2(cryptoUtils.byteArray2hexStr(payload))
    );

    if (checksum[0] ^ newChecksum[0] |
        checksum[1] ^ newChecksum[1] |
        checksum[2] ^ newChecksum[2] |
        checksum[3] ^ newChecksum[3])
        return;
    return payload;
}


// converts byte array to hex string
func byteArray2hexStr(byteArray []byte) string {
	var hexStr string
	for _, b := range byteArray {
		hexStr += fmt.Sprintf("%02x", b)
	}
	return hexStr
}

// converts hex string to byte array
func hexStr2byteArray(hexStr string) []byte {
	byteArray := make([]byte, len(hexStr)/2)
	for i := 0; i < len(hexStr); i += 2 {
		byteArray[i/2] = byte(hexStr[i])
	}
	return byteArray
}

// sha256x2
func sha256x2(hexStr string) string {
	return sha256(sha256(hexStr))
}

// sha256
func sha256(hexStr string) string {
	return fmt.Sprintf("%x", sha256p.Sha256([]byte(hexStr)))
}*/
