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

// StoriesAllStoriesNotModified represents TL type `stories.allStoriesNotModified#1158fe3e`.
// The list of active (or active and hidden) stories¹ has not changed.
//
// Links:
//  1. https://core.telegram.org/api/stories#watching-stories
//
// See https://core.telegram.org/constructor/stories.allStoriesNotModified for reference.
type StoriesAllStoriesNotModified struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// State to use to ask for updates
	State string
	// Current stealth mode¹ information
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#stealth-mode
	StealthMode StoriesStealthMode
}

// StoriesAllStoriesNotModifiedTypeID is TL type id of StoriesAllStoriesNotModified.
const StoriesAllStoriesNotModifiedTypeID = 0x1158fe3e

// construct implements constructor of StoriesAllStoriesClass.
func (a StoriesAllStoriesNotModified) construct() StoriesAllStoriesClass { return &a }

// Ensuring interfaces in compile-time for StoriesAllStoriesNotModified.
var (
	_ bin.Encoder     = &StoriesAllStoriesNotModified{}
	_ bin.Decoder     = &StoriesAllStoriesNotModified{}
	_ bin.BareEncoder = &StoriesAllStoriesNotModified{}
	_ bin.BareDecoder = &StoriesAllStoriesNotModified{}

	_ StoriesAllStoriesClass = &StoriesAllStoriesNotModified{}
)

func (a *StoriesAllStoriesNotModified) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.State == "") {
		return false
	}
	if !(a.StealthMode.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *StoriesAllStoriesNotModified) String() string {
	if a == nil {
		return "StoriesAllStoriesNotModified(nil)"
	}
	type Alias StoriesAllStoriesNotModified
	return fmt.Sprintf("StoriesAllStoriesNotModified%+v", Alias(*a))
}

// FillFrom fills StoriesAllStoriesNotModified from given interface.
func (a *StoriesAllStoriesNotModified) FillFrom(from interface {
	GetState() (value string)
	GetStealthMode() (value StoriesStealthMode)
}) {
	a.State = from.GetState()
	a.StealthMode = from.GetStealthMode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesAllStoriesNotModified) TypeID() uint32 {
	return StoriesAllStoriesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesAllStoriesNotModified) TypeName() string {
	return "stories.allStoriesNotModified"
}

// TypeInfo returns info about TL type.
func (a *StoriesAllStoriesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.allStoriesNotModified",
		ID:   StoriesAllStoriesNotModifiedTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "StealthMode",
			SchemaName: "stealth_mode",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *StoriesAllStoriesNotModified) SetFlags() {
}

// Encode implements bin.Encoder.
func (a *StoriesAllStoriesNotModified) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.allStoriesNotModified#1158fe3e as nil")
	}
	b.PutID(StoriesAllStoriesNotModifiedTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *StoriesAllStoriesNotModified) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.allStoriesNotModified#1158fe3e as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.allStoriesNotModified#1158fe3e: field flags: %w", err)
	}
	b.PutString(a.State)
	if err := a.StealthMode.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.allStoriesNotModified#1158fe3e: field stealth_mode: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *StoriesAllStoriesNotModified) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.allStoriesNotModified#1158fe3e to nil")
	}
	if err := b.ConsumeID(StoriesAllStoriesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.allStoriesNotModified#1158fe3e: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *StoriesAllStoriesNotModified) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.allStoriesNotModified#1158fe3e to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.allStoriesNotModified#1158fe3e: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.allStoriesNotModified#1158fe3e: field state: %w", err)
		}
		a.State = value
	}
	{
		if err := a.StealthMode.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.allStoriesNotModified#1158fe3e: field stealth_mode: %w", err)
		}
	}
	return nil
}

// GetState returns value of State field.
func (a *StoriesAllStoriesNotModified) GetState() (value string) {
	if a == nil {
		return
	}
	return a.State
}

// GetStealthMode returns value of StealthMode field.
func (a *StoriesAllStoriesNotModified) GetStealthMode() (value StoriesStealthMode) {
	if a == nil {
		return
	}
	return a.StealthMode
}

// StoriesAllStories represents TL type `stories.allStories#6efc5e81`.
// Full list of active (or active and hidden) stories¹.
//
// Links:
//  1. https://core.telegram.org/api/stories#watching-stories
//
// See https://core.telegram.org/constructor/stories.allStories for reference.
type StoriesAllStories struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether more results can be fetched as described here »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#watching-stories
	HasMore bool
	// Total number of active (or active and hidden) stories
	Count int
	// State to use for pagination
	State string
	// Stories
	PeerStories []PeerStories
	// Mentioned chats
	Chats []ChatClass
	// Mentioned users
	Users []UserClass
	// Current stealth mode¹ information
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#stealth-mode
	StealthMode StoriesStealthMode
}

// StoriesAllStoriesTypeID is TL type id of StoriesAllStories.
const StoriesAllStoriesTypeID = 0x6efc5e81

// construct implements constructor of StoriesAllStoriesClass.
func (a StoriesAllStories) construct() StoriesAllStoriesClass { return &a }

// Ensuring interfaces in compile-time for StoriesAllStories.
var (
	_ bin.Encoder     = &StoriesAllStories{}
	_ bin.Decoder     = &StoriesAllStories{}
	_ bin.BareEncoder = &StoriesAllStories{}
	_ bin.BareDecoder = &StoriesAllStories{}

	_ StoriesAllStoriesClass = &StoriesAllStories{}
)

func (a *StoriesAllStories) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.HasMore == false) {
		return false
	}
	if !(a.Count == 0) {
		return false
	}
	if !(a.State == "") {
		return false
	}
	if !(a.PeerStories == nil) {
		return false
	}
	if !(a.Chats == nil) {
		return false
	}
	if !(a.Users == nil) {
		return false
	}
	if !(a.StealthMode.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *StoriesAllStories) String() string {
	if a == nil {
		return "StoriesAllStories(nil)"
	}
	type Alias StoriesAllStories
	return fmt.Sprintf("StoriesAllStories%+v", Alias(*a))
}

// FillFrom fills StoriesAllStories from given interface.
func (a *StoriesAllStories) FillFrom(from interface {
	GetHasMore() (value bool)
	GetCount() (value int)
	GetState() (value string)
	GetPeerStories() (value []PeerStories)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetStealthMode() (value StoriesStealthMode)
}) {
	a.HasMore = from.GetHasMore()
	a.Count = from.GetCount()
	a.State = from.GetState()
	a.PeerStories = from.GetPeerStories()
	a.Chats = from.GetChats()
	a.Users = from.GetUsers()
	a.StealthMode = from.GetStealthMode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesAllStories) TypeID() uint32 {
	return StoriesAllStoriesTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesAllStories) TypeName() string {
	return "stories.allStories"
}

// TypeInfo returns info about TL type.
func (a *StoriesAllStories) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.allStories",
		ID:   StoriesAllStoriesTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasMore",
			SchemaName: "has_more",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "PeerStories",
			SchemaName: "peer_stories",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "StealthMode",
			SchemaName: "stealth_mode",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *StoriesAllStories) SetFlags() {
	if !(a.HasMore == false) {
		a.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (a *StoriesAllStories) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.allStories#6efc5e81 as nil")
	}
	b.PutID(StoriesAllStoriesTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *StoriesAllStories) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.allStories#6efc5e81 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field flags: %w", err)
	}
	b.PutInt(a.Count)
	b.PutString(a.State)
	b.PutVectorHeader(len(a.PeerStories))
	for idx, v := range a.PeerStories {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field peer_stories element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Chats))
	for idx, v := range a.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Users))
	for idx, v := range a.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field users element with index %d: %w", idx, err)
		}
	}
	if err := a.StealthMode.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.allStories#6efc5e81: field stealth_mode: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *StoriesAllStories) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.allStories#6efc5e81 to nil")
	}
	if err := b.ConsumeID(StoriesAllStoriesTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.allStories#6efc5e81: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *StoriesAllStories) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.allStories#6efc5e81 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field flags: %w", err)
		}
	}
	a.HasMore = a.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field count: %w", err)
		}
		a.Count = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field state: %w", err)
		}
		a.State = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field peer_stories: %w", err)
		}

		if headerLen > 0 {
			a.PeerStories = make([]PeerStories, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PeerStories
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field peer_stories: %w", err)
			}
			a.PeerStories = append(a.PeerStories, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field chats: %w", err)
		}

		if headerLen > 0 {
			a.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field chats: %w", err)
			}
			a.Chats = append(a.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field users: %w", err)
		}

		if headerLen > 0 {
			a.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field users: %w", err)
			}
			a.Users = append(a.Users, value)
		}
	}
	{
		if err := a.StealthMode.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.allStories#6efc5e81: field stealth_mode: %w", err)
		}
	}
	return nil
}

// SetHasMore sets value of HasMore conditional field.
func (a *StoriesAllStories) SetHasMore(value bool) {
	if value {
		a.Flags.Set(0)
		a.HasMore = true
	} else {
		a.Flags.Unset(0)
		a.HasMore = false
	}
}

// GetHasMore returns value of HasMore conditional field.
func (a *StoriesAllStories) GetHasMore() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// GetCount returns value of Count field.
func (a *StoriesAllStories) GetCount() (value int) {
	if a == nil {
		return
	}
	return a.Count
}

// GetState returns value of State field.
func (a *StoriesAllStories) GetState() (value string) {
	if a == nil {
		return
	}
	return a.State
}

// GetPeerStories returns value of PeerStories field.
func (a *StoriesAllStories) GetPeerStories() (value []PeerStories) {
	if a == nil {
		return
	}
	return a.PeerStories
}

// GetChats returns value of Chats field.
func (a *StoriesAllStories) GetChats() (value []ChatClass) {
	if a == nil {
		return
	}
	return a.Chats
}

// GetUsers returns value of Users field.
func (a *StoriesAllStories) GetUsers() (value []UserClass) {
	if a == nil {
		return
	}
	return a.Users
}

// GetStealthMode returns value of StealthMode field.
func (a *StoriesAllStories) GetStealthMode() (value StoriesStealthMode) {
	if a == nil {
		return
	}
	return a.StealthMode
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (a *StoriesAllStories) MapChats() (value ChatClassArray) {
	return ChatClassArray(a.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (a *StoriesAllStories) MapUsers() (value UserClassArray) {
	return UserClassArray(a.Users)
}

// StoriesAllStoriesClassName is schema name of StoriesAllStoriesClass.
const StoriesAllStoriesClassName = "stories.AllStories"

// StoriesAllStoriesClass represents stories.AllStories generic type.
//
// See https://core.telegram.org/type/stories.AllStories for reference.
//
// Example:
//
//	g, err := tg.DecodeStoriesAllStories(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.StoriesAllStoriesNotModified: // stories.allStoriesNotModified#1158fe3e
//	case *tg.StoriesAllStories: // stories.allStories#6efc5e81
//	default: panic(v)
//	}
type StoriesAllStoriesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoriesAllStoriesClass

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

	// State to use to ask for updates
	GetState() (value string)

	// Current stealth mode¹ information
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#stealth-mode
	GetStealthMode() (value StoriesStealthMode)

	// AsModified tries to map StoriesAllStoriesClass to StoriesAllStories.
	AsModified() (*StoriesAllStories, bool)
}

// AsModified tries to map StoriesAllStoriesNotModified to StoriesAllStories.
func (a *StoriesAllStoriesNotModified) AsModified() (*StoriesAllStories, bool) {
	return nil, false
}

// AsModified tries to map StoriesAllStories to StoriesAllStories.
func (a *StoriesAllStories) AsModified() (*StoriesAllStories, bool) {
	return a, true
}

// DecodeStoriesAllStories implements binary de-serialization for StoriesAllStoriesClass.
func DecodeStoriesAllStories(buf *bin.Buffer) (StoriesAllStoriesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoriesAllStoriesNotModifiedTypeID:
		// Decoding stories.allStoriesNotModified#1158fe3e.
		v := StoriesAllStoriesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoriesAllStoriesClass: %w", err)
		}
		return &v, nil
	case StoriesAllStoriesTypeID:
		// Decoding stories.allStories#6efc5e81.
		v := StoriesAllStories{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoriesAllStoriesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoriesAllStoriesClass: %w", bin.NewUnexpectedID(id))
	}
}

// StoriesAllStories boxes the StoriesAllStoriesClass providing a helper.
type StoriesAllStoriesBox struct {
	AllStories StoriesAllStoriesClass
}

// Decode implements bin.Decoder for StoriesAllStoriesBox.
func (b *StoriesAllStoriesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoriesAllStoriesBox to nil")
	}
	v, err := DecodeStoriesAllStories(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AllStories = v
	return nil
}

// Encode implements bin.Encode for StoriesAllStoriesBox.
func (b *StoriesAllStoriesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AllStories == nil {
		return fmt.Errorf("unable to encode StoriesAllStoriesClass as nil")
	}
	return b.AllStories.Encode(buf)
}
