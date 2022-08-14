//NOT COMPLETED

package cardano_validator

import (
	"bytes"
	//base58 "gowallval/utils/base58"
)

/*func getDecoded(address string) (string, error) {
	decoded, err := base58.Decode(address)
	if err != nil {
		return "", err
	} else {
		cdcd := cbor.NewDecoder(bytes.NewReader(intArrayToByteArray(decoded)))
		d := cdcd.Decode(decoded)
		return cdcd.Decode(), nil
	}

}*/

func intArrayToByteArray(arr []int) []byte {
	var buf bytes.Buffer
	for _, v := range arr {
		buf.WriteByte(byte(v))
	}
	return buf.Bytes()
}
