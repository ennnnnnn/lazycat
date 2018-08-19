package safety

import (
	"time"
	"net/http"
	."../client"
)

type BruteRequest struct {
	req *http.Request
	name string
	n int
	time int64
}

var brs = make(map[string]BruteRequest)

func GetBruteRequest(name string, req *http.Request) *BruteRequest {
	cip := GetIPaddress(req) + "#" + name
	if v, ok := brs[cip]; ok {
		return &v
	}
	return &BruteRequest{name: cip, n: 0, time: time.Now().Unix()}
}

func ClearBruteRequest() {
	for k, v := range brs {
		if time.Now().Unix()-v.time >= 300 {
			delete(brs, k)
		}
	}
}

func (this *BruteRequest)Good() bool {
	return this.n < 3
}

func (this *BruteRequest)Update() {
	if v, ok := brs[this.name]; ok {
		v.n++
	} else {
		brs[this.name] = *this
	}
}