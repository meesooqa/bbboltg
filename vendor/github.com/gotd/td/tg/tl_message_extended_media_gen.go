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

// MessageExtendedMediaPreview represents TL type `messageExtendedMediaPreview#ad628cc8`.
// Paid media preview for not yet purchased paid media, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/paid-media
//
// See https://core.telegram.org/constructor/messageExtendedMediaPreview for reference.
type MessageExtendedMediaPreview struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Width
	//
	// Use SetW and GetW helpers.
	W int
	// Height
	//
	// Use SetH and GetH helpers.
	H int
	// Extremely low resolution thumbnail¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/files#stripped-thumbnails
	//
	// Use SetThumb and GetThumb helpers.
	Thumb PhotoSizeClass
	// Video duration for videos.
	//
	// Use SetVideoDuration and GetVideoDuration helpers.
	VideoDuration int
}

// MessageExtendedMediaPreviewTypeID is TL type id of MessageExtendedMediaPreview.
const MessageExtendedMediaPreviewTypeID = 0xad628cc8

// construct implements constructor of MessageExtendedMediaClass.
func (m MessageExtendedMediaPreview) construct() MessageExtendedMediaClass { return &m }

// Ensuring interfaces in compile-time for MessageExtendedMediaPreview.
var (
	_ bin.Encoder     = &MessageExtendedMediaPreview{}
	_ bin.Decoder     = &MessageExtendedMediaPreview{}
	_ bin.BareEncoder = &MessageExtendedMediaPreview{}
	_ bin.BareDecoder = &MessageExtendedMediaPreview{}

	_ MessageExtendedMediaClass = &MessageExtendedMediaPreview{}
)

func (m *MessageExtendedMediaPreview) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.W == 0) {
		return false
	}
	if !(m.H == 0) {
		return false
	}
	if !(m.Thumb == nil) {
		return false
	}
	if !(m.VideoDuration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageExtendedMediaPreview) String() string {
	if m == nil {
		return "MessageExtendedMediaPreview(nil)"
	}
	type Alias MessageExtendedMediaPreview
	return fmt.Sprintf("MessageExtendedMediaPreview%+v", Alias(*m))
}

// FillFrom fills MessageExtendedMediaPreview from given interface.
func (m *MessageExtendedMediaPreview) FillFrom(from interface {
	GetW() (value int, ok bool)
	GetH() (value int, ok bool)
	GetThumb() (value PhotoSizeClass, ok bool)
	GetVideoDuration() (value int, ok bool)
}) {
	if val, ok := from.GetW(); ok {
		m.W = val
	}

	if val, ok := from.GetH(); ok {
		m.H = val
	}

	if val, ok := from.GetThumb(); ok {
		m.Thumb = val
	}

	if val, ok := from.GetVideoDuration(); ok {
		m.VideoDuration = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageExtendedMediaPreview) TypeID() uint32 {
	return MessageExtendedMediaPreviewTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageExtendedMediaPreview) TypeName() string {
	return "messageExtendedMediaPreview"
}

// TypeInfo returns info about TL type.
func (m *MessageExtendedMediaPreview) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageExtendedMediaPreview",
		ID:   MessageExtendedMediaPreviewTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "W",
			SchemaName: "w",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "H",
			SchemaName: "h",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "Thumb",
			SchemaName: "thumb",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "VideoDuration",
			SchemaName: "video_duration",
			Null:       !m.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MessageExtendedMediaPreview) SetFlags() {
	if !(m.W == 0) {
		m.Flags.Set(0)
	}
	if !(m.H == 0) {
		m.Flags.Set(0)
	}
	if !(m.Thumb == nil) {
		m.Flags.Set(1)
	}
	if !(m.VideoDuration == 0) {
		m.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (m *MessageExtendedMediaPreview) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageExtendedMediaPreview#ad628cc8 as nil")
	}
	b.PutID(MessageExtendedMediaPreviewTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageExtendedMediaPreview) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageExtendedMediaPreview#ad628cc8 as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageExtendedMediaPreview#ad628cc8: field flags: %w", err)
	}
	if m.Flags.Has(0) {
		b.PutInt(m.W)
	}
	if m.Flags.Has(0) {
		b.PutInt(m.H)
	}
	if m.Flags.Has(1) {
		if m.Thumb == nil {
			return fmt.Errorf("unable to encode messageExtendedMediaPreview#ad628cc8: field thumb is nil")
		}
		if err := m.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageExtendedMediaPreview#ad628cc8: field thumb: %w", err)
		}
	}
	if m.Flags.Has(2) {
		b.PutInt(m.VideoDuration)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageExtendedMediaPreview) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageExtendedMediaPreview#ad628cc8 to nil")
	}
	if err := b.ConsumeID(MessageExtendedMediaPreviewTypeID); err != nil {
		return fmt.Errorf("unable to decode messageExtendedMediaPreview#ad628cc8: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageExtendedMediaPreview) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageExtendedMediaPreview#ad628cc8 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageExtendedMediaPreview#ad628cc8: field flags: %w", err)
		}
	}
	if m.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageExtendedMediaPreview#ad628cc8: field w: %w", err)
		}
		m.W = value
	}
	if m.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageExtendedMediaPreview#ad628cc8: field h: %w", err)
		}
		m.H = value
	}
	if m.Flags.Has(1) {
		value, err := DecodePhotoSize(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageExtendedMediaPreview#ad628cc8: field thumb: %w", err)
		}
		m.Thumb = value
	}
	if m.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageExtendedMediaPreview#ad628cc8: field video_duration: %w", err)
		}
		m.VideoDuration = value
	}
	return nil
}

// SetW sets value of W conditional field.
func (m *MessageExtendedMediaPreview) SetW(value int) {
	m.Flags.Set(0)
	m.W = value
}

// GetW returns value of W conditional field and
// boolean which is true if field was set.
func (m *MessageExtendedMediaPreview) GetW() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.W, true
}

// SetH sets value of H conditional field.
func (m *MessageExtendedMediaPreview) SetH(value int) {
	m.Flags.Set(0)
	m.H = value
}

// GetH returns value of H conditional field and
// boolean which is true if field was set.
func (m *MessageExtendedMediaPreview) GetH() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.H, true
}

// SetThumb sets value of Thumb conditional field.
func (m *MessageExtendedMediaPreview) SetThumb(value PhotoSizeClass) {
	m.Flags.Set(1)
	m.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (m *MessageExtendedMediaPreview) GetThumb() (value PhotoSizeClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.Thumb, true
}

// SetVideoDuration sets value of VideoDuration conditional field.
func (m *MessageExtendedMediaPreview) SetVideoDuration(value int) {
	m.Flags.Set(2)
	m.VideoDuration = value
}

// GetVideoDuration returns value of VideoDuration conditional field and
// boolean which is true if field was set.
func (m *MessageExtendedMediaPreview) GetVideoDuration() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.VideoDuration, true
}

// MessageExtendedMedia represents TL type `messageExtendedMedia#ee479c64`.
// Already purchased paid media, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/paid-media
//
// See https://core.telegram.org/constructor/messageExtendedMedia for reference.
type MessageExtendedMedia struct {
	// The media we purchased.
	Media MessageMediaClass
}

// MessageExtendedMediaTypeID is TL type id of MessageExtendedMedia.
const MessageExtendedMediaTypeID = 0xee479c64

// construct implements constructor of MessageExtendedMediaClass.
func (m MessageExtendedMedia) construct() MessageExtendedMediaClass { return &m }

// Ensuring interfaces in compile-time for MessageExtendedMedia.
var (
	_ bin.Encoder     = &MessageExtendedMedia{}
	_ bin.Decoder     = &MessageExtendedMedia{}
	_ bin.BareEncoder = &MessageExtendedMedia{}
	_ bin.BareDecoder = &MessageExtendedMedia{}

	_ MessageExtendedMediaClass = &MessageExtendedMedia{}
)

func (m *MessageExtendedMedia) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Media == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageExtendedMedia) String() string {
	if m == nil {
		return "MessageExtendedMedia(nil)"
	}
	type Alias MessageExtendedMedia
	return fmt.Sprintf("MessageExtendedMedia%+v", Alias(*m))
}

// FillFrom fills MessageExtendedMedia from given interface.
func (m *MessageExtendedMedia) FillFrom(from interface {
	GetMedia() (value MessageMediaClass)
}) {
	m.Media = from.GetMedia()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageExtendedMedia) TypeID() uint32 {
	return MessageExtendedMediaTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageExtendedMedia) TypeName() string {
	return "messageExtendedMedia"
}

// TypeInfo returns info about TL type.
func (m *MessageExtendedMedia) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageExtendedMedia",
		ID:   MessageExtendedMediaTypeID,
	}
	if m == nil {
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
func (m *MessageExtendedMedia) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageExtendedMedia#ee479c64 as nil")
	}
	b.PutID(MessageExtendedMediaTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageExtendedMedia) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageExtendedMedia#ee479c64 as nil")
	}
	if m.Media == nil {
		return fmt.Errorf("unable to encode messageExtendedMedia#ee479c64: field media is nil")
	}
	if err := m.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageExtendedMedia#ee479c64: field media: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageExtendedMedia) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageExtendedMedia#ee479c64 to nil")
	}
	if err := b.ConsumeID(MessageExtendedMediaTypeID); err != nil {
		return fmt.Errorf("unable to decode messageExtendedMedia#ee479c64: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageExtendedMedia) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageExtendedMedia#ee479c64 to nil")
	}
	{
		value, err := DecodeMessageMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageExtendedMedia#ee479c64: field media: %w", err)
		}
		m.Media = value
	}
	return nil
}

// GetMedia returns value of Media field.
func (m *MessageExtendedMedia) GetMedia() (value MessageMediaClass) {
	if m == nil {
		return
	}
	return m.Media
}

// MessageExtendedMediaClassName is schema name of MessageExtendedMediaClass.
const MessageExtendedMediaClassName = "MessageExtendedMedia"

// MessageExtendedMediaClass represents MessageExtendedMedia generic type.
//
// See https://core.telegram.org/type/MessageExtendedMedia for reference.
//
// Example:
//
//	g, err := tg.DecodeMessageExtendedMedia(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessageExtendedMediaPreview: // messageExtendedMediaPreview#ad628cc8
//	case *tg.MessageExtendedMedia: // messageExtendedMedia#ee479c64
//	default: panic(v)
//	}
type MessageExtendedMediaClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageExtendedMediaClass

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

// DecodeMessageExtendedMedia implements binary de-serialization for MessageExtendedMediaClass.
func DecodeMessageExtendedMedia(buf *bin.Buffer) (MessageExtendedMediaClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageExtendedMediaPreviewTypeID:
		// Decoding messageExtendedMediaPreview#ad628cc8.
		v := MessageExtendedMediaPreview{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageExtendedMediaClass: %w", err)
		}
		return &v, nil
	case MessageExtendedMediaTypeID:
		// Decoding messageExtendedMedia#ee479c64.
		v := MessageExtendedMedia{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageExtendedMediaClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageExtendedMediaClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessageExtendedMedia boxes the MessageExtendedMediaClass providing a helper.
type MessageExtendedMediaBox struct {
	MessageExtendedMedia MessageExtendedMediaClass
}

// Decode implements bin.Decoder for MessageExtendedMediaBox.
func (b *MessageExtendedMediaBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageExtendedMediaBox to nil")
	}
	v, err := DecodeMessageExtendedMedia(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageExtendedMedia = v
	return nil
}

// Encode implements bin.Encode for MessageExtendedMediaBox.
func (b *MessageExtendedMediaBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageExtendedMedia == nil {
		return fmt.Errorf("unable to encode MessageExtendedMediaClass as nil")
	}
	return b.MessageExtendedMedia.Encode(buf)
}
