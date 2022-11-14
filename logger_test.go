package jagger

import (
	"log"
	"testing"
)

type LogStruct struct {
	Title string `json:"Title"`
	Msg   string `json:"Msg"`
}

func TestDefaultLogging(t *testing.T) {
	info := "info log"
	warning := "warning log"
	errL := "error log"
	fatal := "fatal log"
	lMsg := LogStruct{Title: "Struct Log", Msg: "This is Struct Log Message"}

	Info(info)
	Warning(warning)
	Error(errL)
	Errorj(lMsg)
	Fatal(fatal)
}

func TestCustomLogging(t *testing.T) {
	l := New("Test")
	l.SetLevel(ParseLevel("Warning"))
	l.SetFlags(log.Ldate)

	info := "info log"
	warning := "warning log"
	errL := "error log"
	fatal := "fatal log"
	lMsg := LogStruct{Title: "Struct Log", Msg: "This is Struct Log Message"}

	l.Info(info)
	l.Warning(warning)
	l.Error(errL)
	l.Errorj(lMsg)
	l.Fatal(fatal)
}
