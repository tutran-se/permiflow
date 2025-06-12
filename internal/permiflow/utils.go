package permiflow

import (
	"crypto/rand"
	"encoding/hex"
)

func ShortRandomString(length int) string {
	b := make([]byte, length/2)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
