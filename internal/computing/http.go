package computing

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type HttpClient struct {
	host   string
	header http.Header
}

func NewHttpClient(host string, header http.Header) *HttpClient {
	return &HttpClient{
		host:   host,
		header: header,
	}
}

func (c *HttpClient) Get(api string, queries url.Values, dest any) error {
	if queries != nil {
		api += "?" + queries.Encode()
	}
	return c.Request(http.MethodGet, api, nil, dest, "")
}

func (c *HttpClient) PostForm(api string, data url.Values, dest any) error {
	return c.Request(http.MethodPost, api, strings.NewReader(data.Encode()), dest, "application/x-www-form-urlencoded")
}

func (c *HttpClient) PostJSON(api string, data any, dest any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.Request(http.MethodPost, api, bytes.NewReader(b), dest, "application/json")
}

func (c *HttpClient) Request(method string, api string, body io.Reader, dest any, contentType ...string) (err error) {
	if body != nil {
		rb, _ := io.ReadAll(body)
		body = bytes.NewReader(rb)
	}

	if reflect.ValueOf(dest).Kind() != reflect.Ptr {
		return errors.New("dest is not a pointer")
	}

	url := c.host
	if strings.HasPrefix(api, "/") {
		url += api
	} else {
		url += "/" + api
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	for key := range c.header {
		req.Header.Set(key, c.header.Get(key))
	}
	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType[0])
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	bd, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(bd, dest); err != nil {
		return
	}

	if check, ok := dest.(ResultChecker); ok {
		return check.Check()
	}
	return
}

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: defaultTransportDialContext(&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}),
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

type ResultChecker interface {
	Check() error
}
