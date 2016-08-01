package main

import (
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
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
	b          bool   // bzip compression flag
	v          bool   // verbose flag
	f          string // input file
	memprofile string
	cpuprofile string
	messages   bool // dump messages or no
	packets    bool // dump packets or no
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

func memprofile(dest string) {
	fmt.Println("writing mem profile to", dest)
	w, err := os.Create(dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open memprofile output %s: %v\n", dest, err)
	} else {
		defer w.Close()
		err := pprof.WriteHeapProfile(w)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to write heap profile: %v\n", err)
		}
	}
}

func cpuprofile(dest string) func() {
	fmt.Println("writing cpu profile to", dest)
	w, err := os.Create(dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open cpuprofile output %s: %v\n", dest, err)
		return func() {}
	} else {
		pprof.StartCPUProfile(w)
		return func() {
			pprof.StopCPUProfile()
			w.Close()
		}
	}
}

func main() {
	var opts options
	flag.BoolVar(&opts.b, "b", false, "input is expected to be bzip-compressed")
	flag.BoolVar(&opts.v, "v", false, "verbose mode")
	flag.StringVar(&opts.f, "f", "--", "input file to be used. -- means stdin")
	flag.BoolVar(&opts.messages, "messages", false, "dump top-level messages to stdout")
	flag.BoolVar(&opts.packets, "packets", false, "dump packets to stdout")
	flag.StringVar(&opts.memprofile, "memprofile", "", "memory profile destination")
	flag.StringVar(&opts.cpuprofile, "cpuprofile", "", "cpu profile destination")
	flag.Parse()

	if opts.memprofile != "" {
		defer memprofile(opts.memprofile)
	}

	if opts.cpuprofile != "" {
		defer cpuprofile(opts.cpuprofile)()
	}

	r, err := opts.input()
	if err != nil {
		bail(1, "input error: %v", err)
	}

	p := newParser(r)
	p.dumpMessages = opts.messages
	p.dumpPackets = opts.packets
	if err := p.start(); err != nil {
		bail(1, "parse error: %v", err)
	}
	if err := p.run(); err != nil && err != io.EOF {
		bail(1, "run error: %v", err)
	}
}
