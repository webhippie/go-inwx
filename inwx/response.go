package inwx

// Response represents a response from the API.
type Response struct {
	Status  int                    `xmlrpc:"code"`
	Message string                 `xmlrpc:"msg"`
	Code    string                 `xmlrpc:"reasonCode"`
	Reason  string                 `xmlrpc:"reason"`
	Data    map[string]interface{} `xmlrpc:"resData"`
}
