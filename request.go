package httpx

type Request struct {
	Method        string
	Proto         string // HTTP/1.0
	Header        Header
	ContentLength int64
}
