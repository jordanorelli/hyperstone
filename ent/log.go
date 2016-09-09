package ent

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug = log.New(ioutil.Discard, "DEBUG ent: ", 0)
	Info  = log.New(os.Stdout, "INFO end: ", 0)
)
