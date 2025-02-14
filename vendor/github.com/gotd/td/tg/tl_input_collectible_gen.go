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

// InputCollectibleUsername represents TL type `inputCollectibleUsername#e39460a9`.
// Represents a username fragment collectible¹
//
// Links:
//  1. https://core.telegram.org/api/fragment
//
// See https://core.telegram.org/constructor/inputCollectibleUsername for reference.
type InputCollectibleUsername struct {
	// Username
	Username string
}

// InputCollectibleUsernameTypeID is TL type id of InputCollectibleUsername.
const InputCollectibleUsernameTypeID = 0xe39460a9

// construct implements constructor of InputCollectibleClass.
func (i InputCollectibleUsername) construct() InputCollectibleClass { return &i }

// Ensuring interfaces in compile-time for InputCollectibleUsername.
var (
	_ bin.Encoder     = &InputCollectibleUsername{}
	_ bin.Decoder     = &InputCollectibleUsername{}
	_ bin.BareEncoder = &InputCollectibleUsername{}
	_ bin.BareDecoder = &InputCollectibleUsername{}

	_ InputCollectibleClass = &InputCollectibleUsername{}
)

func (i *InputCollectibleUsername) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCollectibleUsername) String() string {
	if i == nil {
		return "InputCollectibleUsername(nil)"
	}
	type Alias InputCollectibleUsername
	return fmt.Sprintf("InputCollectibleUsername%+v", Alias(*i))
}

// FillFrom fills InputCollectibleUsername from given interface.
func (i *InputCollectibleUsername) FillFrom(from interface {
	GetUsername() (value string)
}) {
	i.Username = from.GetUsername()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCollectibleUsername) TypeID() uint32 {
	return InputCollectibleUsernameTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCollectibleUsername) TypeName() string {
	return "inputCollectibleUsername"
}

// TypeInfo returns info about TL type.
func (i *InputCollectibleUsername) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCollectibleUsername",
		ID:   InputCollectibleUsernameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCollectibleUsername) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCollectibleUsername#e39460a9 as nil")
	}
	b.PutID(InputCollectibleUsernameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCollectibleUsername) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCollectibleUsername#e39460a9 as nil")
	}
	b.PutString(i.Username)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCollectibleUsername) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCollectibleUsername#e39460a9 to nil")
	}
	if err := b.ConsumeID(InputCollectibleUsernameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCollectibleUsername#e39460a9: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCollectibleUsername) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCollectibleUsername#e39460a9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputCollectibleUsername#e39460a9: field username: %w", err)
		}
		i.Username = value
	}
	return nil
}

// GetUsername returns value of Username field.
func (i *InputCollectibleUsername) GetUsername() (value string) {
	if i == nil {
		return
	}
	return i.Username
}

// InputCollectiblePhone represents TL type `inputCollectiblePhone#a2e214a4`.
// Represents a phone number fragment collectible¹
//
// Links:
//  1. https://core.telegram.org/api/fragment
//
// See https://core.telegram.org/constructor/inputCollectiblePhone for reference.
type InputCollectiblePhone struct {
	// Phone number
	Phone string
}

// InputCollectiblePhoneTypeID is TL type id of InputCollectiblePhone.
const InputCollectiblePhoneTypeID = 0xa2e214a4

// construct implements constructor of InputCollectibleClass.
func (i InputCollectiblePhone) construct() InputCollectibleClass { return &i }

// Ensuring interfaces in compile-time for InputCollectiblePhone.
var (
	_ bin.Encoder     = &InputCollectiblePhone{}
	_ bin.Decoder     = &InputCollectiblePhone{}
	_ bin.BareEncoder = &InputCollectiblePhone{}
	_ bin.BareDecoder = &InputCollectiblePhone{}

	_ InputCollectibleClass = &InputCollectiblePhone{}
)

func (i *InputCollectiblePhone) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Phone == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCollectiblePhone) String() string {
	if i == nil {
		return "InputCollectiblePhone(nil)"
	}
	type Alias InputCollectiblePhone
	return fmt.Sprintf("InputCollectiblePhone%+v", Alias(*i))
}

// FillFrom fills InputCollectiblePhone from given interface.
func (i *InputCollectiblePhone) FillFrom(from interface {
	GetPhone() (value string)
}) {
	i.Phone = from.GetPhone()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCollectiblePhone) TypeID() uint32 {
	return InputCollectiblePhoneTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCollectiblePhone) TypeName() string {
	return "inputCollectiblePhone"
}

// TypeInfo returns info about TL type.
func (i *InputCollectiblePhone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCollectiblePhone",
		ID:   InputCollectiblePhoneTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Phone",
			SchemaName: "phone",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCollectiblePhone) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCollectiblePhone#a2e214a4 as nil")
	}
	b.PutID(InputCollectiblePhoneTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCollectiblePhone) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCollectiblePhone#a2e214a4 as nil")
	}
	b.PutString(i.Phone)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCollectiblePhone) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCollectiblePhone#a2e214a4 to nil")
	}
	if err := b.ConsumeID(InputCollectiblePhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCollectiblePhone#a2e214a4: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCollectiblePhone) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCollectiblePhone#a2e214a4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputCollectiblePhone#a2e214a4: field phone: %w", err)
		}
		i.Phone = value
	}
	return nil
}

// GetPhone returns value of Phone field.
func (i *InputCollectiblePhone) GetPhone() (value string) {
	if i == nil {
		return
	}
	return i.Phone
}

// InputCollectibleClassName is schema name of InputCollectibleClass.
const InputCollectibleClassName = "InputCollectible"

// InputCollectibleClass represents InputCollectible generic type.
//
// See https://core.telegram.org/type/InputCollectible for reference.
//
// Example:
//
//	g, err := tg.DecodeInputCollectible(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputCollectibleUsername: // inputCollectibleUsername#e39460a9
//	case *tg.InputCollectiblePhone: // inputCollectiblePhone#a2e214a4
//	default: panic(v)
//	}
type InputCollectibleClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputCollectibleClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputCollectible implements binary de-serialization for InputCollectibleClass.
func DecodeInputCollectible(buf *bin.Buffer) (InputCollectibleClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputCollectibleUsernameTypeID:
		// Decoding inputCollectibleUsername#e39460a9.
		v := InputCollectibleUsername{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCollectibleClass: %w", err)
		}
		return &v, nil
	case InputCollectiblePhoneTypeID:
		// Decoding inputCollectiblePhone#a2e214a4.
		v := InputCollectiblePhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCollectibleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputCollectibleClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputCollectible boxes the InputCollectibleClass providing a helper.
type InputCollectibleBox struct {
	InputCollectible InputCollectibleClass
}

// Decode implements bin.Decoder for InputCollectibleBox.
func (b *InputCollectibleBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputCollectibleBox to nil")
	}
	v, err := DecodeInputCollectible(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputCollectible = v
	return nil
}

// Encode implements bin.Encode for InputCollectibleBox.
func (b *InputCollectibleBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputCollectible == nil {
		return fmt.Errorf("unable to encode InputCollectibleClass as nil")
	}
	return b.InputCollectible.Encode(buf)
}
