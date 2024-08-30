package global

import (
	"go-reggie/internal/utils"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	MinioClient *utils.MinIOClient
)
