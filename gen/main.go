package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var (
	messageTypes    = make(map[string]bool)
	enumTypes       = make(map[string]bool)
	entityTypes     = make(typeMap)
	cmdTypes        = make(typeMap)
	cmdEnumType     = "EDemoCommands"
	entityEnumTypes = map[string]bool{
		"NET_Messages":        true,
		"SVC_Messages":        true,
		"EBaseUserMessages":   true,
		"EBaseEntityMessages": true,
		"EBaseGameEvents":     true,
		"EDotaUserMessages":   true,
		"ETEProtobufIds":      true,
	}
	prefixes = map[string]string{
		"EDemoCommands_DEM_":         "CDemo",
		"NET_Messages_net_":          "CNETMsg_",
		"SVC_Messages_svc_":          "CSVCMsg_",
		"EBaseUserMessages_UM_":      "CUserMessage",
		"EBaseEntityMessages_EM_":    "CEntityMessage",
		"EBaseGameEvents_GE_":        "CMsg",
		"EDotaUserMessages_DOTA_UM_": "CDOTAUserMsg_",
	}
	specials = map[string]string{
		"EDotaUserMessages_DOTA_UM_StatsHeroDetails":  "CDOTAUserMsg_StatsHeroMinuteDetails",
		"EDotaUserMessages_DOTA_UM_CombatLogDataHLTV": "CMsgDOTACombatLogEntry",
		"EDotaUserMessages_DOTA_UM_TournamentDrop":    "CMsgGCToClientTournamentItemDrop",
		"EDotaUserMessages_DOTA_UM_MatchMetadata":     "CDOTAClientMsg_MatchMetadata",
		"ETEProtobufIds_TE_EffectDispatchId":          "CMsgTEEffectDispatch",
		"EDemoCommands_DEM_SignonPacket":              "CDemoPacket",
	}
	// EBaseUserMessages_UM_HandHapticPulse
	tpl = `package main

////////////////////////////////////////////////////////////////////////////////
//
//                           .aMMMb  .aMMMb  dMMMMb  dMMMMMP
//                          dMP"VMP dMP"dMP dMP VMP dMP
//                         dMP     dMP dMP dMP dMP dMMMP
//                        dMP.aMP dMP.aMP dMP.aMP dMP
//                        VMMMP"  VMMMP" dMMMMP" dMMMMMP
//
//       .aMMMMP dMMMMMP dMMMMb  dMMMMMP dMMMMb  .aMMMb dMMMMMMP dMMMMMP dMMMMb
//      dMP"    dMP     dMP dMP dMP     dMP.dMP dMP"dMP   dMP   dMP     dMP VMP
//     dMP MMP"dMMMP   dMP dMP dMMMP   dMMMMK" dMMMMMP   dMP   dMMMP   dMP dMP
//    dMP.dMP dMP     dMP dMP dMP     dMP"AMF dMP dMP   dMP   dMP     dMP.aMP
//    VMMMP" dMMMMMP dMP dMP dMMMMMP dMP dMP dMP dMP   dMP   dMMMMMP dMMMMP"
//
//
//  This code was generated by a code-generation program. It was NOT written by
//  hand. Do not edit this file by hand! Your edits will be destroyed!
//
//  This file can be regenerated by running "go generate"
//
//  The generator program is defined in "gen/main.go"
//
////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/golang/protobuf/proto"
	"github.com/jordanorelli/hyperstone/dota"
)

type datagramType int32
type entityType int32

const (
	EDemoCommands_DEM_Error datagramType = -1
{{- range $id, $spec := .Commands }}
	{{$spec.EnumName}} datagramType = {{$id}}
{{- end }}
{{- range $id, $spec := .Entities }}
	{{$spec.EnumName}} entityType = {{$id}}
{{- end }}
)

func (d datagramType) String() string {
	switch d {
{{- range $id, $spec := .Commands }}
	case {{$spec.EnumName}}:
		return "{{$spec.EnumName}}"
{{- end }}
	default:
		return "UnknownDatagramType"
	}
}

func (e entityType) String() string {
	switch e {
{{- range $id, $spec := .Entities }}
	case {{$spec.EnumName}}:
		return "{{$spec.EnumName}}"
{{- end }}
	default:
		return "UnknownEntityType"
	}
}

type messageStatus int

const (
	m_Unknown messageStatus = iota
	m_Skipped
)

func (m messageStatus) Error() string {
	switch m {
	case m_Unknown:
		return "unknown message type"
	case m_Skipped:
		return "skipped message type"
	default:
		return "unknown message error"
	}
}

type datagramFactory map[datagramType]func() proto.Message
type entityFactory map[entityType]func() proto.Message

type messageFactory struct {
	datagrams datagramFactory
	entities entityFactory
}

func (m *messageFactory) BuildDatagram(id datagramType) (proto.Message, error) {
	fn, ok := m.datagrams[id]
	if !ok {
		return nil, m_Unknown
	}
	return fn(), nil
}

func (m *messageFactory) BuildEntity(id entityType) (proto.Message, error) {
	fn, ok := m.entities[id]
	if !ok {
		return nil, m_Unknown
	}
	return fn(), nil
}

var messages = messageFactory{
	datagramFactory{
	{{- range $id, $spec := .Commands }}
		{{$spec.EnumName}}: func() proto.Message { return new(dota.{{$spec.TypeName}}) },
	{{- end }}
	},
	entityFactory{
	{{- range $id, $spec := .Entities }}
		{{$spec.EnumName}}: func() proto.Message { return new(dota.{{$spec.TypeName}}) },
	{{- end }}
	},
}
`
)

type messageSpec struct {
	EnumName string
	TypeName string
}

type typeMap map[int]messageSpec

func (m typeMap) fillTypeNames() {
	for id, spec := range m {
		spec.TypeName = typeName(spec.EnumName)
		if spec.TypeName == "" {
			delete(m, id)
		} else {
			m[id] = spec
		}
	}
}

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

// processes a single value specification
func processValueSpec(spec *ast.ValueSpec) {
	if spec.Type == nil {
		return
	}

	t, ok := spec.Type.(*ast.Ident)
	if !ok {
		return
	}

	var isEntityType bool
	switch {
	case t.Name == cmdEnumType:
		// it's a message type
	case entityEnumTypes[t.Name]:
		isEntityType = true
	default:
		return
	}

	for i, name := range spec.Names {
		if name.Name == "_" {
			continue
		}

		valExpr := spec.Values[i]
		litExpr, ok := valExpr.(*ast.BasicLit)
		if !ok {
			continue
		}

		if litExpr.Kind != token.INT {
			continue
		}

		n, err := strconv.Atoi(litExpr.Value)
		if err != nil {
			continue
		}

		if isEntityType {
			entityTypes[n] = messageSpec{EnumName: name.Name}
		} else {
			cmdTypes[n] = messageSpec{EnumName: name.Name}
		}
	}
}

// processes a single type specification from the Go source
// e.g.:
//       type Thing int32
//       type Message struct {
//           ...
//       }
func processTypeSpec(spec *ast.TypeSpec) {
	switch tt := spec.Type.(type) {
	case *ast.StructType:
		// the only structes that are defined in our generated proto code are
		// protobuf message types
		messageTypes[spec.Name.Name] = true
	case *ast.Ident:
		switch tt.Name {
		case "int32":
			// all protobuf enums generate int32s in Go
			enumTypes[spec.Name.Name] = true
		}
	}
}

// processes one specification from the Go source
func processSpec(spec ast.Spec) {
	switch t := spec.(type) {
	case *ast.ValueSpec:
		processValueSpec(t)
	case *ast.TypeSpec:
		processTypeSpec(t)
	}
}

// processes one protobuf-generated Go declaration
func processDeclaration(decl ast.Decl) {
	d, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	for _, spec := range d.Specs {
		processSpec(spec)
	}
}

// processes one protobuf-generated Go file
func processFile(name string, fi *ast.File) {
	for _, decl := range fi.Decls {
		processDeclaration(decl)
	}
}

// processes a package of protobuf-generated Go
func processPackage(name string, pkg *ast.Package) {
	for name, fi := range pkg.Files {
		processFile(name, fi)
	}
}

// given an enum name, finds the appropriate message type
func typeName(enumName string) string {
	if name, ok := specials[enumName]; ok {
		return name
	}
	for prefix, replacement := range prefixes {
		if strings.HasPrefix(enumName, prefix) {
			candidate := strings.Replace(enumName, prefix, replacement, 1)
			if messageTypes[candidate] {
				return candidate
			}
		}
	}
	return ""
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		bail(1, "gen should get exactly one argument: the directory to operate on")
	}
	path := flag.Arg(0)

	fs := token.NewFileSet()
	packages, err := parser.ParseDir(fs, path, nil, 0)
	if err != nil {
		bail(1, "go parser error: %v", err)
	}

	for name, pkg := range packages {
		processPackage(name, pkg)
	}

	cmdTypes.fillTypeNames()
	entityTypes.fillTypeNames()

	var ctx = struct {
		Commands typeMap
		Entities typeMap
	}{
		Commands: cmdTypes,
		Entities: entityTypes,
	}

	t, err := template.New("out.go").Parse(tpl)
	if err != nil {
		bail(1, "bad template: %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, ctx); err != nil {
		bail(1, "template error: %v", err)
	}

	source, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(string(buf.Bytes()))
		bail(1, "fmt error: %v", err)
	}

	if err := ioutil.WriteFile("generated.go", source, 0644); err != nil {
		bail(1, "error writing source output: %v", err)
	}
}
