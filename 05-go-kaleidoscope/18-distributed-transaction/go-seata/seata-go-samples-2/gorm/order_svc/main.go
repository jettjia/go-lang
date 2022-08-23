package main

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentrx/mysql/v2"
	"github.com/opentrx/seata-go-samples/order_svc/dao"
	"github.com/opentrx/seata-golang/v2/pkg/client"
	"github.com/opentrx/seata-golang/v2/pkg/client/config"
	"github.com/opentrx/seata-golang/v2/pkg/client/rm"
	"github.com/opentrx/seata-golang/v2/pkg/util/log"

	dialector "github.com/opentrx/seata-go-samples/dialector/mysql"
)

func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	conf := config.InitConfiguration(configPath)

	log.Init(conf.Log.LogPath, conf.Log.LogLevel)
	client.Init(conf)

	rm.RegisterTransactionServiceServer(mysql.GetDataSourceManager())
	mysql.RegisterResource(config.GetATConfig().DSN)

	db, err := gorm.Open(
		dialector.Open(config.GetATConfig().DSN),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}})
	if err != nil {
		panic(err)
	}
	DB, err := db.DB()
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(4 * time.Hour)
	if err := DB.Ping(); err != nil {
		panic(err)
	}
	d := &dao.Dao{DB: db}

	r.POST("/createSo", func(c *gin.Context) {
		type req struct {
			Req []*dao.SoMaster
		}
		var q req
		if err := c.ShouldBindJSON(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := d.CreateSO(
			context.WithValue(
				context.Background(),
				mysql.XID,
				c.Request.Header.Get("XID")),
			q.Req)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		}
	})
	r.Run(":8002")
}
