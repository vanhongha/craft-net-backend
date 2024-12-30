package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance *Logger
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	file        *os.File
}

func InitLogger(logFilePath string) *Logger {
	once.Do(func() {
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		multiWriter := io.MultiWriter(file, os.Stdout)
		instance = &Logger{
			infoLogger:  log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime),
			errorLogger: log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime),
			file:        file,
		}

		log.New(multiWriter, "-------------------------", log.Lmsgprefix).Println("")
		log.Println("Logger initialized successfully.")
	})

	return instance
}

// singleton instance
func GetLogger() *Logger {
	if instance == nil {
		InitLogger("../../logs/app.log")
	}
	if instance == nil {
		log.Fatalf("Logger has not been initialized. Call InitLogger first.")
	}
	return instance
}

func (l *Logger) LogInfo(message string) {
	l.infoLogger.Println(l.formatLog(message))
}

func (l *Logger) LogErrorWithMsg(message string, isExit bool) {
	l.errorLogger.Println(l.formatLog(message))
	if isExit {
		log.Fatalf("%s", "ERROR: "+message)
	}
}

func (l *Logger) LogErrorWithMsgAndError(message string, err error, isExit bool) {
	errMsg := fmt.Sprintf(message+" - %v", err)
	l.errorLogger.Println(l.formatLog(errMsg))
	if isExit {
		log.Fatalf("%s", "ERROR: "+errMsg)
	}
}

func (l *Logger) formatLog(message string) string {
	projectFolderName := "craft-net-backend"
	// logger called file, line, function's name
	_, file, line, ok := runtime.Caller(2) // logger called's stack trace
	if !ok {
		file = "unknown"
		line = 0
	}
	shortFile := file
	if idx := strings.Index(file, projectFolderName); idx != -1 {
		shortFile = file[idx+len(projectFolderName)+1:]
	}
	return fmt.Sprintf("[%s:%d] %s", shortFile, line, message)
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
