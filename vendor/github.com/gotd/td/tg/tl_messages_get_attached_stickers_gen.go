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

// MessagesGetAttachedStickersRequest represents TL type `messages.getAttachedStickers#cc5b67cc`.
// Get stickers attached to a photo or video
//
// See https://core.telegram.org/method/messages.getAttachedStickers for reference.
type MessagesGetAttachedStickersRequest struct {
	// Stickered media
	Media InputStickeredMediaClass
}

// MessagesGetAttachedStickersRequestTypeID is TL type id of MessagesGetAttachedStickersRequest.
const MessagesGetAttachedStickersRequestTypeID = 0xcc5b67cc

// Ensuring interfaces in compile-time for MessagesGetAttachedStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetAttachedStickersRequest{}
	_ bin.Decoder     = &MessagesGetAttachedStickersRequest{}
	_ bin.BareEncoder = &MessagesGetAttachedStickersRequest{}
	_ bin.BareDecoder = &MessagesGetAttachedStickersRequest{}
)

func (g *MessagesGetAttachedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Media == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAttachedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetAttachedStickersRequest(nil)"
	}
	type Alias MessagesGetAttachedStickersRequest
	return fmt.Sprintf("MessagesGetAttachedStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetAttachedStickersRequest from given interface.
func (g *MessagesGetAttachedStickersRequest) FillFrom(from interface {
	GetMedia() (value InputStickeredMediaClass)
}) {
	g.Media = from.GetMedia()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAttachedStickersRequest) TypeID() uint32 {
	return MessagesGetAttachedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAttachedStickersRequest) TypeName() string {
	return "messages.getAttachedStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAttachedStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAttachedStickers",
		ID:   MessagesGetAttachedStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Media",
			SchemaName: "media",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetAttachedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachedStickers#cc5b67cc as nil")
	}
	b.PutID(MessagesGetAttachedStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAttachedStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachedStickers#cc5b67cc as nil")
	}
	if g.Media == nil {
		return fmt.Errorf("unable to encode messages.getAttachedStickers#cc5b67cc: field media is nil")
	}
	if err := g.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getAttachedStickers#cc5b67cc: field media: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAttachedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachedStickers#cc5b67cc to nil")
	}
	if err := b.ConsumeID(MessagesGetAttachedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAttachedStickers#cc5b67cc: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAttachedStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachedStickers#cc5b67cc to nil")
	}
	{
		value, err := DecodeInputStickeredMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAttachedStickers#cc5b67cc: field media: %w", err)
		}
		g.Media = value
	}
	return nil
}

// GetMedia returns value of Media field.
func (g *MessagesGetAttachedStickersRequest) GetMedia() (value InputStickeredMediaClass) {
	if g == nil {
		return
	}
	return g.Media
}

// MessagesGetAttachedStickers invokes method messages.getAttachedStickers#cc5b67cc returning error if any.
// Get stickers attached to a photo or video
//
// Possible errors:
//
//	400 MEDIA_EMPTY: The provided media object is invalid.
//
// See https://core.telegram.org/method/messages.getAttachedStickers for reference.
func (c *Client) MessagesGetAttachedStickers(ctx context.Context, media InputStickeredMediaClass) ([]StickerSetCoveredClass, error) {
	var result StickerSetCoveredClassVector

	request := &MessagesGetAttachedStickersRequest{
		Media: media,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []StickerSetCoveredClass(result.Elems), nil
}
