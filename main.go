package main

import (
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	replayHeader = "PBDEMS2\000"
)

func ensureNewline(t string) string {
	if strings.HasSuffix(t, "\n") {
		return t
	}
	return t + "\n"
}

func bail(status int, t string, args ...interface{}) {
	var out io.Writer
	if status == 0 {
		out = os.Stdout
	} else {
		out = os.Stderr
	}

	fmt.Fprintf(out, ensureNewline(t), args...)
	os.Exit(status)
}

func wrap(err error, t string, args ...interface{}) error {
	if err == io.EOF {
		return io.EOF
	}
	return fmt.Errorf(t+": %v", append(args, err)...)
}

type options struct {
	b bool   // bzip compression flag
	v bool   // verbose flag
	f string // input file
}

func (o options) input() (io.Reader, error) {
	var r io.Reader
	if o.f == "--" {
		r = os.Stdin
	} else {
		fi, err := os.Open(o.f)
		if err != nil {
			return nil, wrap(err, "unable to open file %s", o.f)
		}
		r = fi
	}

	if o.b {
		r = bzip2.NewReader(r)
	}
	return r, nil
}

func main() {
	var opts options
	flag.BoolVar(&opts.b, "b", false, "input is expected to be bzip-compressed")
	flag.BoolVar(&opts.v, "v", false, "verbose mode")
	flag.StringVar(&opts.f, "f", "--", "input file to be used. -- means stdin")
	flag.Parse()

	r, err := opts.input()
	if err != nil {
		bail(1, "input error: %v", err)
	}

	p := newParser(r)
	if err := p.start(); err != nil {
		bail(1, "parse error: %v", err)
	}
	if err := p.run(); err != nil && err != io.EOF {
		bail(1, "run error: %v", err)
	}
}
