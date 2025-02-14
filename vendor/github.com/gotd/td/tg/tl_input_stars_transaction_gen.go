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

// InputStarsTransaction represents TL type `inputStarsTransaction#206ae6d1`.
// Used to fetch info about a Telegram Star transaction »¹.
//
// Links:
//  1. https://core.telegram.org/api/stars#balance-and-transaction-history
//
// See https://core.telegram.org/constructor/inputStarsTransaction for reference.
type InputStarsTransaction struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, fetches info about the refund transaction for this transaction.
	Refund bool
	// Transaction ID.
	ID string
}

// InputStarsTransactionTypeID is TL type id of InputStarsTransaction.
const InputStarsTransactionTypeID = 0x206ae6d1

// Ensuring interfaces in compile-time for InputStarsTransaction.
var (
	_ bin.Encoder     = &InputStarsTransaction{}
	_ bin.Decoder     = &InputStarsTransaction{}
	_ bin.BareEncoder = &InputStarsTransaction{}
	_ bin.BareDecoder = &InputStarsTransaction{}
)

func (i *InputStarsTransaction) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Refund == false) {
		return false
	}
	if !(i.ID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStarsTransaction) String() string {
	if i == nil {
		return "InputStarsTransaction(nil)"
	}
	type Alias InputStarsTransaction
	return fmt.Sprintf("InputStarsTransaction%+v", Alias(*i))
}

// FillFrom fills InputStarsTransaction from given interface.
func (i *InputStarsTransaction) FillFrom(from interface {
	GetRefund() (value bool)
	GetID() (value string)
}) {
	i.Refund = from.GetRefund()
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStarsTransaction) TypeID() uint32 {
	return InputStarsTransactionTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStarsTransaction) TypeName() string {
	return "inputStarsTransaction"
}

// TypeInfo returns info about TL type.
func (i *InputStarsTransaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStarsTransaction",
		ID:   InputStarsTransactionTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Refund",
			SchemaName: "refund",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputStarsTransaction) SetFlags() {
	if !(i.Refund == false) {
		i.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (i *InputStarsTransaction) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStarsTransaction#206ae6d1 as nil")
	}
	b.PutID(InputStarsTransactionTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStarsTransaction) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStarsTransaction#206ae6d1 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStarsTransaction#206ae6d1: field flags: %w", err)
	}
	b.PutString(i.ID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStarsTransaction) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStarsTransaction#206ae6d1 to nil")
	}
	if err := b.ConsumeID(InputStarsTransactionTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStarsTransaction#206ae6d1: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStarsTransaction) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStarsTransaction#206ae6d1 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStarsTransaction#206ae6d1: field flags: %w", err)
		}
	}
	i.Refund = i.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStarsTransaction#206ae6d1: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// SetRefund sets value of Refund conditional field.
func (i *InputStarsTransaction) SetRefund(value bool) {
	if value {
		i.Flags.Set(0)
		i.Refund = true
	} else {
		i.Flags.Unset(0)
		i.Refund = false
	}
}

// GetRefund returns value of Refund conditional field.
func (i *InputStarsTransaction) GetRefund() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(0)
}

// GetID returns value of ID field.
func (i *InputStarsTransaction) GetID() (value string) {
	if i == nil {
		return
	}
	return i.ID
}
