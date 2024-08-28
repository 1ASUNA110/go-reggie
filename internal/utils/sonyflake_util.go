package utils

import (
	"github.com/sony/sonyflake"
	"log"
)

var sf *sonyflake.Sonyflake

func init() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		log.Fatal("Sonyflake 初始化失败")
	}
}

// GenerateID 生成一个唯一的 Sonyflake ID
func GenerateID() int64 {
	id, err := sf.NextID()
	if err != nil {
		log.Fatalf("Sonyflake ID 生成失败: %v", err)
	}
	return int64(id)
}
