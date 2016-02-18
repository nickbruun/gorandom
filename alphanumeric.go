package random

import (
	"crypto/rand"
	"encoding/binary"
)

// Alphanumeric characters.
const alphanumericChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Alphanumeric bits of random information needed per character.
const alphanumericBits = 6

// Alphanumeric bit mask for random information.
const alphanumericMask = (1 << alphanumericBits) - 1

// Generate a cryptographically strong random alphanumeric string.
//
// An error in generating a string should be considered an exceptional case.
func RandAlphanumericString(n int) (string, error) {
	src := make([]byte, 8)
	b := make([]byte, n)

	pos := 0
	for pos < n {
		// Read 64 bits of information to use as a source.
		if _, err := rand.Read(src); err != nil {
			return "", err
		}

		x := binary.BigEndian.Uint64(src)

		// Attempt to fill into the reandom string.
		for xPos := 0; xPos < 64 / alphanumericBits && pos < n; xPos++ {
			idx := int(x & alphanumericMask)
			x = x >> alphanumericBits

			if idx < len(alphanumericChars) {
				b[pos] = alphanumericChars[idx]
				pos++
			}
		}
	}

	return string(b), nil
}
