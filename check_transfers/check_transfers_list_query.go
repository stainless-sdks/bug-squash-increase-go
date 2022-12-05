package check_transfers

type CheckTransfersListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Check Transfers to those that originated from the specified Account.
	AccountId *string                           `json:"account_id,omitempty"`
	CreatedAt *CheckTransfersCreatedAtListQuery `json:"created_at,omitempty"`
}

// Return the page of entries after this one.
func (r *CheckTransfersListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CheckTransfersListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Transfers to those that originated from the specified Account.
func (r *CheckTransfersListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *CheckTransfersListQuery) GetCreatedAt() (CreatedAt CheckTransfersCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}
