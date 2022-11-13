package jagger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

const (
	flags    = log.Ldate | log.Lmicroseconds | log.Lshortfile
	initText = "Jagger"
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
	sync.Mutex
}

var defaultLogger = New(initText)

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
	l.Lock()
	defer l.Unlock()
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

func (l *Logger) Debug(v ...interface{}) {
	l.output(InfoLevel, fmt.Sprint(v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.output(InfoLevel, fmt.Sprintln(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.output(InfoLevel, fmt.Sprintf(format, v...))
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

func (l *Logger) Warning(v ...interface{}) {
	l.output(WarningLevel, fmt.Sprint(v...))
}

func (l *Logger) Warningln(v ...interface{}) {
	l.output(WarningLevel, fmt.Sprintln(v...))
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.output(WarningLevel, fmt.Sprintf(format, v...))
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

func Debug(v ...interface{}) {
	defaultLogger.output(InfoLevel, fmt.Sprint(v...))
}

func Debugln(v ...interface{}) {
	defaultLogger.output(InfoLevel, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.output(InfoLevel, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	defaultLogger.output(InfoLevel, fmt.Sprint(v...))
}

func Infoln(v ...interface{}) {
	defaultLogger.output(InfoLevel, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	defaultLogger.output(InfoLevel, fmt.Sprintf(format, v...))
}

func Warning(v ...interface{}) {
	defaultLogger.output(WarningLevel, fmt.Sprint(v...))
}

func Warningln(v ...interface{}) {
	defaultLogger.output(WarningLevel, fmt.Sprintln(v...))
}

func Warningf(format string, v ...interface{}) {
	defaultLogger.output(WarningLevel, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	defaultLogger.output(ErrorLevel, fmt.Sprint(v...))
}

func Errorln(v ...interface{}) {
	defaultLogger.output(ErrorLevel, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.output(ErrorLevel, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	defaultLogger.output(FatalLevel, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	defaultLogger.output(FatalLevel, fmt.Sprintln(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.output(FatalLevel, fmt.Sprintf(format, v...))
	os.Exit(1)
}
