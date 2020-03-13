package gb

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/natefinch/lumberjack"
	"log"
	"time"
	"fmt"
)

var Logger *zap.SugaredLogger


//获取日志
//filePath日志文件路径
//level 日志级别
//maxSize 每个日志文件保存的最大尺寸
//maxBackups日志文件最多保存多少个备份
//maxAge文件最多保存多少填
//compress是否压缩
//serviceName服务名

func LogConf(){
	now := time.Now()
	hook := &lumberjack.Logger{
		Filename: fmt.Sprintf("log/%007", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()),
		MaxSize: 5000,
		MaxBackups: 10,
		Compress: false,
	}
	defer hook.Close()

	enConfig := zap.NewProductionEncoderConfig()

	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		zapcore.AddSync(hook),
		zap.InfoLevel,
		)
	logger := zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1))
	_log := log.New(hook,"",log.LstdFlags)
	Logger = logger.Sugar()
	_log.Println("Start...")
}