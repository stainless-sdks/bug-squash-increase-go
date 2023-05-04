package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// ACHPrenotificationService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewACHPrenotificationService] method
// instead.
type ACHPrenotificationService struct {
	Options []option.RequestOption
}

// NewACHPrenotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewACHPrenotificationService(opts ...option.RequestOption) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Options = opts
	return
}

// Create an ACH Prenotification
func (r *ACHPrenotificationService) New(ctx context.Context, body ACHPrenotificationNewParams, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_prenotifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, ach_prenotification_id string, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_prenotifications/%s", ach_prenotification_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) (res *shared.Page[ACHPrenotification], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_prenotifications"
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

// List ACH Prenotifications
func (r *ACHPrenotificationService) ListAutoPaging(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) *shared.PageAutoPager[ACHPrenotification] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// ACH Prenotifications are one way you can verify account and routing numbers by
// Automated Clearing House (ACH).
type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID string `json:"id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// Additional information for the recipient.
	Addendum string `json:"addendum,required,nullable"`
	// The description of the date of the notification.
	CompanyDescriptiveDate string `json:"company_descriptive_date,required,nullable"`
	// Optional data associated with the notification.
	CompanyDiscretionaryData string `json:"company_discretionary_data,required,nullable"`
	// The description of the notification.
	CompanyEntryDescription string `json:"company_entry_description,required,nullable"`
	// The name by which you know the company.
	CompanyName string `json:"company_name,required,nullable"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator ACHPrenotificationCreditDebitIndicator `json:"credit_debit_indicator,required,nullable"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date-time"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn ACHPrenotificationPrenotificationReturn `json:"prenotification_return,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The lifecycle status of the ACH Prenotification.
	Status ACHPrenotificationStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type ACHPrenotificationType `json:"type,required"`
	JSON achPrenotificationJSON
}

// achPrenotificationJSON contains the JSON metadata for the struct
// [ACHPrenotification]
type achPrenotificationJSON struct {
	ID                       apijson.Field
	AccountNumber            apijson.Field
	Addendum                 apijson.Field
	CompanyDescriptiveDate   apijson.Field
	CompanyDiscretionaryData apijson.Field
	CompanyEntryDescription  apijson.Field
	CompanyName              apijson.Field
	CreditDebitIndicator     apijson.Field
	EffectiveDate            apijson.Field
	RoutingNumber            apijson.Field
	PrenotificationReturn    apijson.Field
	CreatedAt                apijson.Field
	Status                   apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ACHPrenotification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHPrenotificationCreditDebitIndicator string

const (
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	ACHPrenotificationCreditDebitIndicatorDebit  ACHPrenotificationCreditDebitIndicator = "debit"
)

// If your prenotification is returned, this will contain details of the return.
type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the Prenotification was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             achPrenotificationPrenotificationReturnJSON
}

// achPrenotificationPrenotificationReturnJSON contains the JSON metadata for the
// struct [ACHPrenotificationPrenotificationReturn]
type achPrenotificationPrenotificationReturnJSON struct {
	CreatedAt        apijson.Field
	ReturnReasonCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ACHPrenotificationPrenotificationReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHPrenotificationStatus string

const (
	ACHPrenotificationStatusPendingSubmitting ACHPrenotificationStatus = "pending_submitting"
	ACHPrenotificationStatusRequiresAttention ACHPrenotificationStatus = "requires_attention"
	ACHPrenotificationStatusReturned          ACHPrenotificationStatus = "returned"
	ACHPrenotificationStatusSubmitted         ACHPrenotificationStatus = "submitted"
)

type ACHPrenotificationType string

const (
	ACHPrenotificationTypeACHPrenotification ACHPrenotificationType = "ach_prenotification"
)

type ACHPrenotificationNewParams struct {
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum param.Field[string] `json:"addendum"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName param.Field[string] `json:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator param.Field[ACHPrenotificationNewParamsCreditDebitIndicator] `json:"credit_debit_indicator"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// Your identifer for the transfer recipient.
	IndividualID param.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName param.Field[string] `json:"individual_name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode param.Field[ACHPrenotificationNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
}

func (r ACHPrenotificationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACHPrenotificationNewParamsCreditDebitIndicator string

const (
	ACHPrenotificationNewParamsCreditDebitIndicatorCredit ACHPrenotificationNewParamsCreditDebitIndicator = "credit"
	ACHPrenotificationNewParamsCreditDebitIndicatorDebit  ACHPrenotificationNewParamsCreditDebitIndicator = "debit"
)

type ACHPrenotificationNewParamsStandardEntryClassCode string

const (
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit        ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	ACHPrenotificationNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHPrenotificationNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHPrenotificationNewParamsStandardEntryClassCodeInternetInitiated             ACHPrenotificationNewParamsStandardEntryClassCode = "internet_initiated"
)

type ACHPrenotificationListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     param.Field[int64]                                 `query:"limit"`
	CreatedAt param.Field[ACHPrenotificationListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes [ACHPrenotificationListParams]'s query parameters as
// `url.Values`.
func (r ACHPrenotificationListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ACHPrenotificationListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [ACHPrenotificationListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r ACHPrenotificationListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of ACH Prenotification objects
type ACHPrenotificationListResponse struct {
	// The contents of the list.
	Data []ACHPrenotification `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       achPrenotificationListResponseJSON
}

// achPrenotificationListResponseJSON contains the JSON metadata for the struct
// [ACHPrenotificationListResponse]
type achPrenotificationListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHPrenotificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
