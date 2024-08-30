package config

import (
	"github.com/spf13/viper"
	"go-reggie/internal/utils"
)

func InitMinio() (*utils.MinIOClient, error) {
	// 初始化 MinIO 客户端
	endpoint := viper.GetString("minio.endpoint")
	accessKeyID := viper.GetString("minio.accessKeyId")
	secretAccessKey := viper.GetString("minio.secretAccessKey")
	bucketName := viper.GetString("minio.bucketName")
	useSSL := viper.GetBool("minio.useSSL")

	minioClient, err := utils.NewMinIOClient(endpoint, accessKeyID, secretAccessKey, bucketName, useSSL)

	return minioClient, err

}
