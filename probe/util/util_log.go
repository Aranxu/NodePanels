package util

import (
	"log"
	"os"
)

var LogFile *os.File

func LogInfo(msg string) {
	debugLog := log.New(LogFile, "[Info]", log.LstdFlags)
	debugLog.Println(msg)
}

func LogDebug(msg string) {
	debugLog := log.New(LogFile, "[Debug]", log.LstdFlags)
	debugLog.Println(msg)
}

func LogError(msg string) {
	debugLog := log.New(LogFile, "[Error]", log.LstdFlags)
	debugLog.Println(msg)
}

func LogWarn(msg string) {
	debugLog := log.New(LogFile, "[Warn]", log.LstdFlags)
	debugLog.Println(msg)
}
