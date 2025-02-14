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

// PhotosUploadContactProfilePhotoRequest represents TL type `photos.uploadContactProfilePhoto#e14c4a71`.
// Upload a custom profile picture for a contact, or suggest a new profile picture to a
// contact.
// The file, video and video_emoji_markup flags are mutually exclusive.
//
// See https://core.telegram.org/method/photos.uploadContactProfilePhoto for reference.
type PhotosUploadContactProfilePhotoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, will send a messageActionSuggestProfilePhoto¹ service message to user_id,
	// suggesting them to use the specified profile picture; otherwise, will set a personal
	// profile picture for the user (only visible to the current user).
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messageActionSuggestProfilePhoto
	Suggest bool
	// If set, removes a previously set personal profile picture (does not affect suggested
	// profile pictures, to remove them simply deleted the messageActionSuggestProfilePhoto¹
	// service message with messages.deleteMessages²).
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messageActionSuggestProfilePhoto
	//  2) https://core.telegram.org/method/messages.deleteMessages
	Save bool
	// The contact
	UserID InputUserClass
	// Profile photo
	//
	// Use SetFile and GetFile helpers.
	File InputFileClass
	// Animated profile picture¹ video
	//
	// Links:
	//  1) https://core.telegram.org/api/files#animated-profile-pictures
	//
	// Use SetVideo and GetVideo helpers.
	Video InputFileClass
	// Floating point UNIX timestamp in seconds, indicating the frame of the video/sticker
	// that should be used as static preview; can only be used if video or video_emoji_markup
	// is set.
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
	// Animated sticker profile picture, must contain either a videoSizeEmojiMarkup¹ or a
	// videoSizeStickerMarkup² constructor.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/videoSizeEmojiMarkup
	//  2) https://core.telegram.org/constructor/videoSizeStickerMarkup
	//
	// Use SetVideoEmojiMarkup and GetVideoEmojiMarkup helpers.
	VideoEmojiMarkup VideoSizeClass
}

// PhotosUploadContactProfilePhotoRequestTypeID is TL type id of PhotosUploadContactProfilePhotoRequest.
const PhotosUploadContactProfilePhotoRequestTypeID = 0xe14c4a71

// Ensuring interfaces in compile-time for PhotosUploadContactProfilePhotoRequest.
var (
	_ bin.Encoder     = &PhotosUploadContactProfilePhotoRequest{}
	_ bin.Decoder     = &PhotosUploadContactProfilePhotoRequest{}
	_ bin.BareEncoder = &PhotosUploadContactProfilePhotoRequest{}
	_ bin.BareDecoder = &PhotosUploadContactProfilePhotoRequest{}
)

func (u *PhotosUploadContactProfilePhotoRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Suggest == false) {
		return false
	}
	if !(u.Save == false) {
		return false
	}
	if !(u.UserID == nil) {
		return false
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.Video == nil) {
		return false
	}
	if !(u.VideoStartTs == 0) {
		return false
	}
	if !(u.VideoEmojiMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *PhotosUploadContactProfilePhotoRequest) String() string {
	if u == nil {
		return "PhotosUploadContactProfilePhotoRequest(nil)"
	}
	type Alias PhotosUploadContactProfilePhotoRequest
	return fmt.Sprintf("PhotosUploadContactProfilePhotoRequest%+v", Alias(*u))
}

// FillFrom fills PhotosUploadContactProfilePhotoRequest from given interface.
func (u *PhotosUploadContactProfilePhotoRequest) FillFrom(from interface {
	GetSuggest() (value bool)
	GetSave() (value bool)
	GetUserID() (value InputUserClass)
	GetFile() (value InputFileClass, ok bool)
	GetVideo() (value InputFileClass, ok bool)
	GetVideoStartTs() (value float64, ok bool)
	GetVideoEmojiMarkup() (value VideoSizeClass, ok bool)
}) {
	u.Suggest = from.GetSuggest()
	u.Save = from.GetSave()
	u.UserID = from.GetUserID()
	if val, ok := from.GetFile(); ok {
		u.File = val
	}

	if val, ok := from.GetVideo(); ok {
		u.Video = val
	}

	if val, ok := from.GetVideoStartTs(); ok {
		u.VideoStartTs = val
	}

	if val, ok := from.GetVideoEmojiMarkup(); ok {
		u.VideoEmojiMarkup = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosUploadContactProfilePhotoRequest) TypeID() uint32 {
	return PhotosUploadContactProfilePhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosUploadContactProfilePhotoRequest) TypeName() string {
	return "photos.uploadContactProfilePhoto"
}

// TypeInfo returns info about TL type.
func (u *PhotosUploadContactProfilePhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.uploadContactProfilePhoto",
		ID:   PhotosUploadContactProfilePhotoRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Suggest",
			SchemaName: "suggest",
			Null:       !u.Flags.Has(3),
		},
		{
			Name:       "Save",
			SchemaName: "save",
			Null:       !u.Flags.Has(4),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "File",
			SchemaName: "file",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "VideoEmojiMarkup",
			SchemaName: "video_emoji_markup",
			Null:       !u.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *PhotosUploadContactProfilePhotoRequest) SetFlags() {
	if !(u.Suggest == false) {
		u.Flags.Set(3)
	}
	if !(u.Save == false) {
		u.Flags.Set(4)
	}
	if !(u.File == nil) {
		u.Flags.Set(0)
	}
	if !(u.Video == nil) {
		u.Flags.Set(1)
	}
	if !(u.VideoStartTs == 0) {
		u.Flags.Set(2)
	}
	if !(u.VideoEmojiMarkup == nil) {
		u.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (u *PhotosUploadContactProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.uploadContactProfilePhoto#e14c4a71 as nil")
	}
	b.PutID(PhotosUploadContactProfilePhotoRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *PhotosUploadContactProfilePhotoRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.uploadContactProfilePhoto#e14c4a71 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field flags: %w", err)
	}
	if u.UserID == nil {
		return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field user_id is nil")
	}
	if err := u.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field user_id: %w", err)
	}
	if u.Flags.Has(0) {
		if u.File == nil {
			return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field file is nil")
		}
		if err := u.File.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field file: %w", err)
		}
	}
	if u.Flags.Has(1) {
		if u.Video == nil {
			return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field video is nil")
		}
		if err := u.Video.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field video: %w", err)
		}
	}
	if u.Flags.Has(2) {
		b.PutDouble(u.VideoStartTs)
	}
	if u.Flags.Has(5) {
		if u.VideoEmojiMarkup == nil {
			return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field video_emoji_markup is nil")
		}
		if err := u.VideoEmojiMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadContactProfilePhoto#e14c4a71: field video_emoji_markup: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *PhotosUploadContactProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.uploadContactProfilePhoto#e14c4a71 to nil")
	}
	if err := b.ConsumeID(PhotosUploadContactProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *PhotosUploadContactProfilePhotoRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.uploadContactProfilePhoto#e14c4a71 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: field flags: %w", err)
		}
	}
	u.Suggest = u.Flags.Has(3)
	u.Save = u.Flags.Has(4)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: field user_id: %w", err)
		}
		u.UserID = value
	}
	if u.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: field file: %w", err)
		}
		u.File = value
	}
	if u.Flags.Has(1) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: field video: %w", err)
		}
		u.Video = value
	}
	if u.Flags.Has(2) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: field video_start_ts: %w", err)
		}
		u.VideoStartTs = value
	}
	if u.Flags.Has(5) {
		value, err := DecodeVideoSize(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadContactProfilePhoto#e14c4a71: field video_emoji_markup: %w", err)
		}
		u.VideoEmojiMarkup = value
	}
	return nil
}

// SetSuggest sets value of Suggest conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) SetSuggest(value bool) {
	if value {
		u.Flags.Set(3)
		u.Suggest = true
	} else {
		u.Flags.Unset(3)
		u.Suggest = false
	}
}

// GetSuggest returns value of Suggest conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) GetSuggest() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(3)
}

// SetSave sets value of Save conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) SetSave(value bool) {
	if value {
		u.Flags.Set(4)
		u.Save = true
	} else {
		u.Flags.Unset(4)
		u.Save = false
	}
}

// GetSave returns value of Save conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) GetSave() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(4)
}

// GetUserID returns value of UserID field.
func (u *PhotosUploadContactProfilePhotoRequest) GetUserID() (value InputUserClass) {
	if u == nil {
		return
	}
	return u.UserID
}

// SetFile sets value of File conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) SetFile(value InputFileClass) {
	u.Flags.Set(0)
	u.File = value
}

// GetFile returns value of File conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadContactProfilePhotoRequest) GetFile() (value InputFileClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.File, true
}

// SetVideo sets value of Video conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) SetVideo(value InputFileClass) {
	u.Flags.Set(1)
	u.Video = value
}

// GetVideo returns value of Video conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadContactProfilePhotoRequest) GetVideo() (value InputFileClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.Video, true
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) SetVideoStartTs(value float64) {
	u.Flags.Set(2)
	u.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadContactProfilePhotoRequest) GetVideoStartTs() (value float64, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.VideoStartTs, true
}

// SetVideoEmojiMarkup sets value of VideoEmojiMarkup conditional field.
func (u *PhotosUploadContactProfilePhotoRequest) SetVideoEmojiMarkup(value VideoSizeClass) {
	u.Flags.Set(5)
	u.VideoEmojiMarkup = value
}

// GetVideoEmojiMarkup returns value of VideoEmojiMarkup conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadContactProfilePhotoRequest) GetVideoEmojiMarkup() (value VideoSizeClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(5) {
		return value, false
	}
	return u.VideoEmojiMarkup, true
}

// PhotosUploadContactProfilePhoto invokes method photos.uploadContactProfilePhoto#e14c4a71 returning error if any.
// Upload a custom profile picture for a contact, or suggest a new profile picture to a
// contact.
// The file, video and video_emoji_markup flags are mutually exclusive.
//
// Possible errors:
//
//	400 CONTACT_MISSING: The specified user is not a contact.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/photos.uploadContactProfilePhoto for reference.
func (c *Client) PhotosUploadContactProfilePhoto(ctx context.Context, request *PhotosUploadContactProfilePhotoRequest) (*PhotosPhoto, error) {
	var result PhotosPhoto

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
