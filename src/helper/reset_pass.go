package helper

import (
	"crypto/rand"
	"fmt"
)

func ResetPass() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
