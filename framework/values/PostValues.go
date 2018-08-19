package values

import "net/http"

func ParsePost(req *http.Request) RequestValues {
	req.ParseForm()
	return RequestValues{values: &req.Form}
}
