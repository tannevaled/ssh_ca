package utils

import (
	"log"
	"runtime"

	"github.com/fatih/color"
)

func Trace() {
	pc, filename, line, _ := runtime.Caller(1)
	color.Set(color.FgGreen)
	log.Printf("[trace] in %s[%s:%d]", runtime.FuncForPC(pc).Name(), filename, line)
	color.Unset()
}
