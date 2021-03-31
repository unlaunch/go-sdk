package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

type HTTPClient interface {
	Get(path string) ([]byte, error)
	Post(path string, body []byte) error
}

type simpleHTTPClient struct {
	host           string
	httpClient     *http.Client
	headers        map[string]string
	logger         logger.Interface
	sdkKey         string
	lastModifiedAt string
}

type GenericHTTPClient struct {
	host       string
	httpClient *http.Client
	logger     logger.Interface
	sdkKey     string
}

// NewHTTPClient returns a new http client
func NewHTTPClient(
	sdkKey string,
	host string,
	timeout time.Duration,
	logger logger.Interface,
	sync0 bool,
) HTTPClient {

	if sync0 {
		client := &http.Client{
			Timeout: timeout,
		}

		return &GenericHTTPClient{
			host:       host,
			httpClient: client,
			logger:     logger,
			sdkKey:     sdkKey,
		}

	} else {
		client := &http.Client{
			Timeout: timeout,
		}

		return &simpleHTTPClient{
			host:       host,
			httpClient: client,
			logger:     logger,
			sdkKey:     sdkKey,
		}
	}
}

func (c *simpleHTTPClient) Get(path string) ([]byte, error) {
	apiEndpoint := c.host + path
	c.logger.Debug("[HTTP GET] ", apiEndpoint)

	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	req.Header.Add("X-Api-Key", c.sdkKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("If-Modified-Since", c.lastModifiedAt)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		c.logger.Error("[HTTP GET] HTTP error ", err)
		return nil, err
	}

	defer resp.Body.Close()

	reader := resp.Body
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		c.logger.Error("[HTTP GET] error reading body", err)
		return nil, err
	}

	c.logger.Debug(fmt.Sprintf("[HTTP GET] status code: %d", resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		c.lastModifiedAt = resp.Header.Get("Last-Modified")
		return body, nil
	} else if resp.StatusCode == 304 {
		return nil, nil
	} else {
		return nil, &dtos.HTTPError{
			Code: resp.StatusCode,
			Msg:  resp.Status,
		}
	}
}

func (c *simpleHTTPClient) Post(path string, body []byte) error {
	apiEndpoint := c.host + path
	c.logger.Debug("[HTTP POST] ", apiEndpoint)

	req, _ := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(body))
	req.Close = true

	req.Header.Add("X-Api-Key", c.sdkKey)
	req.Header.Add("Content-Type", "application/json")

	c.logger.Debug(fmt.Sprintf("Headers: %v", req.Header))

	c.logger.Trace("REQ_BODY -->", string(body), "<--REQ_BODY")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("[HTTP POST] error: ", req.URL.String(), err.Error())
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}

	c.logger.Trace("RES_BODY -->", string(respBody), "<-- RES_BODY")

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	c.logger.Error(fmt.Sprintf("[HTTP POST] Possible error. Status Code: %d", resp.StatusCode))
	return &dtos.HTTPError{
		Code: resp.StatusCode,
		Msg:  resp.Status,
	}
}

func (c *GenericHTTPClient) Get(path string) ([]byte, error) {
	apiEndpoint := path
	c.logger.Debug("[HTTP GET] ", apiEndpoint)

	req, _ := http.NewRequest("GET", apiEndpoint, nil)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		c.logger.Error("[HTTP GET] HTTP error ", err)
		return nil, err
	}

	defer resp.Body.Close()

	reader := resp.Body
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		c.logger.Error("[HTTP GET] error reading body", err)
		return nil, err
	}

	c.logger.Debug(fmt.Sprintf("[HTTP GET] status code: %d", resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return body, nil
	} else {
		return nil, &dtos.HTTPError{
			Code: resp.StatusCode,
			Msg:  resp.Status,
		}
	}
}

func (c *GenericHTTPClient) Post(path string, body []byte) error {
	return nil
}
