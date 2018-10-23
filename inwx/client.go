package inwx

import (
	"github.com/kolo/xmlrpc"
	"github.com/mitchellh/mapstructure"
)

const (
	// Endpoint is the endpoint URL for production usage.
	Endpoint = "https://api.domrobot.com/xmlrpc/"

	// Sandbox is the endpoint URL for development usage.
	Sandbox = "https://api.ote.domrobot.com/xmlrpc/"

	// Language defines the API language for the responses.
	Language = "eng"

	// UserAgent is the value for the User-Agent header sent with each request.
	UserAgent = "go-inwx/" + Version
)

// Client is a client for the INWX API.
type Client struct {
	rpcClient *xmlrpc.Client

	endpoint string
	username string
	password string

	Account AccountClient
	Record  RecordClient
}

// A ClientOption is used to configure a Client.
type ClientOption func(*Client)

// WithUsername configures a Client to use the specified username for authentication.
func WithUsername(value string) ClientOption {
	return func(client *Client) {
		client.username = value
	}
}

// WithPassword configures a Client to use the specified password for authentication.
func WithPassword(value string) ClientOption {
	return func(client *Client) {
		client.password = value
	}
}

// WithEndpoint configures a Client to use another endpoint URL.
func WithEndpoint(value string) ClientOption {
	return func(client *Client) {
		client.endpoint = value
	}
}

// WithXMLRPC configures a Client to use another XMLRPC client.
func WithXMLRPC(value *xmlrpc.Client) ClientOption {
	return func(client *Client) {
		client.rpcClient = value
	}
}

// NewClient creates a new client.
func NewClient(options ...ClientOption) (*Client, error) {
	client := &Client{
		endpoint: Endpoint,
	}

	for _, option := range options {
		option(client)
	}

	if client.rpcClient == nil {
		rpcClient, err := xmlrpc.NewClient(client.endpoint, nil)

		if err != nil {
			return nil, err
		}

		client.rpcClient = rpcClient
	}

	if err := client.Login(); err != nil {
		return nil, err
	}

	client.Account = AccountClient{client: client}
	client.Record = RecordClient{client: client}

	return client, nil
}

// NewRequest creates an XMLRPC request against the API.
func (c *Client) NewRequest(method string, args Args) *Request {
	if args != nil {
		args["lang"] = Language
	}

	return &Request{
		Method: method,
		Args:   args,
	}
}

// Login performs a login to the XMLRPC API.
func (c *Client) Login() error {
	_, err := c.Do(
		c.NewRequest(
			"account.login",
			Args{
				"user": c.username,
				"pass": c.password,
			},
		),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

// Logout performs a logout to the XMLRPC API.
func (c *Client) Logout() error {
	_, err := c.Do(
		c.NewRequest(
			"account.logout",
			nil,
		),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

// Do performs the provided XMLRPC request against the API.
func (c *Client) Do(r *Request, v interface{}) (*Response, error) {
	resp := &Response{}
	err := c.rpcClient.Call(r.Method, r.Args, resp)

	if err != nil {
		return nil, err
	}

	if resp.Status < 1000 || resp.Status > 1500 {
		return nil, &Error{
			Status:  resp.Status,
			Message: resp.Message,
			Reason:  resp.Reason,
			Code:    resp.Code,
		}
	}

	if v != nil {
		if err := mapstructure.Decode(resp.Data, v); err != nil {
			return resp, err
		}
	}

	return resp, nil
}
