package ent

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug = log.New(ioutil.Discard, "DEBUG ent: ", 0)
	Info  = log.New(os.Stdout, "INFO end: ", 0)
)

func wrap(err error, t string, args ...interface{}) error {
	return fmt.Errorf(t+": %v", append(args, err)...)
}
