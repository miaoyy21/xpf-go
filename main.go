package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"path/filepath"
	"time"
	"xpf/cache"
	"xpf/ds"
	"xpf/md"
	"xpf/pack"
)

func main() {
	// Init Cache
	logrus.SetLevel(logrus.DebugLevel)
	if err := cache.Init(); err != nil {
		logrus.Fatalf("cache.Init() Failure :: %s", err.Error())
		return
	}

	cfg := cache.Cache.Config

	// Connecting MySQL
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&loc=UTC&parseTime=true", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbDatabase)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("gorm.Open() Failure :: %s", err.Error())
		return
	}
	db.LogMode(true)

	/*************************************** Start ***************************************/
	//db.DropTable(&ds.User{})
	//db.DropTable(&ds.UserKey{})
	//db.DropTable(&ds.ApplicationAccount{})
	//db.DropTable(&ds.ApplicationAccountField{})
	//db.DropTable(&ds.ApplicationCategory{})
	//db.DropTable(&ds.File{})
	//db.DropTable(&ds.Operate{})
	//db.DropTable(&ds.Score{})
	//db.DropTable(&ds.Task{})
	//db.DropTable(&ds.Purchase{})
	/***************************************  End  ***************************************/

	// Init DataBase
	if err := ds.Init(db); err != nil {
		logrus.Fatalf("ds.Init() Failure :: %s", err.Error())
		return
	}

	// Init Tasks
	go md.Init(db)

	// Set Logrus && Gin && DataBase Mode
	if cfg.Mode == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
		db.LogMode(true)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
		db.LogMode(false)
	}
	logrus.Infof("Running with %s ...", cfg.Mode)

	router := gin.Default()

	// Routers
	router.StaticFS("www", http.Dir(filepath.Join("store", "www")))
	router.StaticFS("psw/icon", http.Dir(filepath.Join("store", "icon")))
	router.StaticFS("psw/file", http.Dir(filepath.Join("store", "file")))
	router.StaticFS("psw/share", http.Dir(filepath.Join("store", "share")))
	router.StaticFS("psw/thumbnail", http.Dir(filepath.Join("store", "thumbnail")))

	router.GET("psw/HelloWorld", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json", []byte(fmt.Sprintf("Hello World :: %s", time.Now().GoString())))
	})
	router.POST("psw/md", pack.Pack(db))

	// Listening
	logrus.Infof("Listen HTTP at Host %s and Post %s ...", cfg.Host, cfg.Port)
	if err := router.Run(net.JoinHostPort(cfg.Host, cfg.Port)); err != nil {
		logrus.Fatalf("GIN Run Failure :: %s", err.Error())
	}
}
