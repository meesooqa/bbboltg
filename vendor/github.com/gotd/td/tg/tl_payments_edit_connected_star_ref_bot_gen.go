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

// PaymentsEditConnectedStarRefBotRequest represents TL type `payments.editConnectedStarRefBot#e4fca4a3`.
// Leave a bot's affiliate program »¹
//
// Links:
//  1. https://core.telegram.org/api/bots/referrals#becoming-an-affiliate
//
// See https://core.telegram.org/method/payments.editConnectedStarRefBot for reference.
type PaymentsEditConnectedStarRefBotRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, leaves the bot's affiliate program
	Revoked bool
	// The peer that was affiliated
	Peer InputPeerClass
	// The affiliate link to revoke
	Link string
}

// PaymentsEditConnectedStarRefBotRequestTypeID is TL type id of PaymentsEditConnectedStarRefBotRequest.
const PaymentsEditConnectedStarRefBotRequestTypeID = 0xe4fca4a3

// Ensuring interfaces in compile-time for PaymentsEditConnectedStarRefBotRequest.
var (
	_ bin.Encoder     = &PaymentsEditConnectedStarRefBotRequest{}
	_ bin.Decoder     = &PaymentsEditConnectedStarRefBotRequest{}
	_ bin.BareEncoder = &PaymentsEditConnectedStarRefBotRequest{}
	_ bin.BareDecoder = &PaymentsEditConnectedStarRefBotRequest{}
)

func (e *PaymentsEditConnectedStarRefBotRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Revoked == false) {
		return false
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PaymentsEditConnectedStarRefBotRequest) String() string {
	if e == nil {
		return "PaymentsEditConnectedStarRefBotRequest(nil)"
	}
	type Alias PaymentsEditConnectedStarRefBotRequest
	return fmt.Sprintf("PaymentsEditConnectedStarRefBotRequest%+v", Alias(*e))
}

// FillFrom fills PaymentsEditConnectedStarRefBotRequest from given interface.
func (e *PaymentsEditConnectedStarRefBotRequest) FillFrom(from interface {
	GetRevoked() (value bool)
	GetPeer() (value InputPeerClass)
	GetLink() (value string)
}) {
	e.Revoked = from.GetRevoked()
	e.Peer = from.GetPeer()
	e.Link = from.GetLink()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsEditConnectedStarRefBotRequest) TypeID() uint32 {
	return PaymentsEditConnectedStarRefBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsEditConnectedStarRefBotRequest) TypeName() string {
	return "payments.editConnectedStarRefBot"
}

// TypeInfo returns info about TL type.
func (e *PaymentsEditConnectedStarRefBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.editConnectedStarRefBot",
		ID:   PaymentsEditConnectedStarRefBotRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoked",
			SchemaName: "revoked",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *PaymentsEditConnectedStarRefBotRequest) SetFlags() {
	if !(e.Revoked == false) {
		e.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (e *PaymentsEditConnectedStarRefBotRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode payments.editConnectedStarRefBot#e4fca4a3 as nil")
	}
	b.PutID(PaymentsEditConnectedStarRefBotRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PaymentsEditConnectedStarRefBotRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode payments.editConnectedStarRefBot#e4fca4a3 as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.editConnectedStarRefBot#e4fca4a3: field flags: %w", err)
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode payments.editConnectedStarRefBot#e4fca4a3: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.editConnectedStarRefBot#e4fca4a3: field peer: %w", err)
	}
	b.PutString(e.Link)
	return nil
}

// Decode implements bin.Decoder.
func (e *PaymentsEditConnectedStarRefBotRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode payments.editConnectedStarRefBot#e4fca4a3 to nil")
	}
	if err := b.ConsumeID(PaymentsEditConnectedStarRefBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.editConnectedStarRefBot#e4fca4a3: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PaymentsEditConnectedStarRefBotRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode payments.editConnectedStarRefBot#e4fca4a3 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.editConnectedStarRefBot#e4fca4a3: field flags: %w", err)
		}
	}
	e.Revoked = e.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.editConnectedStarRefBot#e4fca4a3: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.editConnectedStarRefBot#e4fca4a3: field link: %w", err)
		}
		e.Link = value
	}
	return nil
}

// SetRevoked sets value of Revoked conditional field.
func (e *PaymentsEditConnectedStarRefBotRequest) SetRevoked(value bool) {
	if value {
		e.Flags.Set(0)
		e.Revoked = true
	} else {
		e.Flags.Unset(0)
		e.Revoked = false
	}
}

// GetRevoked returns value of Revoked conditional field.
func (e *PaymentsEditConnectedStarRefBotRequest) GetRevoked() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (e *PaymentsEditConnectedStarRefBotRequest) GetPeer() (value InputPeerClass) {
	if e == nil {
		return
	}
	return e.Peer
}

// GetLink returns value of Link field.
func (e *PaymentsEditConnectedStarRefBotRequest) GetLink() (value string) {
	if e == nil {
		return
	}
	return e.Link
}

// PaymentsEditConnectedStarRefBot invokes method payments.editConnectedStarRefBot#e4fca4a3 returning error if any.
// Leave a bot's affiliate program »¹
//
// Links:
//  1. https://core.telegram.org/api/bots/referrals#becoming-an-affiliate
//
// Possible errors:
//
//	400 STARREF_HASH_REVOKED: The specified affiliate link was already revoked.
//
// See https://core.telegram.org/method/payments.editConnectedStarRefBot for reference.
func (c *Client) PaymentsEditConnectedStarRefBot(ctx context.Context, request *PaymentsEditConnectedStarRefBotRequest) (*PaymentsConnectedStarRefBots, error) {
	var result PaymentsConnectedStarRefBots

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
