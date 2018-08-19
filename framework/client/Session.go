package client

import (
	"github.com/gorilla/sessions"
)

/*
 * 获取会话
 */
func (this *Client) Store(name string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(name))
}
func Store(name string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(name))
}
/*
 * 获取默认会话
 */
func (this *Client) DefaultStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte("something-very-secret"))
}
func DefaultStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte("something-very-secret"))
}