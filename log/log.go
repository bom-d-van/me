package log

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	*log.Logger
}

func NewLogger(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{log.New(out, prefix, flag)}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Fatal(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Fatalf(format, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Fatalln(v...)
}

func (l *Logger) Flags() int {
	return l.Logger.Flags()
}

func (l *Logger) Output(calldepth int, s string) error {
	l.PrintFileLine()
	return l.Logger.Output(calldepth, s)
}

func (l *Logger) Panic(v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Panic(v...)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Panicf(format, v...)
}

func (l *Logger) Panicln(v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Panicln(v...)
}

func (l *Logger) Prefix() string {
	return l.Logger.Prefix()
}

func (l *Logger) Print(v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Print(v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Printf(format, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.PrintFileLine()
	l.Logger.Println(v...)
}

var mePath = os.Getenv("GOPATH") + "/src/github.com/bom-d-van/me/"

func (l *Logger) PrintFileLine() {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		l.Logger.Println("Can't Retrieve Caller")
		return
	}

	file = strings.Replace(file, mePath, "", 1)
	l.Logger.Printf("---> %s:%d", file, line)
}
