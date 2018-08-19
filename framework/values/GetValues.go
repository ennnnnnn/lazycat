package values

import "net/http"

func ParseGet(req *http.Request) RequestValues {
	query := req.URL.Query()
	return RequestValues{values: &query}
}
