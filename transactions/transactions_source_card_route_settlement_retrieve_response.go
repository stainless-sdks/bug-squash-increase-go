package transactions

type TransactionsSourceCardRouteSettlementRetrieveResponse struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency *string `json:"currency,omitempty"`

	MerchantAcceptorId *string `json:"merchant_acceptor_id,omitempty"`

	MerchantCity *string `json:"merchant_city,omitempty"`

	MerchantCountry *string `json:"merchant_country,omitempty"`

	MerchantDescriptor *string `json:"merchant_descriptor,omitempty"`

	MerchantState *string `json:"merchant_state,omitempty"`

	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`
}

// The settled amount in the minor unit of the settlement currency. For dollars,
// for example, this is cents.
func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *TransactionsSourceCardRouteSettlementRetrieveResponse) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}
