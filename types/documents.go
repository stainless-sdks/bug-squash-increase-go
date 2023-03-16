package types

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type Document struct {
	// The Document identifier.
	ID *string `pjson:"id"`
	// The type of document.
	Category *DocumentCategory `pjson:"category"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// The identifier of the Entity the document was generated for.
	EntityID *string `pjson:"entity_id"`
	// The identifier of the File containing the Document's contents.
	FileID *string `pjson:"file_id"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type       *DocumentType          `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Document using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Document into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Document) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Document identifier.
func (r Document) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The type of document.
func (r Document) GetCategory() (Category DocumentCategory) {
	if r.Category != nil {
		Category = *r.Category
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
// Document was created.
func (r Document) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the Entity the document was generated for.
func (r Document) GetEntityID() (EntityID string) {
	if r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

// The identifier of the File containing the Document's contents.
func (r Document) GetFileID() (FileID string) {
	if r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `document`.
func (r Document) GetType() (Type DocumentType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Document) String() (result string) {
	return fmt.Sprintf("&Document{ID:%s Category:%s CreatedAt:%s EntityID:%s FileID:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.Category), core.FmtP(r.CreatedAt), core.FmtP(r.EntityID), core.FmtP(r.FileID), core.FmtP(r.Type))
}

type DocumentCategory string

const (
	DocumentCategoryForm_1099Int DocumentCategory = "form_1099_int"
)

type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

type DocumentListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID   *string                      `query:"entity_id"`
	Category   *DocumentListParamsCategory  `query:"category"`
	CreatedAt  *DocumentListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentListParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentListParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DocumentListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DocumentListParams into a url.Values of the query parameters
// associated with this value
func (r *DocumentListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r DocumentListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r DocumentListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Documents to ones belonging to the specified Entity.
func (r DocumentListParams) GetEntityID() (EntityID string) {
	if r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

func (r DocumentListParams) GetCategory() (Category DocumentListParamsCategory) {
	if r.Category != nil {
		Category = *r.Category
	}
	return
}

func (r DocumentListParams) GetCreatedAt() (CreatedAt DocumentListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r DocumentListParams) String() (result string) {
	return fmt.Sprintf("&DocumentListParams{Cursor:%s Limit:%s EntityID:%s Category:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.EntityID), r.Category, r.CreatedAt)
}

type DocumentListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]DocumentListParamsCategoryIn `pjson:"in"`
	jsonFields map[string]interface{}          `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentListParamsCategory
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DocumentListParamsCategory) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentListParamsCategory into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DocumentListParamsCategory) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DocumentListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *DocumentListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r DocumentListParamsCategory) GetIn() (In []DocumentListParamsCategoryIn) {
	if r.In != nil {
		In = *r.In
	}
	return
}

func (r DocumentListParamsCategory) String() (result string) {
	return fmt.Sprintf("&DocumentListParamsCategory{In:%s}", core.Fmt(r.In))
}

type DocumentListParamsCategoryIn string

const (
	DocumentListParamsCategoryInForm_1099Int DocumentListParamsCategoryIn = "form_1099_int"
)

type DocumentListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *time.Time `pjson:"after" format:"2006-01-02T15:04:05Z07:00"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *time.Time `pjson:"before" format:"2006-01-02T15:04:05Z07:00"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *time.Time `pjson:"on_or_after" format:"2006-01-02T15:04:05Z07:00"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *time.Time             `pjson:"on_or_before" format:"2006-01-02T15:04:05Z07:00"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentListParamsCreatedAt
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DocumentListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentListParamsCreatedAt into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DocumentListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DocumentListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *DocumentListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r DocumentListParamsCreatedAt) GetAfter() (After time.Time) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r DocumentListParamsCreatedAt) GetBefore() (Before time.Time) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r DocumentListParamsCreatedAt) GetOnOrAfter() (OnOrAfter time.Time) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r DocumentListParamsCreatedAt) GetOnOrBefore() (OnOrBefore time.Time) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r DocumentListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DocumentListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type DocumentList struct {
	// The contents of the list.
	Data *[]Document `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DocumentList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DocumentList into a url.Values of the query parameters
// associated with this value
func (r *DocumentList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r DocumentList) GetData() (Data []Document) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r DocumentList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r DocumentList) String() (result string) {
	return fmt.Sprintf("&DocumentList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type DocumentsPage struct {
	*pagination.Page[Document]
}

func (r *DocumentsPage) Document() *Document {
	return r.Current()
}

func (r *DocumentsPage) NextPage() (*DocumentsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &DocumentsPage{page}, nil
	}
}