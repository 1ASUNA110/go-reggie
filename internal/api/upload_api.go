package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-reggie/internal/utils/response"
	"path/filepath"
)

type UploadApi struct {
}

var uploadApi *UploadApi

func NewUploadApi() *UploadApi {
	if uploadApi == nil {
		uploadApi = &UploadApi{}
	}

	return uploadApi
}

func (m *UploadApi) Upload(c *gin.Context) {
	// 1、获取文件
	file, err := c.FormFile("file")

	if err != nil {

		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、验证文件类型是否合法
	// 2.1、获取文件类型
	contentType := file.Header.Get("Content-Type")

	if contentType != "image/jpeg" && contentType != "image/png" {
		response.Fail(response.UOLOAD_FILE_TYPE_ERROR(), c)
		return
	}

	// 3、获取文件后缀
	ext := filepath.Ext(file.Filename)

	// 4、生成UUID重新生成文件名，防止文件名称重复造成文件覆盖
	fileName := uuid.New().String() + ext

	// 5、保存文件
	// todo

	// 6、返回响应
	response.Ok(fileName, c)
}
