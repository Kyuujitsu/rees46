package rees46

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseUrl = "http://api.rees46.com/"
	GET_METHOD     = "GET"
	POST_METHOD    = "POST"
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	ShopID    string
	SessionID string

	Recommendation *RecommendService
}

type Response struct {
	*http.Response
	RawBody string
}

type ErrorResponse struct {
	Response *http.Response
	Status   string `json:"status"`
	Message  string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Status, r.Message)
}

func NewClient(shopId string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL, ShopID: shopId}
	c.Recommendation = &RecommendService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)
	err = checkResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		err = json.Unmarshal([]byte(response.RawBody), &v)
	}

	return response, err
}

func newResponse(r *http.Response) *Response {
	rawBody, _ := ioutil.ReadAll(r.Body)
	response := &Response{Response: r, RawBody: string(rawBody)}
	return response
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
