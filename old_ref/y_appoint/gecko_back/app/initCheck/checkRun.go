// Package initCheck checks the required files and minimum requirements for start application
package initCheck

import (
	"fmt"
	"gecko/crossLogging"
	"os"
	"path/filepath"
)

// CheckRun checks if the .env file exists
// and configure the logger
func CheckRun() {

	logPath := os.Getenv("LOG_PATH")

	fullPath := filepath.Join("/var/log/", logPath)

	fmt.Println(fullPath)

	//check the existence of the log directory
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err := os.Mkdir(fullPath, 0750)
		if err != nil {
			panic(err)
		}
	}

	crossLogging.LoggerSetup(fullPath)

	//
	//config := crossLogging.Config{
	//	ConsoleLoggingEnabled: true,
	//	EncodeLogsAsJson:      true,
	//	FileLoggingEnabled:    true,
	//	Directory:             filepath.Join("/var/log/", logPath),
	//	Filename:              "server.log",
	//	MaxSize:               50,
	//	MaxBackups:            50,
	//	MaxAge:                60,
	//}
	//
	//err := crossLogging.Watcher(e, config)
	//if err != nil {
	//	e.Logger.Errorf(err.Error())
	//	e.Logger.Fatal("Failed to configure logger")
	//
	//}
	//
	//crossLogging.Logger.Info().Msg("log file was created, and the logger started correctly")

}
