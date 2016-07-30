package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ensureNewline(t string) string {
	if strings.HasSuffix(t, "\n") {
		return t
	}
	return t + "\n"
}

func usage(t string, args ...interface{}) {
	t = ensureNewline(t) + "\n"
	fmt.Fprintf(os.Stderr, t, args...)
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "\thyperstone replay-id")
	os.Exit(1)
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

func openReplayFile(path string) (io.ReadCloser, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open replay file: %v", err)
	}

	switch filepath.Ext(path) {
	case ".bz2":
		return newBzipCloser(fi), nil
	default:
		return fi, nil
	}
	return nil, io.EOF
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage("please supply the path to a valid replay file")
	}
	f, err := openReplayFile(flag.Arg(0))
	if err != nil {
		bail(1, "unable to open replay file: %v", err)
	}
	defer f.Close()
}
