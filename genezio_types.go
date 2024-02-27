package genezio_types

// payload for genezio http methods
type GenezioHttpRequest struct {
	Headers map[string]string `json:"headers"`
	Http    struct {
		Method    string `json:"method"`
		Path      string `json:"path"`
		Protocol  string `json:"protocol"`
		UserAgent string `json:"userAgent"`
		SourceIp  string `json:"sourceIp"`
	} `json:"http"`
	QueryStringParameters *map[string]string `json:"queryStringParameters"`
	TimeEpoch             int64              `json:"timeEpoch"`
	RawBody               string             `json:"rawBody"`
	Body                  interface{}        `json:"body"`
}

// return type for genezio http methods
type GenezioHttpResponse struct {
	Body            interface{}        `json:"body"`
	StatusCode      string             `json:"statusCode"`
	Headers         *map[string]string `json:"headers"`
	IsBase64Encoded *bool              `json:"isBase64Encoded"`
}
