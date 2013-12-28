package configs

import (
	"github.com/bom-d-van/me/log"
	"os"
)

var (
	Port           = ":80"
	Log            = log.NewLogger(os.Stdout, "", 0)
	ReLoadTemplate = false
)
