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

// Bytes represents TL type `bytes#e937bb82`.
//
// See https://core.telegram.org/constructor/bytes for reference.
type Bytes struct {
}

// BytesTypeID is TL type id of Bytes.
const BytesTypeID = 0xe937bb82

// Ensuring interfaces in compile-time for Bytes.
var (
	_ bin.Encoder     = &Bytes{}
	_ bin.Decoder     = &Bytes{}
	_ bin.BareEncoder = &Bytes{}
	_ bin.BareDecoder = &Bytes{}
)

func (b *Bytes) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *Bytes) String() string {
	if b == nil {
		return "Bytes(nil)"
	}
	type Alias Bytes
	return fmt.Sprintf("Bytes%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Bytes) TypeID() uint32 {
	return BytesTypeID
}

// TypeName returns name of type in TL schema.
func (*Bytes) TypeName() string {
	return "bytes"
}

// TypeInfo returns info about TL type.
func (b *Bytes) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bytes",
		ID:   BytesTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *Bytes) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytes#e937bb82 as nil")
	}
	buf.PutID(BytesTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *Bytes) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytes#e937bb82 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *Bytes) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytes#e937bb82 to nil")
	}
	if err := buf.ConsumeID(BytesTypeID); err != nil {
		return fmt.Errorf("unable to decode bytes#e937bb82: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *Bytes) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytes#e937bb82 to nil")
	}
	return nil
}
