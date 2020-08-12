package httpx

import "time"

type Client struct {
	Timeout time.Duration
}

var DefaultHttpxClient = &Client{}

func Get() {

}
