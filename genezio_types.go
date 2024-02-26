package genezio_types

// payload for genezio http methods
type GenezioHttpRequest struct {
	Headers map[string]string
	Http    struct {
		Method    string
		Path      string
		Protocol  string
		UserAgent string
		SourceIp  string
	}
	QueryStringParameters *map[string]string
	TimeEpoch             int64
	RawBody               string
	Body                  interface{}
}

// return type for genezio http methods
type GenezioHttpResponse struct {
	Body            interface{}
	StatusCode      string
	Headers         *map[string]string
	IsBase64Encoded *bool
}
