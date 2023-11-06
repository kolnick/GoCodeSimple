package encrypt

import (
	"fmt"
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
