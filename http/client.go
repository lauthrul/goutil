package http

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"time"
)

type Client struct {
	Proxy   string
	Timeout time.Duration
	cli     *fasthttp.Client
	req     *fasthttp.Request
	resp    *fasthttp.Response
}

func (c *Client) Init() {
	c.cli = &fasthttp.Client{
		ReadTimeout:  c.Timeout,
		WriteTimeout: c.Timeout,
	}
	if c.Proxy != "" {
		c.cli.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(c.Proxy, c.Timeout)
	}
	c.req = fasthttp.AcquireRequest()
	c.resp = fasthttp.AcquireResponse()
}

func (c *Client) Head(url string) (*fasthttp.Response, error) {
	c.req.Reset()
	c.resp.Reset()

	c.req.SetRequestURI(url)
	c.req.Header.SetMethod("HEAD")

	err := c.cli.Do(c.req, c.resp)
	if err != nil {
		return nil, err
	}

	return c.resp, nil
}

func (c *Client) Get(url string) (*fasthttp.Response, error) {
	c.req.Reset()
	c.resp.Reset()

	c.req.SetRequestURI(url)
	c.req.Header.SetMethod("GET")

	err := c.cli.Do(c.req, c.resp)
	if err != nil {
		return nil, err
	}

	return c.resp, nil
}

func (c *Client) Post(url string, body []byte) (*fasthttp.Response, error) {
	c.req.Reset()
	c.resp.Reset()

	c.req.SetRequestURI(url)
	c.req.Header.SetMethod("POST")
	c.req.SetBody(body)

	err := c.cli.Do(c.req, c.resp)
	if err != nil {
		return nil, err
	}

	return c.resp, nil
}
