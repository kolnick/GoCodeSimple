package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	data := []byte("hello")
	hash := sha256.Sum256(data)
	fmt.Println(hex.EncodeToString(hash[:]))
}
