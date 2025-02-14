// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ContactStatusVector is a box for Vector<ContactStatus>
type ContactStatusVector struct {
	// Elements of Vector<ContactStatus>
	Elems []ContactStatus
}

// ContactStatusVectorTypeID is TL type id of ContactStatusVector.
const ContactStatusVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for ContactStatusVector.
var (
	_ bin.Encoder     = &ContactStatusVector{}
	_ bin.Decoder     = &ContactStatusVector{}
	_ bin.BareEncoder = &ContactStatusVector{}
	_ bin.BareDecoder = &ContactStatusVector{}
)

func (vec *ContactStatusVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *ContactStatusVector) String() string {
	if vec == nil {
		return "ContactStatusVector(nil)"
	}
	type Alias ContactStatusVector
	return fmt.Sprintf("ContactStatusVector%+v", Alias(*vec))
}

// FillFrom fills ContactStatusVector from given interface.
func (vec *ContactStatusVector) FillFrom(from interface {
	GetElems() (value []ContactStatus)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactStatusVector) TypeID() uint32 {
	return ContactStatusVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactStatusVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *ContactStatusVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   ContactStatusVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *ContactStatusVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<ContactStatus> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *ContactStatusVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<ContactStatus> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<ContactStatus>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *ContactStatusVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<ContactStatus> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *ContactStatusVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<ContactStatus> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<ContactStatus>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]ContactStatus, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ContactStatus
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<ContactStatus>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *ContactStatusVector) GetElems() (value []ContactStatus) {
	if vec == nil {
		return
	}
	return vec.Elems
}
