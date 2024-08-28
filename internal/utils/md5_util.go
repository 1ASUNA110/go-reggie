package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash 对输入字符串进行 MD5 加密并返回十六进制字符串
func MD5Hash(input string) string {
	// 创建 MD5 哈希对象
	hash := md5.New()

	// 写入要加密的数据
	hash.Write([]byte(input))

	// 获取 MD5 哈希值（字节数组）
	hashBytes := hash.Sum(nil)

	// 将字节数组转换为十六进制字符串
	return hex.EncodeToString(hashBytes)
}
