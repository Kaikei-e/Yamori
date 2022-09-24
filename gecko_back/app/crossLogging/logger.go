package crossLogging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

var Logger *zap.Logger

func LoggerSetup(path string) {
	fp := filepath.Join(path + "/server.log")

	//ws, err := os.OpenFile(fp, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}

	rotateLogger := &lumberjack.Logger{
		Filename:   fp,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}

	writerSyncer := zapcore.AddSync(rotateLogger)

	//writerSyncer := zapcore.AddSync(ws)

	conf := zap.NewProductionConfig()
	encoderConf := zap.NewProductionEncoderConfig()

	encoderConf = zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		MessageKey:     "message",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	conf.EncoderConfig = encoderConf
	conf.OutputPaths = []string{"stdout", path}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(conf.EncoderConfig),
		writerSyncer,
		zapcore.InfoLevel,
	)

	Logger = zap.New(core)
	//defer func(logger *zap.Logger) {
	//	err := logger.Sync()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//}(Logger)

	Logger.Info("log file was created, and the logger started correctly")
}
