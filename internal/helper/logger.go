package helper

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	InfoLog = log.New(os.Stdout, "[INFO]: ", 0)
	ErrorLog = log.New(os.Stderr, "[ERROR]: ", log.Lshortfile)
}
