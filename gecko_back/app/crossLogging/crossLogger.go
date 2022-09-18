package crossLogging

//
//import (
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/echo/v4/middleware"
//	"github.com/rs/zerolog"
//	"gopkg.in/natefinch/lumberjack.v2"
//	"io"
//	"os"
//	"path"
//)
//
//var Logger *zerolog.Logger
//
//type Config struct {
//	ConsoleLoggingEnabled bool
//	EncodeLogsAsJson      bool
//	FileLoggingEnabled    bool
//	Directory             string
//	Filename              string
//	MaxSize               int
//	MaxBackups            int
//	MaxAge                int
//}
//
//func Watcher(e *echo.Echo, config Config) error {
//	Logger = configureLogger(config)
//	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
//		LogURI:    true,
//		LogStatus: true,
//		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
//			Logger.Info().
//				Str("URI", v.URI).
//				Int("status", v.Status).
//				Timestamp().
//				Dur("latency", v.Latency).
//				Msg("request")
//
//			return nil
//		},
//	}))
//
//	return nil
//}
//
//func configureLogger(config Config) *zerolog.Logger {
//	var writers []io.Writer
//
//	if config.ConsoleLoggingEnabled {
//		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
//	}
//	if config.FileLoggingEnabled {
//		writers = append(writers, logRotation(config))
//	}
//
//	multiWriter := io.MultiWriter(writers...)
//
//	logger := zerolog.New(multiWriter).With().Timestamp().Logger()
//	logger.Info().
//		Bool("fileLogging", config.FileLoggingEnabled).
//		Bool("jsonLogOutput", config.EncodeLogsAsJson).
//		Str("logDirectory", config.Directory).
//		Int("macSizeMB", config.MaxSize).
//		Int("maxBackups", config.MaxBackups).
//		Int("maxAgeInDays", config.MaxAge).Msg("logger was configured")
//
//	return &logger
//
//}
//
//func logRotation(config Config) io.Writer {
//	return &lumberjack.Logger{
//		Filename:   path.Join(config.Directory, config.Filename),
//		MaxBackups: config.MaxBackups,
//		MaxSize:    config.MaxSize,
//		MaxAge:     config.MaxAge,
//	}
//}
