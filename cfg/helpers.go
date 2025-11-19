package cfg

import (
	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
)

func AttributeTypeString() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_String_{String_: true}}
}

func AttributeTypeDelimited(s string) *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Delimited{
		Delimited: s,
	}}
}

func AttributeTypeSpaceDelimited() *pb.Attribute_Type {
	return AttributeTypeDelimited(" ")
}

func AttributeTypeCommaDelimited() *pb.Attribute_Type {
	return AttributeTypeDelimited(",")
}

func AttributeTypeKV(keyValueDelim, pairDelim string) *pb.Attribute_Type {
	return &pb.Attribute_Type{
		Type: &pb.Attribute_Type_Kv{
			Kv: &pb.Attribute_KV{
				KeyValueDelimiter: keyValueDelim,
				PairDelimiter:     pairDelim,
			},
		},
	}
}

func AttributeTypeKVColonSemicolon() *pb.Attribute_Type {
	return AttributeTypeKV(":", ";")
}

func AttributeTypeBool() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Bool{
		Bool: true,
	}}
}

func AttributeTypeInt() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Integer{
		Integer: true,
	}}
}

func AttributeTypeNumber() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Number{
		Number: true,
	}}
}

func AttributeTypeChoices(choices ...*pb.Attribute_Choice) *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Choices{
		Choices: &pb.Attribute_Choices{Choices: choices},
	}}
}

func AttributeTypeChoice(name, description string) *pb.Attribute_Choice {
	return &pb.Attribute_Choice{Name: name, Description: description}
}

func AttributeTypeRune() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Rune{
		Rune: true,
	}}
}

func AttributeTypeCustomModifier(name string, hasKey bool, valueType *pb.Attribute_Type, modifiers ...*pb.Attribute_Custom_Modifier) *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Custom{
		Custom: &pb.Attribute_Custom{
			Name:      name,
			Type:      valueType,
			Modifiers: modifiers,
			HasKey:    hasKey,
		},
	}}
}

func AttributeTypeCustom(hasKey bool, valueType *pb.Attribute_Type, modifiers ...*pb.Attribute_Custom_Modifier) *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Custom{
		Custom: &pb.Attribute_Custom{
			Type:      valueType,
			Modifiers: modifiers,
			HasKey:    hasKey,
		},
	}}
}

func AttributeTypeJSON() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_Json{
		Json: true,
	}}
}

func AttributeTypeDurationMs() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_DurationMs{
		DurationMs: true,
	}}
}

func AttributeTypeDurationSec() *pb.Attribute_Type {
	return &pb.Attribute_Type{Type: &pb.Attribute_Type_DurationSec{
		DurationSec: true,
	}}
}
