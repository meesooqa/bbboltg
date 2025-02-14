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

// ChannelsCreateForumTopicRequest represents TL type `channels.createForumTopic#f40c0224`.
// Create a forum topic¹; requires manage_topics rights².
//
// Links:
//  1. https://core.telegram.org/api/forum
//  2. https://core.telegram.org/api/rights
//
// See https://core.telegram.org/method/channels.createForumTopic for reference.
type ChannelsCreateForumTopicRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The forum¹
	//
	// Links:
	//  1) https://core.telegram.org/api/forum
	Channel InputChannelClass
	// Topic title (maximum UTF-8 length: 128)
	Title string
	// If no custom emoji icon is specified, specifies the color of the fallback topic icon
	// (RGB), one of 0x6FB9F0, 0xFFD67E, 0xCB86DB, 0x8EEE98, 0xFF93B2, or 0xFB6F5F.
	//
	// Use SetIconColor and GetIconColor helpers.
	IconColor int
	// ID of the custom emoji¹ used as topic icon. Telegram Premium² users can use any
	// custom emoji, other users can only use the custom emojis contained in the
	// inputStickerSetEmojiDefaultTopicIcons³ emoji pack.
	//
	// Links:
	//  1) https://core.telegram.org/api/custom-emoji
	//  2) https://core.telegram.org/api/premium
	//  3) https://core.telegram.org/constructor/inputStickerSetEmojiDefaultTopicIcons
	//
	// Use SetIconEmojiID and GetIconEmojiID helpers.
	IconEmojiID int64
	// Unique client message ID to prevent duplicate sending of the same event
	RandomID int64
	// Create the topic as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// ChannelsCreateForumTopicRequestTypeID is TL type id of ChannelsCreateForumTopicRequest.
const ChannelsCreateForumTopicRequestTypeID = 0xf40c0224

// Ensuring interfaces in compile-time for ChannelsCreateForumTopicRequest.
var (
	_ bin.Encoder     = &ChannelsCreateForumTopicRequest{}
	_ bin.Decoder     = &ChannelsCreateForumTopicRequest{}
	_ bin.BareEncoder = &ChannelsCreateForumTopicRequest{}
	_ bin.BareDecoder = &ChannelsCreateForumTopicRequest{}
)

func (c *ChannelsCreateForumTopicRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Channel == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.IconColor == 0) {
		return false
	}
	if !(c.IconEmojiID == 0) {
		return false
	}
	if !(c.RandomID == 0) {
		return false
	}
	if !(c.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsCreateForumTopicRequest) String() string {
	if c == nil {
		return "ChannelsCreateForumTopicRequest(nil)"
	}
	type Alias ChannelsCreateForumTopicRequest
	return fmt.Sprintf("ChannelsCreateForumTopicRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsCreateForumTopicRequest from given interface.
func (c *ChannelsCreateForumTopicRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetTitle() (value string)
	GetIconColor() (value int, ok bool)
	GetIconEmojiID() (value int64, ok bool)
	GetRandomID() (value int64)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	c.Channel = from.GetChannel()
	c.Title = from.GetTitle()
	if val, ok := from.GetIconColor(); ok {
		c.IconColor = val
	}

	if val, ok := from.GetIconEmojiID(); ok {
		c.IconEmojiID = val
	}

	c.RandomID = from.GetRandomID()
	if val, ok := from.GetSendAs(); ok {
		c.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsCreateForumTopicRequest) TypeID() uint32 {
	return ChannelsCreateForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsCreateForumTopicRequest) TypeName() string {
	return "channels.createForumTopic"
}

// TypeInfo returns info about TL type.
func (c *ChannelsCreateForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.createForumTopic",
		ID:   ChannelsCreateForumTopicRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "IconColor",
			SchemaName: "icon_color",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "IconEmojiID",
			SchemaName: "icon_emoji_id",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !c.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChannelsCreateForumTopicRequest) SetFlags() {
	if !(c.IconColor == 0) {
		c.Flags.Set(0)
	}
	if !(c.IconEmojiID == 0) {
		c.Flags.Set(3)
	}
	if !(c.SendAs == nil) {
		c.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (c *ChannelsCreateForumTopicRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.createForumTopic#f40c0224 as nil")
	}
	b.PutID(ChannelsCreateForumTopicRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsCreateForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.createForumTopic#f40c0224 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.createForumTopic#f40c0224: field flags: %w", err)
	}
	if c.Channel == nil {
		return fmt.Errorf("unable to encode channels.createForumTopic#f40c0224: field channel is nil")
	}
	if err := c.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.createForumTopic#f40c0224: field channel: %w", err)
	}
	b.PutString(c.Title)
	if c.Flags.Has(0) {
		b.PutInt(c.IconColor)
	}
	if c.Flags.Has(3) {
		b.PutLong(c.IconEmojiID)
	}
	b.PutLong(c.RandomID)
	if c.Flags.Has(2) {
		if c.SendAs == nil {
			return fmt.Errorf("unable to encode channels.createForumTopic#f40c0224: field send_as is nil")
		}
		if err := c.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.createForumTopic#f40c0224: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelsCreateForumTopicRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.createForumTopic#f40c0224 to nil")
	}
	if err := b.ConsumeID(ChannelsCreateForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsCreateForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.createForumTopic#f40c0224 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field channel: %w", err)
		}
		c.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field icon_color: %w", err)
		}
		c.IconColor = value
	}
	if c.Flags.Has(3) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field icon_emoji_id: %w", err)
		}
		c.IconEmojiID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field random_id: %w", err)
		}
		c.RandomID = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.createForumTopic#f40c0224: field send_as: %w", err)
		}
		c.SendAs = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (c *ChannelsCreateForumTopicRequest) GetChannel() (value InputChannelClass) {
	if c == nil {
		return
	}
	return c.Channel
}

// GetTitle returns value of Title field.
func (c *ChannelsCreateForumTopicRequest) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// SetIconColor sets value of IconColor conditional field.
func (c *ChannelsCreateForumTopicRequest) SetIconColor(value int) {
	c.Flags.Set(0)
	c.IconColor = value
}

// GetIconColor returns value of IconColor conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateForumTopicRequest) GetIconColor() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.IconColor, true
}

// SetIconEmojiID sets value of IconEmojiID conditional field.
func (c *ChannelsCreateForumTopicRequest) SetIconEmojiID(value int64) {
	c.Flags.Set(3)
	c.IconEmojiID = value
}

// GetIconEmojiID returns value of IconEmojiID conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateForumTopicRequest) GetIconEmojiID() (value int64, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.IconEmojiID, true
}

// GetRandomID returns value of RandomID field.
func (c *ChannelsCreateForumTopicRequest) GetRandomID() (value int64) {
	if c == nil {
		return
	}
	return c.RandomID
}

// SetSendAs sets value of SendAs conditional field.
func (c *ChannelsCreateForumTopicRequest) SetSendAs(value InputPeerClass) {
	c.Flags.Set(2)
	c.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateForumTopicRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.SendAs, true
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (c *ChannelsCreateForumTopicRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return c.Channel.AsNotEmpty()
}

// ChannelsCreateForumTopic invokes method channels.createForumTopic#f40c0224 returning error if any.
// Create a forum topic¹; requires manage_topics rights².
//
// Links:
//  1. https://core.telegram.org/api/forum
//  2. https://core.telegram.org/api/rights
//
// Possible errors:
//
//	400 CHANNEL_FORUM_MISSING: This supergroup is not a forum.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	403 PREMIUM_ACCOUNT_REQUIRED: A premium account is required to execute this action.
//	400 TOPIC_TITLE_EMPTY: The specified topic title is empty.
//
// See https://core.telegram.org/method/channels.createForumTopic for reference.
// Can be used by bots.
func (c *Client) ChannelsCreateForumTopic(ctx context.Context, request *ChannelsCreateForumTopicRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
