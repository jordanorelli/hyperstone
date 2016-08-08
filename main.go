package main

// go:generate ./mkprotos
//go:generate go run ./gen/main.go ./dota

import (
	"bytes"
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime/pprof"
	"strings"

	"github.com/golang/protobuf/proto"
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

func check(msg proto.Message) {}

func printTypes(m proto.Message) {
	fmt.Println(reflect.TypeOf(m))
}

func prettySlice(v reflect.Value) string {
    if v.Type().Elem().Kind() == reflect.Uint8 {
        l := v.Len()
        if l > 16 {
            l = 16
        }
        b := make([]byte, l)
        for i := 0; i < l; i++ {
            b[i] = byte(v.Index(i).Uint())
        }
        return fmt.Sprintf("%x", b)
    }

    width := 0
    parts := make([]string, 0, v.Len())
    for i := 0; i < v.Len() && width <= 32; i++ {
        parts = append(parts, pretty(v.Index(i)))
        width += len(parts[i]) // obligatory byte count is not rune count rabble
    }
    return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}

func prettyStruct(v reflect.Value) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "{%s ", v.Type())
	for fn := 0; fn < v.NumField(); fn++ {
		field := v.Type().Field(fn)
		if field.Name == "XXX_unrecognized" {
			continue
		}
        fv := v.Field(fn)
        fmt.Fprintf(&buf, "%s: %s ", field.Name, pretty(fv))
	}
    fmt.Fprint(&buf, "}")
    return buf.String()
}

func pretty(v reflect.Value) string {
	switch v.Kind() {
    case reflect.Ptr:
        if v.IsNil() {
            return "nil"
        }
        return pretty(v.Elem())
	case reflect.Struct:
		return prettyStruct(v)
    case reflect.Slice:
        return prettySlice(v)
	case reflect.String:
        return fmt.Sprintf("%q", v.String())
	case reflect.Int32:
        return fmt.Sprintf("%d", v.Int())
	case reflect.Uint8, reflect.Uint32:
        return fmt.Sprintf("%d", v.Uint())
	case reflect.Bool:
        return fmt.Sprintf("%t", v.Bool())
    default:
        return v.Type().Name()
	}
}

func prettyVals(m proto.Message) {
	v := reflect.ValueOf(m)
	fmt.Println(pretty(v))
}

func main() {
	var opts options
	flag.BoolVar(&opts.b, "b", false, "input is expected to be bzip-compressed")
	flag.BoolVar(&opts.v, "v", false, "verbose mode")
	flag.StringVar(&opts.f, "f", "--", "input file to be used. -- means stdin")
	flag.StringVar(&opts.memprofile, "memprofile", "", "memory profile destination")
	flag.StringVar(&opts.cpuprofile, "cpuprofile", "", "cpu profile destination")
	flag.Parse()

	var handle func(proto.Message)
	switch flag.Arg(0) {
	case "":
		handle = check
	case "types":
		handle = printTypes
	case "pretty":
		handle = prettyVals
	default:
		bail(1, "no such action: %s", flag.Arg(0))
	}

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

	c := make(chan maybe, 32)
	p := newParser(r)
	delete(p.ewl, EBaseGameEvents_GE_SosStartSoundEvent)
	delete(p.ewl, EDotaUserMessages_DOTA_UM_SpectatorPlayerUnitOrders)
	delete(p.ewl, EDotaUserMessages_DOTA_UM_SpectatorPlayerClick)
	delete(p.ewl, EDotaUserMessages_DOTA_UM_TE_UnitAnimation)
	delete(p.ewl, EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd)
	go p.run(c)
	for m := range c {
		if m.error != nil {
			fmt.Fprintln(os.Stderr, m.error)
		} else {
			handle(m.Message)
			messages.Return(m.Message)
		}
	}
	if p.err != nil {
		fmt.Printf("parser error: %v\n", p.err)
	}
}
