package jagger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

const (
	flags       = log.Ldate | log.Lmicroseconds | log.Lshortfile
	defaultName = "Jagger"
)

const (
	tagDebug   = "DEBUG : "
	tagInfo    = "INFO : "
	tagWarning = "WARN : "
	tagError   = "ERROR : "
	tagFatal   = "FATAL : "
)

type Level int

type Logger struct {
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
	level      Level
}

var defaultLogger = New(defaultName)

func New(name string) *Logger {
	return &Logger{
		debugLog:   log.New(os.Stderr, "["+name+"] "+tagDebug, flags),
		infoLog:    log.New(os.Stderr, "["+name+"] "+tagInfo, flags),
		warningLog: log.New(os.Stderr, "["+name+"] "+tagWarning, flags),
		errorLog:   log.New(os.Stderr, "["+name+"] "+tagError, flags),
		fatalLog:   log.New(os.Stderr, "["+name+"] "+tagFatal, flags),
		level:      InfoLevel,
	}
}

func ParseLevel(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarningLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	default:
		return InfoLevel
	}
}

func SetFlags(flag int) {
	defaultLogger.SetFlags(flag)
}

func SetLevel(lvl Level) {
	defaultLogger.SetLevel(lvl)
}

func SetOutput(w io.Writer) {
	defaultLogger.debugLog.SetOutput(w)
}

func (l *Logger) SetLevel(lvl Level) {
	l.level = lvl
}

func (l *Logger) SetFlags(flag int) {
	l.debugLog.SetFlags(flag)
	l.infoLog.SetFlags(flag)
	l.warningLog.SetFlags(flag)
	l.errorLog.SetFlags(flag)
	l.fatalLog.SetFlags(flag)
}

func (l *Logger) SetOutput(w io.Writer) {
	l.debugLog.SetOutput(w)
	l.infoLog.SetOutput(w)
	l.warningLog.SetOutput(w)
	l.errorLog.SetOutput(w)
	l.fatalLog.SetOutput(w)
}

func (l *Logger) isLevelEnabled(level Level) bool {
	return level >= l.level
}

func (l *Logger) output(s Level, txt string) {
	if l.isLevelEnabled(s) {
		switch s {
		case DebugLevel:
			l.debugLog.Output(2, txt)
		case InfoLevel:
			l.infoLog.Output(2, txt)
		case WarningLevel:
			l.warningLog.Output(2, txt)
		case ErrorLevel:
			l.errorLog.Output(2, txt)
		case FatalLevel:
			l.fatalLog.Output(2, txt)
		default:
			panic(fmt.Sprintln("unrecognized severity:", s))
		}
	}
}

func getJsonMessage(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	message := string(b)
	return message
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(DebugLevel, fmt.Sprint(v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.output(DebugLevel, fmt.Sprintln(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.output(DebugLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Debugj(v interface{}) {
	l.output(DebugLevel, getJsonMessage(v))
}

func (l *Logger) Info(v ...interface{}) {
	l.output(InfoLevel, fmt.Sprint(v...))
}

func (l *Logger) Infoln(v ...interface{}) {
	l.output(InfoLevel, fmt.Sprintln(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.output(InfoLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Infoj(v interface{}) {
	l.output(InfoLevel, getJsonMessage(v))
}

func (l *Logger) Warning(v ...interface{}) {
	l.output(WarningLevel, fmt.Sprint(v...))
}

func (l *Logger) Warningln(v ...interface{}) {
	l.output(WarningLevel, fmt.Sprintln(v...))
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.output(WarningLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Warningj(v interface{}) {
	l.output(WarningLevel, getJsonMessage(v))
}

func (l *Logger) Error(v ...interface{}) {
	l.output(ErrorLevel, fmt.Sprint(v...))
}

func (l *Logger) Errorln(v ...interface{}) {
	l.output(ErrorLevel, fmt.Sprintln(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.output(ErrorLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorj(v interface{}) {
	l.output(ErrorLevel, getJsonMessage(v))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.output(FatalLevel, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.output(FatalLevel, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.output(FatalLevel, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Fatalj(v interface{}) {
	l.output(FatalLevel, getJsonMessage(v))
}

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Debugln(v ...interface{}) {
	defaultLogger.Debugln(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Debugj(v interface{}) {
	defaultLogger.Debugj(v)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infoln(v ...interface{}) {
	defaultLogger.Infoln(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Infoj(v interface{}) {
	defaultLogger.Infoj(v)
}

func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

func Warningln(v ...interface{}) {
	defaultLogger.Warningln(v...)
}

func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

func Warningj(v interface{}) {
	defaultLogger.Warningj(v)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Errorln(v ...interface{}) {
	defaultLogger.Errorln(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Errorj(v interface{}) {
	defaultLogger.Errorj(v)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Fatalln(v ...interface{}) {
	defaultLogger.Fatalln(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func Fatalj(v interface{}) {
	defaultLogger.Fatalj(v)
}
