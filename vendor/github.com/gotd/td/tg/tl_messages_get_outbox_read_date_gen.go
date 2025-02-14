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

// MessagesGetOutboxReadDateRequest represents TL type `messages.getOutboxReadDate#8c4bfe5d`.
// Get the exact read date of one of our messages, sent to a private chat with another
// user.
// Can be only done for private outgoing messages not older than appConfig
// pm_read_date_expire_period »¹.
// If the peer's userFull¹.read_dates_private flag is set, we will not be able to fetch
// the exact read date of messages we send to them, and a USER_PRIVACY_RESTRICTED RPC
// error will be emitted.
// The exact read date of messages might still be unavailable for other reasons, see here
// »² for more info.
// To set userFull³.read_dates_private for ourselves invoke account
// setGlobalPrivacySettings⁴, setting the settings.hide_read_marks flag.
//
// Links:
//  1. https://core.telegram.org/api/config#pm-read-date-expire-period
//  2. https://core.telegram.org/constructor/userFull
//  3. https://core.telegram.org/constructor/globalPrivacySettings
//  4. https://core.telegram.org/constructor/userFull
//  5. https://core.telegram.org/method/account.setGlobalPrivacySettings
//
// See https://core.telegram.org/method/messages.getOutboxReadDate for reference.
type MessagesGetOutboxReadDateRequest struct {
	// The user to whom we sent the message.
	Peer InputPeerClass
	// The message ID.
	MsgID int
}

// MessagesGetOutboxReadDateRequestTypeID is TL type id of MessagesGetOutboxReadDateRequest.
const MessagesGetOutboxReadDateRequestTypeID = 0x8c4bfe5d

// Ensuring interfaces in compile-time for MessagesGetOutboxReadDateRequest.
var (
	_ bin.Encoder     = &MessagesGetOutboxReadDateRequest{}
	_ bin.Decoder     = &MessagesGetOutboxReadDateRequest{}
	_ bin.BareEncoder = &MessagesGetOutboxReadDateRequest{}
	_ bin.BareDecoder = &MessagesGetOutboxReadDateRequest{}
)

func (g *MessagesGetOutboxReadDateRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetOutboxReadDateRequest) String() string {
	if g == nil {
		return "MessagesGetOutboxReadDateRequest(nil)"
	}
	type Alias MessagesGetOutboxReadDateRequest
	return fmt.Sprintf("MessagesGetOutboxReadDateRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetOutboxReadDateRequest from given interface.
func (g *MessagesGetOutboxReadDateRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetOutboxReadDateRequest) TypeID() uint32 {
	return MessagesGetOutboxReadDateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetOutboxReadDateRequest) TypeName() string {
	return "messages.getOutboxReadDate"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetOutboxReadDateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getOutboxReadDate",
		ID:   MessagesGetOutboxReadDateRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetOutboxReadDateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getOutboxReadDate#8c4bfe5d as nil")
	}
	b.PutID(MessagesGetOutboxReadDateRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetOutboxReadDateRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getOutboxReadDate#8c4bfe5d as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getOutboxReadDate#8c4bfe5d: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getOutboxReadDate#8c4bfe5d: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetOutboxReadDateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getOutboxReadDate#8c4bfe5d to nil")
	}
	if err := b.ConsumeID(MessagesGetOutboxReadDateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getOutboxReadDate#8c4bfe5d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetOutboxReadDateRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getOutboxReadDate#8c4bfe5d to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOutboxReadDate#8c4bfe5d: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOutboxReadDate#8c4bfe5d: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetOutboxReadDateRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetOutboxReadDateRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// MessagesGetOutboxReadDate invokes method messages.getOutboxReadDate#8c4bfe5d returning error if any.
// Get the exact read date of one of our messages, sent to a private chat with another
// user.
// Can be only done for private outgoing messages not older than appConfig
// pm_read_date_expire_period »¹.
// If the peer's userFull¹.read_dates_private flag is set, we will not be able to fetch
// the exact read date of messages we send to them, and a USER_PRIVACY_RESTRICTED RPC
// error will be emitted.
// The exact read date of messages might still be unavailable for other reasons, see here
// »² for more info.
// To set userFull³.read_dates_private for ourselves invoke account
// setGlobalPrivacySettings⁴, setting the settings.hide_read_marks flag.
//
// Links:
//  1. https://core.telegram.org/api/config#pm-read-date-expire-period
//  2. https://core.telegram.org/constructor/userFull
//  3. https://core.telegram.org/constructor/globalPrivacySettings
//  4. https://core.telegram.org/constructor/userFull
//  5. https://core.telegram.org/method/account.setGlobalPrivacySettings
//
// Possible errors:
//
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 MESSAGE_NOT_READ_YET: The specified message wasn't read yet.
//	400 MESSAGE_TOO_OLD: The message is too old, the requested information is not available.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this.
//	403 YOUR_PRIVACY_RESTRICTED: You cannot fetch the read date of this message because you have disallowed other users to do so for your messages; to fix, allow other users to see your exact last online date OR purchase a Telegram Premium subscription.
//
// See https://core.telegram.org/method/messages.getOutboxReadDate for reference.
func (c *Client) MessagesGetOutboxReadDate(ctx context.Context, request *MessagesGetOutboxReadDateRequest) (*OutboxReadDate, error) {
	var result OutboxReadDate

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
