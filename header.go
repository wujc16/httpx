package httpx

import "net/textproto"

type Header map[string][]string

func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}

func (h Header) Values(key string) []string {
	return textproto.MIMEHeader(h).Values(key)
}
