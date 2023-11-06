package encrypt

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"hash/crc32"
	"testing"
)

func TestCRC32(t *testing.T) {

	v1 := stringToInt("test")
	v2 := stringToInt("test")
	fmt.Println(v1)
	fmt.Println(v2)
}

func stringToInt(strValue string) int {
	v := int(crc32.ChecksumIEEE([]byte(strValue)))
	return v
}

func TestBase64(t *testing.T) {

	inputStr := "this is string"
	bytes := []byte(inputStr)
	encoded := base64.RawStdEncoding.EncodeToString(bytes)
	decodedData, err := base64.RawStdEncoding.DecodeString(encoded)
	assert.Nil(t, err)

	fmt.Println(encoded)
	fmt.Println(decodedData)
}
