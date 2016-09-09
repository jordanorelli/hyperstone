package stbl

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug = log.New(ioutil.Discard, "DEBUG stbl: ", 0)
	Info  = log.New(os.Stdout, "INFO stbl: ", 0)
)
