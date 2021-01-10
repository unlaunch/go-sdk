package util

import (
	"bytes"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPClient struct {
	host       string
	httpClient *http.Client
	headers    map[string]string
	logger     logger.Interface
	sdkKey     string
}

func NewHTTPClient(
	sdkKey string,
	host string,
	timeout int,
	logger logger.Interface,
) *HTTPClient {

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	return &HTTPClient{
		host:       host,
		httpClient: client,
		logger:     logger,
		sdkKey:     sdkKey,
	}
}

func (c *HTTPClient) Get(path string) ([]byte, error) {
	apiEndpoint := c.host + path
	c.logger.Debug("[HTTP GET] ", apiEndpoint)

	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	req.Header.Add("X-Api-Key", c.sdkKey)
	req.Header.Add("Content-Type", "application/json")
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

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return body, nil
	} else {
		c.logger.Error(fmt.Sprintf("[HTTP GET] status code: %d", resp.StatusCode))
		return nil, &dtos.HTTPError{
			Code: resp.StatusCode,
			Msg:  resp.Status,
		}
	}
}

func (c *HTTPClient) Post(path string, body []byte) error {
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
