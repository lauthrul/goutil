package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// With the exception of the Lmsgprefix flag, there is no
// control over the order they appear (the order listed here)
// or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//	2009/01/23 01:23:23 message
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
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
		//mv = io.MultiWriter(os.Stdout, file)
		mv = file
	}
	logger = log.New(mv, "", LstdFlags|Lmicroseconds|Lshortfile)
}

func SetFlag(flag int) {
	logger.SetFlags(flag)
}

func AppendFlag(flag int) {
	f := logger.Flags()
	logger.SetFlags(f | flag)
}

func RemoveFlag(flag int) {
	f := logger.Flags()
	logger.SetFlags(f & ^flag)
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
