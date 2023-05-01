package increase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type FileService struct {
	Options []option.RequestOption
}

func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a File
func (r *FileService) Get(ctx context.Context, file_id string, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("files/%s", file_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Files
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *shared.Page[File], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Files
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *shared.PageAutoPager[File] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type File struct {
	// The time the File was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The File's identifier.
	ID string `json:"id,required"`
	// What the File will be used for. We may add additional possible values for this
	// enum over time; your application should be able to handle such additions
	// gracefully.
	Purpose FilePurpose `json:"purpose,required"`
	// A description of the File.
	Description string `json:"description,required,nullable"`
	// Whether the File was generated by Increase or by you and sent to Increase.
	Direction FileDirection `json:"direction,required"`
	// The filename that was provided upon upload or generated by Increase.
	Filename string `json:"filename,required,nullable"`
	// A URL from where the File can be downloaded at this point in time. The location
	// of this URL may change over time.
	DownloadURL string `json:"download_url,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `file`.
	Type FileType `json:"type,required"`
	JSON FileJSON
}

type FileJSON struct {
	CreatedAt   apijson.Metadata
	ID          apijson.Metadata
	Purpose     apijson.Metadata
	Description apijson.Metadata
	Direction   apijson.Metadata
	Filename    apijson.Metadata
	DownloadURL apijson.Metadata
	Type        apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into File using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilePurpose string

const (
	FilePurposeCheckImageFront            FilePurpose = "check_image_front"
	FilePurposeCheckImageBack             FilePurpose = "check_image_back"
	FilePurposeForm1099Int                FilePurpose = "form_1099_int"
	FilePurposeFormSS4                    FilePurpose = "form_ss_4"
	FilePurposeIdentityDocument           FilePurpose = "identity_document"
	FilePurposeIncreaseStatement          FilePurpose = "increase_statement"
	FilePurposeOther                      FilePurpose = "other"
	FilePurposeTrustFormationDocument     FilePurpose = "trust_formation_document"
	FilePurposeDigitalWalletArtwork       FilePurpose = "digital_wallet_artwork"
	FilePurposeDigitalWalletAppIcon       FilePurpose = "digital_wallet_app_icon"
	FilePurposeDocumentRequest            FilePurpose = "document_request"
	FilePurposeEntitySupplementalDocument FilePurpose = "entity_supplemental_document"
	FilePurposeExport                     FilePurpose = "export"
)

type FileDirection string

const (
	FileDirectionToIncrease   FileDirection = "to_increase"
	FileDirectionFromIncrease FileDirection = "from_increase"
)

type FileType string

const (
	FileTypeFile FileType = "file"
)

type FileNewParams struct {
	// The file contents. This should follow the specifications of
	// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
	// transfers for the multipart/form-data protocol.
	File field.Field[io.Reader] `form:"file,required" format:"binary"`
	// The description you choose to give the File.
	Description field.Field[string] `form:"description"`
	// What the File will be used for in Increase's systems.
	Purpose field.Field[FileNewParamsPurpose] `form:"purpose,required"`
}

func (r FileNewParams) MarshalMultipart() (data []byte, err error) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	{
		name := "anonymous_file"
		if nameable, ok := r.File.Value.(interface{ Name() string }); ok {
			name = nameable.Name()
		}
		part, err := writer.CreateFormFile("file", name)
		if err != nil {
			return nil, err
		}
		io.Copy(part, r.File.Value)
	}
	{
		bdy, err := apijson.Marshal(r.Description)
		if err != nil {
			return nil, err
		}
		writer.WriteField("description", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Purpose)
		if err != nil {
			return nil, err
		}
		writer.WriteField("purpose", string(bdy))
	}
	return body.Bytes(), nil
}

type FileNewParamsPurpose string

const (
	FileNewParamsPurposeCheckImageFront            FileNewParamsPurpose = "check_image_front"
	FileNewParamsPurposeCheckImageBack             FileNewParamsPurpose = "check_image_back"
	FileNewParamsPurposeFormSS4                    FileNewParamsPurpose = "form_ss_4"
	FileNewParamsPurposeIdentityDocument           FileNewParamsPurpose = "identity_document"
	FileNewParamsPurposeOther                      FileNewParamsPurpose = "other"
	FileNewParamsPurposeTrustFormationDocument     FileNewParamsPurpose = "trust_formation_document"
	FileNewParamsPurposeDigitalWalletArtwork       FileNewParamsPurpose = "digital_wallet_artwork"
	FileNewParamsPurposeDigitalWalletAppIcon       FileNewParamsPurpose = "digital_wallet_app_icon"
	FileNewParamsPurposeDocumentRequest            FileNewParamsPurpose = "document_request"
	FileNewParamsPurposeEntitySupplementalDocument FileNewParamsPurpose = "entity_supplemental_document"
)

type FileListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     field.Field[int64]                   `query:"limit"`
	CreatedAt field.Field[FileListParamsCreatedAt] `query:"created_at"`
	Purpose   field.Field[FileListParamsPurpose]   `query:"purpose"`
}

// URLQuery serializes FileListParams into a url.Values of the query parameters
// associated with this value
func (r FileListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FileListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes FileListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r FileListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FileListParamsPurpose struct {
	// Filter Files for those with the specified purpose or purposes. For GET requests,
	// this should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In field.Field[[]FileListParamsPurposeIn] `query:"in"`
}

// URLQuery serializes FileListParamsPurpose into a url.Values of the query
// parameters associated with this value
func (r FileListParamsPurpose) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type FileListParamsPurposeIn string

const (
	FileListParamsPurposeInCheckImageFront            FileListParamsPurposeIn = "check_image_front"
	FileListParamsPurposeInCheckImageBack             FileListParamsPurposeIn = "check_image_back"
	FileListParamsPurposeInForm1099Int                FileListParamsPurposeIn = "form_1099_int"
	FileListParamsPurposeInFormSS4                    FileListParamsPurposeIn = "form_ss_4"
	FileListParamsPurposeInIdentityDocument           FileListParamsPurposeIn = "identity_document"
	FileListParamsPurposeInIncreaseStatement          FileListParamsPurposeIn = "increase_statement"
	FileListParamsPurposeInOther                      FileListParamsPurposeIn = "other"
	FileListParamsPurposeInTrustFormationDocument     FileListParamsPurposeIn = "trust_formation_document"
	FileListParamsPurposeInDigitalWalletArtwork       FileListParamsPurposeIn = "digital_wallet_artwork"
	FileListParamsPurposeInDigitalWalletAppIcon       FileListParamsPurposeIn = "digital_wallet_app_icon"
	FileListParamsPurposeInDocumentRequest            FileListParamsPurposeIn = "document_request"
	FileListParamsPurposeInEntitySupplementalDocument FileListParamsPurposeIn = "entity_supplemental_document"
	FileListParamsPurposeInExport                     FileListParamsPurposeIn = "export"
)

type FileListResponse struct {
	// The contents of the list.
	Data []File `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       FileListResponseJSON
}

type FileListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into FileListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *FileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}