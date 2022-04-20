package logger

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getLogLevel() (level zapcore.Level) {
	logLevelFromCfg := viper.GetString("LOGGER.LEVEL")
	logLevel := strings.ToLower(logLevelFromCfg)
	switch logLevel {
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal", "panic":
		level = zapcore.DPanicLevel
	default:
		level = zapcore.DebugLevel
	}
	return level
}

func getLogWriter() zapcore.WriteSyncer {
	var (
		maxSize    = viper.GetInt("LOGGER.MAX_SIZE_IN_MB")
		maxBackups = viper.GetInt("LOGGER.MAX_BACKUP_FILES")
		maxAge     = viper.GetInt("LOGGER.MAX_DAY_TO_KEEP_BACKUP_FILES")
		compress   = viper.GetBool("LOGGER.COMPRESS_BACKUP_FILE")
	)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./server.log",
		MaxSize:    maxSize,    // the maximum size of the log file (in MB) before cutting
		MaxBackups: maxBackups, // the maximum number of old files to keep
		MaxAge:     maxAge,     // the maximum number of days to keep old files
		Compress:   compress,   // compress / archive old files
	}
	writeSyncers := []zapcore.WriteSyncer{
		zapcore.AddSync(os.Stdout),
	}
	if viper.GetBool("LOGGER.ENABLE_FILE") {
		writeSyncers = append(
			writeSyncers, zapcore.AddSync(lumberJackLogger),
		)
	}
	return zapcore.NewMultiWriteSyncer(writeSyncers...)
}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	logInJSONFmt := viper.GetBool("LOGGER.JSON")
	if logInJSONFmt {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}

func LoggerWithRequestID(ctx echo.Context, logger *zap.SugaredLogger) *zap.SugaredLogger {
	requestId := ctx.Response().Header().Get(echo.HeaderXRequestID)
	return logger.With(zap.String("req_id", requestId))
}

func NewZapLogger() *zap.Logger {
	logLevel := getLogLevel()
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())

	defer logger.Sync()

	return logger
}

func New() *zap.SugaredLogger {
	logLevel := getLogLevel()
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller()).Sugar()

	defer logger.Sync()

	return logger
}
