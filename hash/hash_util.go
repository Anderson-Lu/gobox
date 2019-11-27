package hash

import "hash/crc32"

//根据key计算返回[0:limitSize)的散列随机值
func HashKey(key string, limitSize int) uint32 {
	return crc32.ChecksumIEEE([]byte(key)) % uint32(limitSize)
}
