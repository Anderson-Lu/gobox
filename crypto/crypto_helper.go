package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
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

//Md5加密
func GetMD5Hex(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	hexstr := hex.EncodeToString(h.Sum(nil))
	return hexstr
}
