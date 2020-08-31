package middleware

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Logger() gin.HandlerFunc {
	// 1
	var fileName string
	_, err := os.Stat(viper.GetString("log.path"))
	if err != nil {
		fileName = path.Join(".", viper.GetString("log.name"))
	} else {
		fileName = path.Join(viper.GetString("log.path"), viper.GetString("log.name"))
	}
	// 2
	//f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := io.MultiWriter(f, os.Stdout)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(writer)

	// 3
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+"%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{

		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.InfoLevel:  logWriter,
		log.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	log.AddHook(lfHook)

	// 4
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		log.WithFields(log.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURI,
		}).Info()
	}
}
