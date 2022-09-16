// Package initCheck checks the required files and minimum requirements for start application
package initCheck

import (
	"fmt"
	"gecko/crossLogging"
	"github.com/labstack/echo/v4"
	"os"
	"path/filepath"
)

// CheckRun checks if the .env file exists
// and configure the logger
func CheckRun(e *echo.Echo) {

	logPath := os.Getenv("LOG_PATH")

	fmt.Println(logPath)

	if _, err := os.Stat(filepath.Join("/var/log/", logPath)); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Join("/var/log/", logPath), 0750)
		if err != nil {
			e.Logger.Fatal(err)
		}
	}

	config := crossLogging.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             filepath.Join("/var/log/", logPath),
		Filename:              "server.log",
		MaxSize:               50,
		MaxBackups:            50,
		MaxAge:                60,
	}

	err := crossLogging.Watcher(e, config)
	if err != nil {
		e.Logger.Errorf(err.Error())
		e.Logger.Fatal("Failed to configure logger")

	}

	crossLogging.Logger.Info().Msg("log file was created, and the logger started correctly")

}
