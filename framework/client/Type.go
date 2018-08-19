package client

import (
	"net/http"
	"net"
	."../values"
	."../safety"
)

type Client struct {
	req *http.Request
}
/*
 * 获取客户端
 */
func GetClient(req *http.Request) *Client {
	return &Client{req: req}
}
/*
 * 获取ip地址
 */
func (this *Client) IPaddress() string {
	ip, _, err := net.SplitHostPort(this.req.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}
func GetIPaddress(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}
/*
 * 暴力请求
 */
func (this *Client) GetBruteRequest(name string) *BruteRequest {
	return GetBruteRequest(name, this.req)
}
/*
 * Get参数
 */
func (this *Client) GetValue() RequestValues {
	return ParseGet(this.req)
}
/*
 * Post参数
 */
func (this *Client) PostValue() RequestValues {
	return ParseGet(this.req)
}
