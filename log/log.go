package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level byte

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var prefixes = map[Level]string{
	LevelDebug: "DEBUG\t",
	LevelInfo:  "INFO\t",
	LevelWarn:  "WARN\t",
	LevelError: "ERROR\t",
	LevelFatal: "FATAL\t",
}

var (
	logger *log.Logger
	lvl    Level = LevelInfo
)

func Init(filePath string) {
	mv := io.Writer(os.Stdout)
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println(err)
		}
		mv = io.MultiWriter(os.Stdout, file)
	}
	logger = log.New(mv, "", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
}

func SetLevel(level Level) {
	lvl = level
}

func GetLevel() (Level, string) {
	return lvl, prefixes[lvl]
}

func write(level Level, v ...interface{}) {
	if level < lvl || logger == nil {
		return
	}
	logger.SetPrefix(prefixes[level])
	logger.Output(3, fmt.Sprintln(v...))
}

func writeF(level Level, format string, v ...interface{}) {
	if level < lvl || logger == nil {
		return
	}
	logger.SetPrefix(prefixes[level])
	logger.Output(3, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	write(LevelDebug, v...)
}

func Info(v ...interface{}) {
	write(LevelInfo, v...)
}

func Warn(v ...interface{}) {
	write(LevelWarn, v...)
}

func Error(v ...interface{}) {
	write(LevelError, v...)
}

func Fatal(v ...interface{}) {
	write(LevelFatal, v...)
}

func DebugF(fmt string, v ...interface{}) {
	writeF(LevelDebug, fmt, v...)
}

func InfoF(fmt string, v ...interface{}) {
	writeF(LevelInfo, fmt, v...)
}

func WarnF(fmt string, v ...interface{}) {
	writeF(LevelWarn, fmt, v...)
}

func ErrorF(fmt string, v ...interface{}) {
	writeF(LevelError, fmt, v...)
}

func FatalF(fmt string, v ...interface{}) {
	writeF(LevelFatal, fmt, v...)
}
