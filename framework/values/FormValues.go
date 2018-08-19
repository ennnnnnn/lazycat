package values

import "net/http"

func ParseForm(req *http.Request) RequestValues {
	req.ParseForm()
	return RequestValues{values: &req.Form}
}
