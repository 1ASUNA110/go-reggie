package api

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/global"
	"go-reggie/internal/model/vo/response"
	"io"
	"net/http"
)

type CommonApi struct {
}

var commonApi *CommonApi

func NewCommonApi() *CommonApi {
	if commonApi == nil {
		commonApi = &CommonApi{}
	}

	return commonApi
}

func (m *CommonApi) FileUpload(c *gin.Context) {
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

	// 3、保存文件
	fileName, err := global.MinioClient.FileUpload(file)

	if err != nil {
		response.Fail(response.FILE_UPLOAD_ERROR(), c)
		return
	}

	// 4、返回响应
	response.Ok(fileName, c)
}

// FileDownload 文件下载
func (m *CommonApi) FileDownload(c *gin.Context) {

	// 1、获取文件名
	fileName := c.Query("name")

	// 2、尝试下载文件
	minioObject, err := global.MinioClient.FileDownload(fileName)

	if err != nil {
		// 如果下载失败，返回错误图片
		m.returnErrorImage(c, "images/2e5ef4d4-bf83-4144-9d91-5ada7bdf2352.jpg")
		return
	}
	defer minioObject.Close()

	// 3、获取文件信息
	stat, err := minioObject.Stat()
	if err != nil {
		// 如果获取文件信息失败，返回错误图片
		m.returnErrorImage(c, "images/2e5ef4d4-bf83-4144-9d91-5ada7bdf2352.jpg")
		return
	}

	// 4、返回文件
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", stat.ContentType)
	//c.Header("Content-Length", string(stat.Size))

	// 将文件内容写入到响应中
	_, err = io.Copy(c.Writer, minioObject)
	if err != nil {
		// 如果写入响应失败，返回错误图片
		m.returnErrorImage(c, "images/2e5ef4d4-bf83-4144-9d91-5ada7bdf2352.jpg")
		return
	}
}

// returnErrorImage 返回指定的错误图片
func (m *CommonApi) returnErrorImage(c *gin.Context, errorImagePath string) {
	// 尝试下载错误图片
	errorObject, err := global.MinioClient.FileDownload(errorImagePath)
	if err != nil {
		// 如果错误图片下载失败，返回 500 错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send error image"})
		return
	}
	defer errorObject.Close()

	// 获取错误图片信息
	stat, err := errorObject.Stat()
	if err != nil {
		// 如果获取错误图片信息失败，返回 500 错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve error image information"})
		return
	}

	// 设置响应头，返回错误图片
	c.Header("Content-Disposition", "inline; filename="+errorImagePath)
	c.Header("Content-Type", stat.ContentType)
	c.Header("Content-Length", string(stat.Size))

	// 将错误图片内容写入到响应中
	_, err = io.Copy(c.Writer, errorObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send error image"})
	}
}
