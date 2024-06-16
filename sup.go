package sup

import (
	"fmt"
	"log"
)

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Cyan   Color = "\033[36m"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

var (
	cyan   = Colored("[APP]", Cyan)
	yellow = Colored("[APP]", Yellow)
	red    = Colored("[APP]", Red)
	green  = Colored("[APP]", Green)

	logLevel Level = InfoLevel
)

func SetProjectName(name string) {
	appName := fmt.Sprintf("[%s]", name)

	cyan = Colored(appName, Cyan)
	yellow = Colored(appName, Yellow)
	red = Colored(appName, Red)
	green = Colored(appName, Green)
}

func SetLogLevel(level Level) {
	logLevel = level
}

func Colored(content string, color Color) string {
	return string(color) + content + string(Reset)
}

func Debug(message interface{}) {
	if logLevel <= DebugLevel {
		log.Printf("%s: %v \n", green, message)
	}
}

func Debugf(message string, v ...interface{}) {
	Debug(fmt.Sprintf(message, v...))
}

func Info(message interface{}) {
	if logLevel <= InfoLevel {
		log.Printf("%s: %v \n", cyan, message)
	}
}

func Infof(message string, v ...interface{}) {
	Info(fmt.Sprintf(message, v...))
}

func Warning(message interface{}) {
	if logLevel <= WarningLevel {
		log.Printf("%s: %v \n", yellow, message)
	}
}

func Warningf(message string, v ...interface{}) {
	Warning(fmt.Sprintf(message, v...))
}

func Error(message interface{}) {
	if logLevel <= ErrorLevel {
		log.Printf("%s: %v \n", red, message)
	}
}

func Errorf(message string, v ...interface{}) {
	Error(fmt.Sprintf(message, v...))
}

func Fatal(message interface{}) {
	if logLevel <= FatalLevel {
		log.Fatalf("%s: %v \n", red, message)
	}
}
