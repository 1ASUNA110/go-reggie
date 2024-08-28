基于黑马程序员的瑞吉外卖项目 魔改的 go语言版瑞吉外卖项目
使用到的技术栈
gin + gorm + jwt + redis + mysql



~~~
# 安装 GORM
go get -u gorm.io/gorm
# 使用 mysql 作为数据库
go get -u gorm.io/driver/mysql  

# 安装 gin
go get -u github.com/gin-gonic/gin

# 安装 Viper
go get -u github.com/spf13/viper

# 安装 gin session 管理库
go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/cookie

# 安装 sonyflake 主要为了数据库主键的雪花id
go get github.com/sony/sonyflake

~~~