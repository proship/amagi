package logger

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type (
	// LogToStream log messages to stream
	LogToStream struct {
		MaxProgress     int
		CurrentProgress int

		LogStream         io.Writer
		ProgressStream    io.Writer
		MaxProgressStream io.Writer
	}
)

// Initialize initialize the logger with the ID
func (log *LogToStream) Initialize(id string) {}

// Info send [INFO] message to log
func (log *LogToStream) Info(message string) {
	logMessageToStream(log.LogStream, "Info", message)
}

// Warn send [WARN] message to log
func (log *LogToStream) Warn(message string) {
	logMessageToStream(log.LogStream, "Warn", message)
}

// Error send [ERROR] message to log
func (log *LogToStream) Error(message string) {
	logMessageToStream(log.LogStream, "Error", message)
}

// Fatal send [FATAL] message to log
func (log *LogToStream) Fatal(message string) {
	logMessageToStream(log.LogStream, "Fatal", message)
}

// SetProgressMax sets the maximum Progress in int
func (log *LogToStream) SetProgressMax(max int) {
	log.MaxProgressStream.Write([]byte(strconv.Itoa(max)))
	log.MaxProgress = max
}

// ProgressInc incease current progress with int as param
func (log *LogToStream) ProgressInc(progress int) {
	log.ProgressStream.Write([]byte(strconv.Itoa(progress)))
	log.CurrentProgress = log.CurrentProgress + progress
}

// Finalize finalize the execution and max out progress
func (log *LogToStream) Finalize() {
	log.ProgressStream.Write(nil)
}

func logMessageToStream(stream io.Writer, logType, message string) {
	stream.Write([]byte(fmt.Sprintf("[%s] %s", strings.ToUpper(logType), message)))
}
