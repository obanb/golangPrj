package passiveEndpoint

type RegisterEndpointRequest struct {
	EndpointName string                 `json:"endpointName"`
	Description  *string                `json:"description,omitempty"`
	DataPattern  map[string]interface{} `json:"dataPattern"`
}

type RegisterEndpointResponse struct {
	EndpointName string                 `json:"endpointName"`
	Description  *string                `json:"description,omitempty"`
	DataPattern  map[string]interface{} `json:"dataPattern"`
}
