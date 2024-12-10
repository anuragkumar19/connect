package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
)

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomToken() (string, error) {
	b, err := generateRandomBytes(48)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func RandomNumber(length int) (string, error) {
	max := math.Pow10(length) - 1
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("%06d\n", n.Int64())
	return str, nil
}
