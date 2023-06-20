package core

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "[server]", log.Llongfile|log.Ltime)

}
