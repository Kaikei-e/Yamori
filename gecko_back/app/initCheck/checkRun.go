package initCheck

import (
	"gecko/crossLogging"
	"github.com/labstack/echo/v4"
	"os"
	"path/filepath"
)

func CheckRun(e *echo.Echo) {

	err := EnvLoader()
	if err != nil {
		e.Logger.Fatal(err)
	}

	logPath := os.Getenv("LOG_PATH")

	if _, err := os.Stat(filepath.Join("/var/log/", logPath)); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Join("/var/log/", logPath), 0755)
		if err != nil {
			crossLogging.Logger.Fatal().Err(err).Msg("failed to create log directory")
		}
	}

	config := crossLogging.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             filepath.Join("/var/log/", logPath, "server.log"),
		Filename:              "server.json",
		MaxSize:               50,
		MaxBackups:            50,
		MaxAge:                60,
	}

	err = crossLogging.Watcher(e, config)
	if err != nil {
		e.Logger.Errorf(err.Error())
		e.Logger.Fatal("Failed to configure logger")

	}

}
