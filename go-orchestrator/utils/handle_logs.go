package utils

// This file sets up the logging configuration using zerolog and lumberjack for log rotation.
// It initializes a global logger that can be used throughout the application to
// - log structured messages to both the console and a rotating log file.
// The InitLogger function configures the logger to write to a file with
// - rotation settings and also outputs to the console in a human-readable format.
// The global log level is set to Info, meaning that only messages at this level or higher will be logged.

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"orchestrator/utils"
)

var Log zerolog.Logger

func InitLogger() {
	utils.LoadConfigs()
	Filename := utils.AppConfig.Logging.LogDir + "/" + utils.AppConfig.Logging.LogFile


	// 1. setup lumberjack for log rotation
	fileLogger := &lumberjack.Logger{
		Filename: Filename,
		MaxSize: 10,
		MaxBackups: 3,
		MaxAge: 30,
		Compress: true,
	}

	// 2. setup console output
	consoleWriter := zerolog.ConsoleWriter{
		Out: os.Stdout,
		TimeFormat: time.DateTime,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileLogger)

	Log = zerolog.New(multi).With().Timestamp().Caller().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

}
