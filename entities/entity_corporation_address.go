package entities

type EntityCorporationAddress struct {
	// The first line of the address.
	Line1 *string `json:"line1,omitempty"`

	// The second line of the address.
	Line2 *string `json:"line2,omitempty"`

	// The city of the address.
	City *string `json:"city,omitempty"`

	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State *string `json:"state,omitempty"`

	// The ZIP code of the address.
	Zip *string `json:"zip,omitempty"`
}

// The first line of the address.
func (r *EntityCorporationAddress) GetLine1() (Line1 string) {
	if r != nil && r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r *EntityCorporationAddress) GetLine2() (Line2 string) {
	if r != nil && r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r *EntityCorporationAddress) GetCity() (City string) {
	if r != nil && r.City != nil {
		City = *r.City
	}
	return
}

// The two-letter United States Postal Service (USPS) abbreviation for the state of
// the address.
func (r *EntityCorporationAddress) GetState() (State string) {
	if r != nil && r.State != nil {
		State = *r.State
	}
	return
}

// The ZIP code of the address.
func (r *EntityCorporationAddress) GetZip() (Zip string) {
	if r != nil && r.Zip != nil {
		Zip = *r.Zip
	}
	return
}
