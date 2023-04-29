package Transfer

import(
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/proxy"
)

type TransferClient struct {
	Client   http.Client
	Transport *http.Transport
}

func NewTransfer(client *http.Client, transport *http.Transport) *TransferClient {
	if transport.Dial == nil {
		transport.Dial = client.Transport.(*http.Transport).Dial
	}
	return &TransferClient{Client: *client, Transport: transport}
}



func (t *TransferClient) Request(requestType string, URL string, data []byte )(*http.Response, error) {

	method := strings.ToUpper(requestType)
	if method != "GET" && method != "POST"{
		return nil, fmt.Errorf("Incorrect Request type: %v", method)
	}
	if !strings.HasPrefix(URL, "http://") {
		URL = "http://" + URL
	}
		// Encode the struct as JSON
		//jsonData, err := json.Marshal(data)
		//if err != nil {
		//	panic(err)
		//} //callers of this function will have to marshal the data before sending it here
		req, err := http.NewRequest(requestType, URL, bytes.NewBuffer(data))
		if err != nil {
			
			return nil, err
		}
		// use the http client to fetch the page
		resp, err2 := t.Client.Do(req)
		if err2 != nil {
			
			return resp, fmt.Errorf("Can't reach page: %v", err2)
		}
		
		return resp, nil
	}
	




func Dialer(proxyAddress string)  *TransferClient{
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
	return NewTransfer(httpClient, httpTransport)


}
