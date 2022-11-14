package jagger

import (
	"log"
	"testing"
)

func TestDefaultLogging(t *testing.T) {
	info := "info log"
	warning := "warning log"
	errL := "error log"
	fatal := "fatal log"

	Info(info)
	Warning(warning)
	Error(errL)
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

	l.Info(info)
	l.Warning(warning)
	l.Error(errL)
	l.Fatal(fatal)
}
