package comm

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger zerolog.Logger

func InitLogger() error {
	// 设置全局的日志级别为 debug
	level, err := zerolog.ParseLevel(CfgLoader.GetString("log.level"))
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	// 创建一个lumberjack.Logger对象，设置日志文件的参数
	logFile := &lumberjack.Logger{
		Filename:   CfgLoader.GetString("log.filename"), // 日志文件名
		MaxSize:    CfgLoader.GetInt("log.max_size"),    // 最大大小（MB）
		MaxBackups: CfgLoader.GetInt("log.max_backups"), // 最大备份数
		MaxAge:     CfgLoader.GetInt("log.max_age"),     // 最大保留天数
		Compress:   true,                                // 是否压缩
	}

	Logger = zerolog.New(logFile).With().Timestamp().Caller().Logger().Level(zerolog.DebugLevel)
	Logger.Info().Str("EventName", "Init").Msg("初始化日志")

	return nil
}
