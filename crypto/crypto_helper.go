package crypto_helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

//HMACSHA256加密算法
func HmacSha256(data string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//SHA256加密算法
func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}
