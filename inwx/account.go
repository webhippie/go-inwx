package inwx

import (
	"github.com/webhippie/go-inwx/inwx/schema"
)

// Account represents an account in the INWX API.
type Account struct {
	ID         int
	Customer   int
	Username   string
	Salutation string
	Company    string
	Firstname  string
	Lastname   string
	Street     string
	Zip        string
	City       string
	Country    string
	Phone      string
	Fax        string
	Website    string
	Email      string
}

// AccountClient is a client for the accounts API.
type AccountClient struct {
	client *Client
}

// Info retrieves a account information.
func (c *AccountClient) Info() (*Account, *Response, error) {
	result := schema.Account{}

	resp, err := c.client.Do(
		c.client.NewRequest(
			"account.info",
			nil,
		),
		&result,
	)

	if err != nil {
		return nil, resp, err
	}

	return AccountFromSchema(result), resp, nil
}
