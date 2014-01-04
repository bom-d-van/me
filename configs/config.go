package configs

import (
	"log"
	"os"
)

var (
	Port           = ":80"
	Log            = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	ReLoadTemplate = false
)
