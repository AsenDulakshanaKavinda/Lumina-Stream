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
)

var Log zerolog.Logger

func InitLogger() {
	
	logFile := GetConfig().Logging.LogFile
	logDir := GetConfig().Logging.LogDir
	fileName := logDir + "/" + logFile


	// 1. setup lumberjack for log rotation
	fileLogger := &lumberjack.Logger{
		Filename: fileName,
		MaxSize: GetConfig().Logging.MaxSize,
		MaxBackups: GetConfig().Logging.MaxBackups,
		MaxAge: GetConfig().Logging.MaxAge,
		Compress: GetConfig().Logging.Compress,
	}

	// 2. setup console output
	consoleWriter := zerolog.ConsoleWriter{
		Out: os.Stdout,
		TimeFormat: time.DateTime,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileLogger)

	// - setup the logging format
	Log = zerolog.New(multi).With().Timestamp().Caller().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

}
