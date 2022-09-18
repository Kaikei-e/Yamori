package crossLogging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func LoggerSetup(path string) {
	//ws, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}

	rotateLogger := &lumberjack.Logger{
		Filename:   "server.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default

	}

	writerSyncer := zapcore.AddSync(rotateLogger)

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

	Logger = zap.New(core, zap.AddCaller())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}

	}(Logger)

	for i := 0; i < 100; i++ {
		Logger.Info("test")
	}
}
