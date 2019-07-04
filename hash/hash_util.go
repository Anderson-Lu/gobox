package hash

import "hash/crc32"

//根据key计算返回[0:limitSize)的散列随机值
func HashKey(key string, limitSize int) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)]) % uint32(limitSize)
	}
	return crc32.ChecksumIEEE([]byte(key)) % uint32(limitSize)
}
