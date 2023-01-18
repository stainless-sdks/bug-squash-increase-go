package wire_transfers

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type WireTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewWireTransferService(requester core.Requester) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type WireTransfer struct {
	// The wire transfer's identifier.
	ID string `json:"id"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient"`
	// The transfer amount in USD cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency WireTransferCurrency `json:"currency"`
	// The destination account number.
	AccountNumber string `json:"account_number"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID *string `json:"external_account_id"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *WireTransferApproval `json:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *WireTransferCancellation `json:"cancellation"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal *WireTransferReversal `json:"reversal"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt string `json:"created_at"`
	// The transfer's network.
	Network WireTransferNetwork `json:"network"`
	// The lifecycle status of the transfer.
	Status WireTransferStatus `json:"status"`
	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission *WireTransferSubmission `json:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `json:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionID *string `json:"transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type WireTransferType `json:"type"`
}

// The identifier of the External Account the transfer was made to, if any.
func (r *WireTransfer) GetExternalAccountID() (ExternalAccountID string) {
	if r != nil && r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *WireTransfer) GetApproval() (Approval WireTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *WireTransfer) GetCancellation() (Cancellation WireTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// If your transfer is reversed, this will contain details of the reversal.
func (r *WireTransfer) GetReversal() (Reversal WireTransferReversal) {
	if r != nil && r.Reversal != nil {
		Reversal = *r.Reversal
	}
	return
}

// After the transfer is submitted to Fedwire, this will contain supplemental
// details.
func (r *WireTransfer) GetSubmission() (Submission WireTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *WireTransfer) GetTemplateID() (TemplateID string) {
	if r != nil && r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *WireTransfer) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

type WireTransferCurrency string

const (
	WireTransferCurrencyCad WireTransferCurrency = "CAD"
	WireTransferCurrencyChf WireTransferCurrency = "CHF"
	WireTransferCurrencyEur WireTransferCurrency = "EUR"
	WireTransferCurrencyGbp WireTransferCurrency = "GBP"
	WireTransferCurrencyJpy WireTransferCurrency = "JPY"
	WireTransferCurrencyUsd WireTransferCurrency = "USD"
)

//
type WireTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt string `json:"approved_at"`
}

//
type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt string `json:"canceled_at"`
}

//
type WireTransferReversal struct {
	// The amount that was reversed.
	Amount int `json:"amount"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate string `json:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate string `json:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation *string `json:"receiver_financial_institution_information"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation *string `json:"financial_institution_to_financial_institution_information"`
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *WireTransferReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *WireTransferReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

type WireTransferNetwork string

const (
	WireTransferNetworkWire WireTransferNetwork = "wire"
)

type WireTransferStatus string

const (
	WireTransferStatusCanceled          WireTransferStatus = "canceled"
	WireTransferStatusRequiresAttention WireTransferStatus = "requires_attention"
	WireTransferStatusPendingApproval   WireTransferStatus = "pending_approval"
	WireTransferStatusRejected          WireTransferStatus = "rejected"
	WireTransferStatusReversed          WireTransferStatus = "reversed"
	WireTransferStatusComplete          WireTransferStatus = "complete"
	WireTransferStatusPendingCreating   WireTransferStatus = "pending_creating"
)

//
type WireTransferSubmission struct {
	// The accountability data for the submission.
	InputMessageAccountabilityData string `json:"input_message_accountability_data"`
}

type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

type CreateAWireTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID string `json:"account_id"`
	// The account number for the destination account.
	AccountNumber string `json:"account_number,omitempty"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number,omitempty"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID string `json:"external_account_id,omitempty"`
	// The transfer amount in cents.
	Amount int `json:"amount"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient"`
	// The beneficiary's name.
	BeneficiaryName string `json:"beneficiary_name,omitempty"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,omitempty"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,omitempty"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,omitempty"`
}

type ListWireTransfersQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID string `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID string                          `query:"external_account_id"`
	CreatedAt         ListWireTransfersQueryCreatedAt `query:"created_at"`
}

type ListWireTransfersQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore string `json:"on_or_before,omitempty"`
}

//
type WireTransferList struct {
	// The contents of the list.
	Data []WireTransfer `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *WireTransferList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *WireTransferService) Create(ctx context.Context, body *CreateAWireTransferParameters, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.post(
		ctx,
		"/wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *WireTransferService) Retrieve(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/wire_transfers/%s", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type WireTransfersPage struct {
	*pagination.Page[WireTransfer]
}

func (r *WireTransfersPage) WireTransfer() *WireTransfer {
	return r.Current()
}

func (r *WireTransfersPage) GetNextPage() (*WireTransfersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &WireTransfersPage{page}, nil
	}
}

func (r *WireTransferService) List(ctx context.Context, query *ListWireTransfersQuery, opts ...*core.RequestOpts) (res *WireTransfersPage, err error) {
	page := &WireTransfersPage{
		Page: &pagination.Page[WireTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/wire_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Simulates the reversal of an Wire Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `complete`.
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/reverse", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// Simulates the submission of a Wire Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *WireTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/submit", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}
