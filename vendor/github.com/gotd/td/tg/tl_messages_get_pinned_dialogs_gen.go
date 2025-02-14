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

// MessagesGetPinnedDialogsRequest represents TL type `messages.getPinnedDialogs#d6b94df2`.
// Get pinned dialogs
//
// See https://core.telegram.org/method/messages.getPinnedDialogs for reference.
type MessagesGetPinnedDialogsRequest struct {
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// MessagesGetPinnedDialogsRequestTypeID is TL type id of MessagesGetPinnedDialogsRequest.
const MessagesGetPinnedDialogsRequestTypeID = 0xd6b94df2

// Ensuring interfaces in compile-time for MessagesGetPinnedDialogsRequest.
var (
	_ bin.Encoder     = &MessagesGetPinnedDialogsRequest{}
	_ bin.Decoder     = &MessagesGetPinnedDialogsRequest{}
	_ bin.BareEncoder = &MessagesGetPinnedDialogsRequest{}
	_ bin.BareDecoder = &MessagesGetPinnedDialogsRequest{}
)

func (g *MessagesGetPinnedDialogsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPinnedDialogsRequest) String() string {
	if g == nil {
		return "MessagesGetPinnedDialogsRequest(nil)"
	}
	type Alias MessagesGetPinnedDialogsRequest
	return fmt.Sprintf("MessagesGetPinnedDialogsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetPinnedDialogsRequest from given interface.
func (g *MessagesGetPinnedDialogsRequest) FillFrom(from interface {
	GetFolderID() (value int)
}) {
	g.FolderID = from.GetFolderID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetPinnedDialogsRequest) TypeID() uint32 {
	return MessagesGetPinnedDialogsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetPinnedDialogsRequest) TypeName() string {
	return "messages.getPinnedDialogs"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetPinnedDialogsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getPinnedDialogs",
		ID:   MessagesGetPinnedDialogsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetPinnedDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPinnedDialogs#d6b94df2 as nil")
	}
	b.PutID(MessagesGetPinnedDialogsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetPinnedDialogsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPinnedDialogs#d6b94df2 as nil")
	}
	b.PutInt(g.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPinnedDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPinnedDialogs#d6b94df2 to nil")
	}
	if err := b.ConsumeID(MessagesGetPinnedDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPinnedDialogs#d6b94df2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetPinnedDialogsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPinnedDialogs#d6b94df2 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPinnedDialogs#d6b94df2: field folder_id: %w", err)
		}
		g.FolderID = value
	}
	return nil
}

// GetFolderID returns value of FolderID field.
func (g *MessagesGetPinnedDialogsRequest) GetFolderID() (value int) {
	if g == nil {
		return
	}
	return g.FolderID
}

// MessagesGetPinnedDialogs invokes method messages.getPinnedDialogs#d6b94df2 returning error if any.
// Get pinned dialogs
//
// Possible errors:
//
//	400 FOLDER_ID_INVALID: Invalid folder ID.
//
// See https://core.telegram.org/method/messages.getPinnedDialogs for reference.
func (c *Client) MessagesGetPinnedDialogs(ctx context.Context, folderid int) (*MessagesPeerDialogs, error) {
	var result MessagesPeerDialogs

	request := &MessagesGetPinnedDialogsRequest{
		FolderID: folderid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
