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

// MessagesFeaturedStickersNotModified represents TL type `messages.featuredStickersNotModified#c6dc0c66`.
// Featured stickers haven't changed
//
// See https://core.telegram.org/constructor/messages.featuredStickersNotModified for reference.
type MessagesFeaturedStickersNotModified struct {
	// Total number of featured stickers
	Count int
}

// MessagesFeaturedStickersNotModifiedTypeID is TL type id of MessagesFeaturedStickersNotModified.
const MessagesFeaturedStickersNotModifiedTypeID = 0xc6dc0c66

// construct implements constructor of MessagesFeaturedStickersClass.
func (f MessagesFeaturedStickersNotModified) construct() MessagesFeaturedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFeaturedStickersNotModified.
var (
	_ bin.Encoder     = &MessagesFeaturedStickersNotModified{}
	_ bin.Decoder     = &MessagesFeaturedStickersNotModified{}
	_ bin.BareEncoder = &MessagesFeaturedStickersNotModified{}
	_ bin.BareDecoder = &MessagesFeaturedStickersNotModified{}

	_ MessagesFeaturedStickersClass = &MessagesFeaturedStickersNotModified{}
)

func (f *MessagesFeaturedStickersNotModified) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFeaturedStickersNotModified) String() string {
	if f == nil {
		return "MessagesFeaturedStickersNotModified(nil)"
	}
	type Alias MessagesFeaturedStickersNotModified
	return fmt.Sprintf("MessagesFeaturedStickersNotModified%+v", Alias(*f))
}

// FillFrom fills MessagesFeaturedStickersNotModified from given interface.
func (f *MessagesFeaturedStickersNotModified) FillFrom(from interface {
	GetCount() (value int)
}) {
	f.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFeaturedStickersNotModified) TypeID() uint32 {
	return MessagesFeaturedStickersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFeaturedStickersNotModified) TypeName() string {
	return "messages.featuredStickersNotModified"
}

// TypeInfo returns info about TL type.
func (f *MessagesFeaturedStickersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.featuredStickersNotModified",
		ID:   MessagesFeaturedStickersNotModifiedTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFeaturedStickersNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickersNotModified#c6dc0c66 as nil")
	}
	b.PutID(MessagesFeaturedStickersNotModifiedTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFeaturedStickersNotModified) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickersNotModified#c6dc0c66 as nil")
	}
	b.PutInt(f.Count)
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFeaturedStickersNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickersNotModified#c6dc0c66 to nil")
	}
	if err := b.ConsumeID(MessagesFeaturedStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.featuredStickersNotModified#c6dc0c66: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFeaturedStickersNotModified) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickersNotModified#c6dc0c66 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickersNotModified#c6dc0c66: field count: %w", err)
		}
		f.Count = value
	}
	return nil
}

// GetCount returns value of Count field.
func (f *MessagesFeaturedStickersNotModified) GetCount() (value int) {
	if f == nil {
		return
	}
	return f.Count
}

// MessagesFeaturedStickers represents TL type `messages.featuredStickers#be382906`.
// Featured stickersets
//
// See https://core.telegram.org/constructor/messages.featuredStickers for reference.
type MessagesFeaturedStickers struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a premium stickerset
	Premium bool
	// Hash used for caching, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// Total number of featured stickers
	Count int
	// Featured stickersets
	Sets []StickerSetCoveredClass
	// IDs of new featured stickersets
	Unread []int64
}

// MessagesFeaturedStickersTypeID is TL type id of MessagesFeaturedStickers.
const MessagesFeaturedStickersTypeID = 0xbe382906

// construct implements constructor of MessagesFeaturedStickersClass.
func (f MessagesFeaturedStickers) construct() MessagesFeaturedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFeaturedStickers.
var (
	_ bin.Encoder     = &MessagesFeaturedStickers{}
	_ bin.Decoder     = &MessagesFeaturedStickers{}
	_ bin.BareEncoder = &MessagesFeaturedStickers{}
	_ bin.BareDecoder = &MessagesFeaturedStickers{}

	_ MessagesFeaturedStickersClass = &MessagesFeaturedStickers{}
)

func (f *MessagesFeaturedStickers) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Flags.Zero()) {
		return false
	}
	if !(f.Premium == false) {
		return false
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Count == 0) {
		return false
	}
	if !(f.Sets == nil) {
		return false
	}
	if !(f.Unread == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFeaturedStickers) String() string {
	if f == nil {
		return "MessagesFeaturedStickers(nil)"
	}
	type Alias MessagesFeaturedStickers
	return fmt.Sprintf("MessagesFeaturedStickers%+v", Alias(*f))
}

// FillFrom fills MessagesFeaturedStickers from given interface.
func (f *MessagesFeaturedStickers) FillFrom(from interface {
	GetPremium() (value bool)
	GetHash() (value int64)
	GetCount() (value int)
	GetSets() (value []StickerSetCoveredClass)
	GetUnread() (value []int64)
}) {
	f.Premium = from.GetPremium()
	f.Hash = from.GetHash()
	f.Count = from.GetCount()
	f.Sets = from.GetSets()
	f.Unread = from.GetUnread()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFeaturedStickers) TypeID() uint32 {
	return MessagesFeaturedStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFeaturedStickers) TypeName() string {
	return "messages.featuredStickers"
}

// TypeInfo returns info about TL type.
func (f *MessagesFeaturedStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.featuredStickers",
		ID:   MessagesFeaturedStickersTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Premium",
			SchemaName: "premium",
			Null:       !f.Flags.Has(0),
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
		{
			Name:       "Unread",
			SchemaName: "unread",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (f *MessagesFeaturedStickers) SetFlags() {
	if !(f.Premium == false) {
		f.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (f *MessagesFeaturedStickers) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickers#be382906 as nil")
	}
	b.PutID(MessagesFeaturedStickersTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFeaturedStickers) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickers#be382906 as nil")
	}
	f.SetFlags()
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.featuredStickers#be382906: field flags: %w", err)
	}
	b.PutLong(f.Hash)
	b.PutInt(f.Count)
	b.PutVectorHeader(len(f.Sets))
	for idx, v := range f.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.featuredStickers#be382906: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.featuredStickers#be382906: field sets element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Unread))
	for _, v := range f.Unread {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFeaturedStickers) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickers#be382906 to nil")
	}
	if err := b.ConsumeID(MessagesFeaturedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.featuredStickers#be382906: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFeaturedStickers) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickers#be382906 to nil")
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field flags: %w", err)
		}
	}
	f.Premium = f.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field count: %w", err)
		}
		f.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field sets: %w", err)
		}

		if headerLen > 0 {
			f.Sets = make([]StickerSetCoveredClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field sets: %w", err)
			}
			f.Sets = append(f.Sets, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field unread: %w", err)
		}

		if headerLen > 0 {
			f.Unread = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.featuredStickers#be382906: field unread: %w", err)
			}
			f.Unread = append(f.Unread, value)
		}
	}
	return nil
}

// SetPremium sets value of Premium conditional field.
func (f *MessagesFeaturedStickers) SetPremium(value bool) {
	if value {
		f.Flags.Set(0)
		f.Premium = true
	} else {
		f.Flags.Unset(0)
		f.Premium = false
	}
}

// GetPremium returns value of Premium conditional field.
func (f *MessagesFeaturedStickers) GetPremium() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(0)
}

// GetHash returns value of Hash field.
func (f *MessagesFeaturedStickers) GetHash() (value int64) {
	if f == nil {
		return
	}
	return f.Hash
}

// GetCount returns value of Count field.
func (f *MessagesFeaturedStickers) GetCount() (value int) {
	if f == nil {
		return
	}
	return f.Count
}

// GetSets returns value of Sets field.
func (f *MessagesFeaturedStickers) GetSets() (value []StickerSetCoveredClass) {
	if f == nil {
		return
	}
	return f.Sets
}

// GetUnread returns value of Unread field.
func (f *MessagesFeaturedStickers) GetUnread() (value []int64) {
	if f == nil {
		return
	}
	return f.Unread
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassArray helper.
func (f *MessagesFeaturedStickers) MapSets() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(f.Sets)
}

// MessagesFeaturedStickersClassName is schema name of MessagesFeaturedStickersClass.
const MessagesFeaturedStickersClassName = "messages.FeaturedStickers"

// MessagesFeaturedStickersClass represents messages.FeaturedStickers generic type.
//
// See https://core.telegram.org/type/messages.FeaturedStickers for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesFeaturedStickers(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesFeaturedStickersNotModified: // messages.featuredStickersNotModified#c6dc0c66
//	case *tg.MessagesFeaturedStickers: // messages.featuredStickers#be382906
//	default: panic(v)
//	}
type MessagesFeaturedStickersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesFeaturedStickersClass

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

	// Total number of featured stickers
	GetCount() (value int)

	// AsModified tries to map MessagesFeaturedStickersClass to MessagesFeaturedStickers.
	AsModified() (*MessagesFeaturedStickers, bool)
}

// AsModified tries to map MessagesFeaturedStickersNotModified to MessagesFeaturedStickers.
func (f *MessagesFeaturedStickersNotModified) AsModified() (*MessagesFeaturedStickers, bool) {
	return nil, false
}

// AsModified tries to map MessagesFeaturedStickers to MessagesFeaturedStickers.
func (f *MessagesFeaturedStickers) AsModified() (*MessagesFeaturedStickers, bool) {
	return f, true
}

// DecodeMessagesFeaturedStickers implements binary de-serialization for MessagesFeaturedStickersClass.
func DecodeMessagesFeaturedStickers(buf *bin.Buffer) (MessagesFeaturedStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFeaturedStickersNotModifiedTypeID:
		// Decoding messages.featuredStickersNotModified#c6dc0c66.
		v := MessagesFeaturedStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", err)
		}
		return &v, nil
	case MessagesFeaturedStickersTypeID:
		// Decoding messages.featuredStickers#be382906.
		v := MessagesFeaturedStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFeaturedStickers boxes the MessagesFeaturedStickersClass providing a helper.
type MessagesFeaturedStickersBox struct {
	FeaturedStickers MessagesFeaturedStickersClass
}

// Decode implements bin.Decoder for MessagesFeaturedStickersBox.
func (b *MessagesFeaturedStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFeaturedStickersBox to nil")
	}
	v, err := DecodeMessagesFeaturedStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FeaturedStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesFeaturedStickersBox.
func (b *MessagesFeaturedStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FeaturedStickers == nil {
		return fmt.Errorf("unable to encode MessagesFeaturedStickersClass as nil")
	}
	return b.FeaturedStickers.Encode(buf)
}
