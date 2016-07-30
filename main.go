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
	replay_header = "PBDEMS2\000"
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
			return nil, fmt.Errorf("unable to open file %s: %v", o.f, err)
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

	buf := make([]byte, 8)
	if _, err := r.Read(buf); err != nil {
		bail(1, "error reading header: %v", err)
	}
	if string(buf) != replay_header {
		bail(1, "unexpected replay header: %s", string(buf))
	}
	fmt.Println(string(buf))
}
