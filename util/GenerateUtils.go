package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成一个随机的16位字符串
func GenerateId(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	unionId := make([]byte, length)
	for i := range unionId {
		unionId[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(unionId)
}
