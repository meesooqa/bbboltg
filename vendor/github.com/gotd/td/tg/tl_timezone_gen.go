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

// Timezone represents TL type `timezone#ff9289f5`.
// Timezone information.
//
// See https://core.telegram.org/constructor/timezone for reference.
type Timezone struct {
	// Unique timezone ID.
	ID string
	// Human-readable and localized timezone name.
	Name string
	// UTC offset in seconds, which may be displayed in hh:mm format by the client together
	// with the human-readable name (i.e. $name UTC -01:00).
	UtcOffset int
}

// TimezoneTypeID is TL type id of Timezone.
const TimezoneTypeID = 0xff9289f5

// Ensuring interfaces in compile-time for Timezone.
var (
	_ bin.Encoder     = &Timezone{}
	_ bin.Decoder     = &Timezone{}
	_ bin.BareEncoder = &Timezone{}
	_ bin.BareDecoder = &Timezone{}
)

func (t *Timezone) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ID == "") {
		return false
	}
	if !(t.Name == "") {
		return false
	}
	if !(t.UtcOffset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Timezone) String() string {
	if t == nil {
		return "Timezone(nil)"
	}
	type Alias Timezone
	return fmt.Sprintf("Timezone%+v", Alias(*t))
}

// FillFrom fills Timezone from given interface.
func (t *Timezone) FillFrom(from interface {
	GetID() (value string)
	GetName() (value string)
	GetUtcOffset() (value int)
}) {
	t.ID = from.GetID()
	t.Name = from.GetName()
	t.UtcOffset = from.GetUtcOffset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Timezone) TypeID() uint32 {
	return TimezoneTypeID
}

// TypeName returns name of type in TL schema.
func (*Timezone) TypeName() string {
	return "timezone"
}

// TypeInfo returns info about TL type.
func (t *Timezone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "timezone",
		ID:   TimezoneTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "UtcOffset",
			SchemaName: "utc_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *Timezone) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode timezone#ff9289f5 as nil")
	}
	b.PutID(TimezoneTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *Timezone) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode timezone#ff9289f5 as nil")
	}
	b.PutString(t.ID)
	b.PutString(t.Name)
	b.PutInt(t.UtcOffset)
	return nil
}

// Decode implements bin.Decoder.
func (t *Timezone) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode timezone#ff9289f5 to nil")
	}
	if err := b.ConsumeID(TimezoneTypeID); err != nil {
		return fmt.Errorf("unable to decode timezone#ff9289f5: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *Timezone) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode timezone#ff9289f5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode timezone#ff9289f5: field id: %w", err)
		}
		t.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode timezone#ff9289f5: field name: %w", err)
		}
		t.Name = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode timezone#ff9289f5: field utc_offset: %w", err)
		}
		t.UtcOffset = value
	}
	return nil
}

// GetID returns value of ID field.
func (t *Timezone) GetID() (value string) {
	if t == nil {
		return
	}
	return t.ID
}

// GetName returns value of Name field.
func (t *Timezone) GetName() (value string) {
	if t == nil {
		return
	}
	return t.Name
}

// GetUtcOffset returns value of UtcOffset field.
func (t *Timezone) GetUtcOffset() (value int) {
	if t == nil {
		return
	}
	return t.UtcOffset
}
