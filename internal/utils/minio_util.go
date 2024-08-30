package utils

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

// MinIOClient 是一个封装了 MinIO 客户端的结构体
type MinIOClient struct {
	Client     *minio.Client
	BucketName string
}

// NewMinIOClient 创建并返回一个 MinIOClient 实例
func NewMinIOClient(endpoint, accessKeyID, secretAccessKey, bucketName string, useSSL bool) (*MinIOClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &MinIOClient{
		Client:     client,
		BucketName: bucketName,
	}, nil
}

// FileUpload 是 MinIOClient 的成员方法，用于上传文件
func (m *MinIOClient) FileUpload(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	// 3、获取文件后缀
	ext := filepath.Ext(file.Filename)

	// 4、生成UUID重新生成文件名，防止文件名称重复造成文件覆盖
	fileName := generateFileName(ext)

	// 打开文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 获取文件的内容类型
	contentType := file.Header.Get("Content-Type")

	// 上传文件到 MinIO
	_, err = m.Client.PutObject(ctx, m.BucketName, fileName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	return fileName, nil
}

// FileDownload 是 MinIOClient 的成员方法，用于从 MinIO 下载文件并返回给客户端
func (m *MinIOClient) FileDownload(fileName string) (*minio.Object, error) {
	ctx := context.Background()

	// 从 MinIO 获取文件对象
	object, err := m.Client.GetObject(ctx, m.BucketName, fileName, minio.GetObjectOptions{})

	return object, err

}

// generateFileName 生成文件名
func generateFileName(ext string) string {
	// 获取当前时间
	now := time.Now()

	// 格式化日期为年/月/日
	datePath := now.Format("2006/01/02")

	// 生成 UUID
	fileName := uuid.New().String()

	// 组合路径和文件名
	fullPath := filepath.Join(datePath, fileName+ext)

	fullPath = strings.ReplaceAll(fullPath, "\\", "/")

	return fullPath
}
