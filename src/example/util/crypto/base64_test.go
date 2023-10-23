package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {

	str := "心怀不惧 方能翱翔于天际"
	encodingString := Base64EncodingString(str)
	decodeString := Base64DecodeString(encodingString)
	fmt.Println("base64 编码后", encodingString)
	fmt.Println("base64 解码后", decodeString)
}

func Base64EncodingString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64DecodeString(str string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	return string(result)
}
