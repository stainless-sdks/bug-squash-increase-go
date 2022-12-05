package entities

type EntitiesCorporationBeneficialOwnersIndividualIdentificationCreateParams struct {
	// A method that can be used to verify the individual's identity.
	Method *string `json:"method,omitempty"`

	// An identification number that can be used to verify the individual's identity,
	// such as a social security number.
	Number *string `json:"number,omitempty"`

	// Information about the passport used for identification. Required if `method` is
	// equal to `passport`.
	Passport *EntitiesCorporationBeneficialOwnersIndividualIdentificationPassportCreateParams `json:"passport,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *EntitiesCorporationBeneficialOwnersIndividualIdentificationCreateParams) GetMethod() (Method string) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// An identification number that can be used to verify the individual's identity,
// such as a social security number.
func (r *EntitiesCorporationBeneficialOwnersIndividualIdentificationCreateParams) GetNumber() (Number string) {
	if r != nil && r.Number != nil {
		Number = *r.Number
	}
	return
}

// Information about the passport used for identification. Required if `method` is
// equal to `passport`.
func (r *EntitiesCorporationBeneficialOwnersIndividualIdentificationCreateParams) GetPassport() (Passport EntitiesCorporationBeneficialOwnersIndividualIdentificationPassportCreateParams) {
	if r != nil && r.Passport != nil {
		Passport = *r.Passport
	}
	return
}
