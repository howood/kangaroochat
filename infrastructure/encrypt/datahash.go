package encrypt

import (
	"encoding/base64"

	"golang.org/x/crypto/sha3"
)

//DataHash struct
type DataHash struct {
}

//GetHash get hashed data
func (dh DataHash) GetHash(data string) string {
	// A MAC with 64 bytes of output has 512-bit security strength
	h := make([]byte, 64)
	d := sha3.NewShake256()
	d.Write([]byte(data))
	d.Read(h)
	return base64.URLEncoding.EncodeToString(h)
}
