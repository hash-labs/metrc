package metrc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// It's fine to store the below variables in plaintext, because these are strictly for the testing sandbox.
const vendorKey string = "xT9DuxBkwB0U8VbomUp28JRKEnsvh3BwVdvHP0PBNATnMWk-"
const metrcUrl string = "https://sandbox-api-ca.metrc.com"
const userKey string = "FusVbe4Yv6W1DGNuxKNhByXU6RO6jSUPcbRCoRDD98VNXc4D"

// ClientInterface specifies an interface for a generic client.
type ClientInterface interface {
	Get(endpoint string) ([]byte, error)
	Post(endpoint string, body []byte) ([]byte, error)
	Delete(endpoint string) ([]byte, error)
	Put(endpoint string, body []byte) ([]byte, error)
}

// HttpClient is a convenience client for raw HTTP calls.
// Wraps a generic `http.Client` and implements `ClientInterface`.
type HttpClient struct {
	Client    *http.Client
	VendorKey string
	UserKey   string
}

// MakeHttpClient lets external packages generate a wrapped HTTP client.
// This exposes the internals used to call Metrc.
func MakeHttpClient(vendorKey string, userKey string) *HttpClient {
	return &HttpClient{
		Client:    &http.Client{},
		VendorKey: vendorKey,
		UserKey:   userKey,
	}
}

// Get executes `GET` requests to the specified endpoint.
func (c *HttpClient) Get(endpoint string) ([]byte, error) {
	endpointUrl := fmt.Sprintf("%s/%s", metrcUrl, endpoint)
	req, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("could not make new request: %s", err)
	}
	return c.do(req)
}

// Post executes `POST` requests to the specified endpoint and body.
func (c *HttpClient) Post(endpoint string, body []byte) ([]byte, error) {
	endpointUrl := fmt.Sprintf("%s/%s", metrcUrl, endpoint)
	req, err := http.NewRequest("POST", endpointUrl, bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, fmt.Errorf("could not form post request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

// Delete executes `DELETE` requests to the specified endpoint and body.
func (c *HttpClient) Delete(endpoint string) ([]byte, error) {
	endpointUrl := fmt.Sprintf("%s/%s", metrcUrl, endpoint)
	req, err := http.NewRequest("DELETE", endpointUrl, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("could not make delete request: %s", err)
	}

	return c.do(req)
}

// Put executes `PUT` requests to the specified endpoint and body.
func (c *HttpClient) Put(endpoint string, body []byte) ([]byte, error) {
	endpointUrl := fmt.Sprintf("%s/%s", metrcUrl, endpoint)
	req, err := http.NewRequest("PUT", endpointUrl, bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, fmt.Errorf("could not form put request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

// do is a boilerplate funtion that executes a request and returns a response.
func (c *HttpClient) do(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.VendorKey, c.UserKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("could not do request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not read from response: %s", err)
	}

	respCode := resp.StatusCode
	if respCode != 200 {
		return []byte{}, fmt.Errorf("response failed with code %d and body %s", respCode, string(body))
	}
	return body, nil
}
