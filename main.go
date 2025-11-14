package main

import (
	"blog/common"
	"blog/entity"
	"blog/route"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var configFile = flag.String("f", "config.yml", "the config file")

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//初始化数据库
	initDB()

	//初始化路由
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	route.InitRoute(router)
	router.Run("localhost:8080")
}

func initDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略找不到记录的错误
			Colorful:                  true,        // 彩色输出
		},
	)
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"),
		viper.GetInt("db.port"), viper.GetString("db.dbname"))
	common.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动创建和迁移表
	err = common.Db.AutoMigrate(&entity.User{})
	// 这里的db就是上述gorm.open的db哈
	if err != nil {
		panic("User 创建/迁移表格失败, error = " + err.Error())
	}

	err = common.Db.AutoMigrate(&entity.Post{})
	// 这里的db就是上述gorm.open的db哈
	if err != nil {
		panic("Post 创建/迁移表格失败, error = " + err.Error())
	}

	err = common.Db.AutoMigrate(&entity.Comment{})
	// 这里的db就是上述gorm.open的db哈
	if err != nil {
		panic("Comment 创建/迁移表格失败, error = " + err.Error())
	}
}
