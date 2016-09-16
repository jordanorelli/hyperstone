package ent

import (
	"bytes"
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

type decodeFn func(bit.Reader, *Dict) (interface{}, error)

// a Type in the entity type system. Note that not every type is necessarily an
// entity type, since there are necessarily primitives, and above that, arrays
// and generics.
type Type interface {
	// creates a new value of the given type.
	New(...interface{}) interface{}

	// name is primarily of interest for debugging
	Name() string

	// whether or not the produced values are expected to be slotted.
	Slotted() bool

	// reads a value of this type off of the bit reader
	Read(bit.Reader, *Dict) (interface{}, error)
}

func parseType(n *Namespace, flat *dota.ProtoFlattenedSerializerFieldT) (Type, error) {
	Debug.Printf("parseType: %s", prettyFlatField(n, flat))
	type_name := n.Symbol(int(flat.GetVarTypeSym())).String()

	if prim, ok := primitives[type_name]; ok {
		Debug.Printf("  parseType: found primitive with name %s", type_name)
		return &prim, nil
	}

	if n.HasClass(type_name) {
		Debug.Printf("  parseType: found class with name %s", type_name)
		return nil, nil
	}

	switch type_name {
	case "float32", "CNetworkedQuantizedFloat":
		Debug.Printf("  parseType: parsing as float type")
		return parseFloatType(n, flat)
	case "CGameSceneNodeHandle":
		return &Handle{name: "CGameSceneNodeHandle"}, nil
	case "QAngle":
		return parseQAngleType(n, flat)
	}

	Debug.Printf("  parseType: failed")
	// type ProtoFlattenedSerializerFieldT struct {
	// 	VarTypeSym             *int32
	// 	VarNameSym             *int32
	// 	VarEncoderSym          *int32
	// 	FieldSerializerNameSym *int32
	// 	FieldSerializerVersion *int32
	// 	BitCount               *int32
	// 	LowValue               *float32
	// 	HighValue              *float32
	// 	EncodeFlags            *int32
	// 	SendNodeSym            *int32
	// }
	return nil, nil
}

func prettyFlatField(n *Namespace, ff *dota.ProtoFlattenedSerializerFieldT) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "{type: %s", n.Symbol(int(ff.GetVarTypeSym())))
	fmt.Fprintf(&buf, " name: %s", n.Symbol(int(ff.GetVarNameSym())))
	if ff.BitCount != nil {
		fmt.Fprintf(&buf, " bits: %d", ff.GetBitCount())
	}
	if ff.LowValue != nil {
		fmt.Fprintf(&buf, " low: %f", ff.GetLowValue())
	}
	if ff.HighValue != nil {
		fmt.Fprintf(&buf, " high: %f", ff.GetHighValue())
	}
	if ff.EncodeFlags != nil {
		fmt.Fprintf(&buf, " flags: %d", ff.GetEncodeFlags())
	}
	if ff.FieldSerializerNameSym != nil {
		fmt.Fprintf(&buf, " serializer: %s", n.Symbol(int(ff.GetFieldSerializerNameSym())))
	}
	if ff.FieldSerializerVersion != nil {
		fmt.Fprintf(&buf, " version: %d", ff.GetFieldSerializerVersion())
	}
	if ff.SendNodeSym != nil {
		fmt.Fprintf(&buf, " send: %s", n.Symbol(int(ff.GetSendNodeSym())))
	}
	if ff.VarEncoderSym != nil {
		fmt.Fprintf(&buf, " encoder: %s", n.Symbol(int(ff.GetVarEncoderSym())))
	}
	fmt.Fprintf(&buf, "}")
	return buf.String()
}
