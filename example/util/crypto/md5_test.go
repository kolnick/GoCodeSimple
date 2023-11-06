package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("hello")
	hash := md5.Sum(data)
	fmt.Println(hex.EncodeToString(hash[:]))
}
