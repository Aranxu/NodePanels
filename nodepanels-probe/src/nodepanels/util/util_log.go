package util

import (
	"fmt"
	"log"
	"os"
)

var LogFile *os.File

func LogInfo(msg string) {
	fmt.Println(msg)
	debugLog := log.New(LogFile, "[Info]", log.LstdFlags)
	debugLog.Println(msg)
}

func LogDebug(msg string) {
	fmt.Println(msg)
	debugLog := log.New(LogFile, "[Debug]", log.LstdFlags)
	debugLog.Println(msg)
}

func LogError(msg string) {
	fmt.Println(msg)
	debugLog := log.New(LogFile, "[Error]", log.LstdFlags)
	debugLog.Println(msg)
}

func LogWarn(msg string) {
	fmt.Println(msg)
	debugLog := log.New(LogFile, "[Warn]", log.LstdFlags)
	debugLog.Println(msg)
}
