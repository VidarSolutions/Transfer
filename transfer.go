package Transfer

import(
	"fmt"
	"net/http"
	"strings"
)

type Transfer struct {
	client   http.Client
	transport *http.Transport
}

func NewTransfer(client *http.Client, transport *http.Transport) *Transfer {
	if transport.Dial == nil {
		transport.Dial = client.Transport.(*http.Transport).Dial
	}
	return &Transfer{client: *client, transport: transport}
}



func (t *Transfer) request(type string, URL string, json []byte )(http.Response, error) {

	method = strings.ToUpper(type)
	if method != "GET" || method != "POST{
		return http.Response{}, fmt.Errorf("Incorrect Request type: %v", method)
	}
		req, err := http.NewRequest(type, URL, json)
		if err != nil {
			
			return http.Response{}, err
		}
		// use the http client to fetch the page
		resp, err2 := t.client.Do(req)
		if err != nil {
			
			return resp, fmt.Errorf("Can't reach page: %v", err2)
		}
		defer resp.Body.Close()
		
	}
	return resp, fmt.Errorf("Request Successful : %v", err2)
}



func Dialer(proxyAddress string)  Transfer{
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
