package operations

type CreateEndpointRequestBodyEndpoint struct {
	Events    []string `json:"events"`
	Filters   []string `json:"filters,omitempty"`
	TargetURL string   `json:"target_url"`
	Token     *string  `json:"token,omitempty"`
}

type CreateEndpointRequestBody struct {
	Endpoint CreateEndpointRequestBodyEndpoint `json:"endpoint"`
}

type CreateEndpointRequest struct {
	Request *CreateEndpointRequestBody `request:"mediaType=application/json"`
}

type CreateEndpointResponse struct {
	ContentType string
	StatusCode  int
}
