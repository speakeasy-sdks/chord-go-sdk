package operations

type FindOrCreateReferralIdentifierRequestBody struct {
	Email string `json:"email"`
}

type FindOrCreateReferralIdentifierRequest struct {
	Request *FindOrCreateReferralIdentifierRequestBody `request:"mediaType=application/json"`
}

type FindOrCreateReferralIdentifierResponse struct {
	ContentType string
	StatusCode  int
}
