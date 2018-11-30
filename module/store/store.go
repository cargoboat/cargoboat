package store

import (
	"log"
	"os"

	"github.com/dkeng/pkg/logger"
	"github.com/jinzhu/gorm"
	// postgres驱动
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DB 数据库连接
var DB *gorm.DB

func initDB() {
	// 初始化数据库
	db, err := gorm.Open("postgres", viper.GetString("postgres.address"))
	if err != nil {
		logger.Fatalf(
			"初始化 postgres 连接失败: %s \n",
			errors.Wrap(err, "打开 postgres 连接失败"),
		)
		os.Exit(-1)
	}
	err = db.DB().Ping()
	if err != nil {
		logger.Fatalf(
			"初始化 postgres 连接失败: %s \n",
			errors.Wrap(err, "Ping postgres 失败"),
		)
		os.Exit(-1)
	}

	db.LogMode(viper.GetBool("postgres.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("postgres.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("postgres.max_idle"))
	// db.DB().SetConnMaxLifetime(time.Hour)

	DB = db
}

// Start 启动存储
func Start() {
	initDB()
}

// Close 关闭
func Close() {
	err := DB.Close()
	if err != nil {
		log.Println(err)
	}
}
