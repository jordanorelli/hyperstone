package main

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug = log.New(ioutil.Discard, "DEBUG main: ", 0)
	Info  = log.New(os.Stdout, "INFO main: ", 0)
)
