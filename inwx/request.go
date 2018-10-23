package inwx

// Args defines a general string map for arguments.
type Args map[string]interface{}

// Request defines a simple request for the XMLRPC API.
type Request struct {
	Method string
	Args   Args
}
