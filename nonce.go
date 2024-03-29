// package nonce creates random strings.
package nonce

import (
	"crypto/rand"
	"encoding/base64"
	"math"
	"strings"
	"unicode"
)

// New returns a random, alpha-numeric string of length l, or an error if the
// random number generator fails.
func New(l int) (string, error) {
	buf := make([]byte, int(math.Ceil(float64(l)*4/3))+2)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return strings.Map(alphaNumeric, base64.RawURLEncoding.EncodeToString(buf))[:l], nil
}

// Must is a wrapper around New but panics if New returns an error.
func Must(l int) string {
	n, err := New(l)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func alphaNumeric(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return r
	}
	return -1
}
