package initCheck

import (
	"gecko/crossLogging"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
	"path/filepath"
)

func CheckRun(e *echo.Echo) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	logPath := os.Getenv("LOG_PATH")

	config := crossLogging.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             filepath.Join("/var/log/", logPath),
		Filename:              "test.json",
		MaxSize:               50,
		MaxBackups:            50,
		MaxAge:                60,
	}

	crossLogging.Watcher(e, config)
}
