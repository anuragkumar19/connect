package token

import (
	"crypto/rand"
	"encoding/base64"
)

func Opaque() (string, error) {
	var b [48]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b[:]), nil
}
