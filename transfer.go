package Transfer

import(
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/proxy"
)

type TransferClient struct {
	client   http.Client
	transport *http.Transport
}

func NewTransfer(client *http.Client, transport *http.Transport) *TransferClient {
	if transport.Dial == nil {
		transport.Dial = client.Transport.(*http.Transport).Dial
	}
	return &TransferClient{client: *client, transport: transport}
}



func (t *TransferClient) request(requestType string, URL string, json []byte )(*http.Response, error) {

	method := strings.ToUpper(requestType)
	if method != "GET" || method != "POST"{
		return http.Response{}, fmt.Errorf("Incorrect Request type: %v", method)
	}
		// Encode the struct as JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest(requestType, URL, bytes.NewBuffer(jsonData))
		if err != nil {
			
			return http.Response{}, err
		}
		// use the http client to fetch the page
		resp, err2 := t.client.Do(req)
		if err != nil {
			
			return resp, fmt.Errorf("Can't reach page: %v", err2)
		}
		defer resp.Body.Close()
		return resp, fmt.Errorf("Request Successful : %v", err2)
	}
	




func Dialer(proxyAddress string)  TransferClient{
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", proxyAddress, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial

	// create a new Transfer struct
	return Transfer.NewTransfer(httpClient, httpTransport)


}
