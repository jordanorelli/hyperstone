package ent

import (
	"bytes"
	"fmt"
	"github.com/jordanorelli/hyperstone/dota"
)

type field struct {
	name string
	tÿpe
}

func (f *field) fromProto(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) error {
	var_name := env.symbol(int(flat.GetVarNameSym()))
	var_type := env.symbol(int(flat.GetVarTypeSym()))

	var pretty bytes.Buffer
	fmt.Fprintf(&pretty, "{name: %s type: %s", var_name, var_type)
	if flat.BitCount != nil {
		fmt.Fprintf(&pretty, " bits: %d", flat.GetBitCount())
	}
	if flat.LowValue != nil {
		fmt.Fprintf(&pretty, " low: %f", flat.GetLowValue())
	}
	if flat.HighValue != nil {
		fmt.Fprintf(&pretty, " high: %f", flat.GetHighValue())
	}
	if flat.EncodeFlags != nil {
		fmt.Fprintf(&pretty, " flags: %d", flat.GetEncodeFlags())
	}
	if flat.FieldSerializerNameSym != nil {
		fmt.Fprintf(&pretty, " serializer: %s", env.symbol(int(flat.GetFieldSerializerNameSym())))
	}
	if flat.FieldSerializerVersion != nil {
		fmt.Fprintf(&pretty, " s_version: %d", flat.GetFieldSerializerVersion())
	}
	if flat.SendNodeSym != nil {
		fmt.Fprintf(&pretty, " send: %s", env.symbol(int(flat.GetSendNodeSym())))
	}
	if flat.VarEncoderSym != nil {
		fmt.Fprintf(&pretty, " var_enc: %s", env.symbol(int(flat.GetVarEncoderSym())))
	}
	fmt.Fprint(&pretty, "}")

	return fmt.Errorf("unable to parse type: %s", pretty.String())
}
