package inwx

// import (
// 	"context"
// 	"strconv"
// 	"errors"
// 	// "encoding/json"
// 	// "bytes"
// 	"fmt"

// 	"github.com/webhippie/go-inwx/inwx/schema"
// )

// Record represents a record in the INWX API.
type Record struct {
	ID       int
	Domain   string
	Name     string
	Type     string
	Value    string
	TTL      int
	Priority int
}

// RecordClient is a client for the records API.
type RecordClient struct {
	client *Client
}

// GetByID retrieves a record by its ID.
// func (c *RecordClient) GetByID(ctx context.Context, id int) (*Record, *Response, error) {
// req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/records/%d", id), nil)

// if err != nil {
// 	return nil, nil, err
// }

// var body schema.RecordGetResponse
// resp, err := c.client.Do(req, &body)

// if err != nil {
// 	if IsError(err, ErrorCodeNotFound) {
// 		return nil, resp, nil
// 	}

// 	return nil, nil, err
// }

// returnRecordFromSchema(body.Record), resp, nil

// 	return &Record{}, nil, nil
// }

// GetByName retreives a record by its name.
// func (c *RecordClient) GetByName(ctx context.Context, name string) (*Record, *Response, error) {
// path := "/servers?name=" + url.QueryEscape(name)
// req, err := c.client.NewRequest(ctx, "GET", path, nil)

// if err != nil {
// 	return nil, nil, err
// }

// var body schema.RecordListResponse
// resp, err := c.client.Do(req, &body)

// if err != nil {
// 	return nil, nil, err
// }

// if len(body.Records) == 0 {
// 	return nil, resp, nil
// }

// return RecordFromSchema(body.Records[0]), resp, nil

// 	return &Record{}, nil, nil
// }

// Get retrieves a record by its ID if the input can be parsed as an integer, otherwise it retrieves a record by its name.
// func (c *RecordClient) Get(ctx context.Context, idOrName string) (*Record, *Response, error) {
// 	if id, err := strconv.Atoi(idOrName); err == nil {
// 		return c.GetByID(ctx, int(id))
// 	}

// 	return c.GetByName(ctx, idOrName)
// }

// RecordListOpts specifies options for listing records.
// type RecordListOpts struct {
// 	ListOpts
// }

// List returns a list of records for a specific page.
// func (c *RecordClient) List(ctx context.Context, opts RecordListOpts) ([]*Record, *Response, error) {
// 	path := "/servers?" + valuesForListOpts(opts.ListOpts).Encode()
// 	req, err := c.client.NewRequest(ctx, "GET", path, nil)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	var body schema.RecordListResponse
// 	resp, err := c.client.Do(req, &body)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	records := make([]*Record, 0, len(body.Records))

// 	for _, r := range body.Records {
// 		records = append(records, RecordFromSchema(r))
// 	}

// 	return records, resp, nil
// }

// All returns all records.
// func (c *RecordClient) All(ctx context.Context) ([]*Record, error) {
// 	return c.AllWithOpts(ctx, RecordListOpts{ListOpts{PerPage: 50}})
// }

// AllWithOpts returns all records for the given options.
// func (c *RecordClient) AllWithOpts(ctx context.Context, opts RecordListOpts) ([]*Record, error) {
// 	all := []*Record{}

// 	_, err := c.client.all(func(page int) (*Response, error) {
// 		opts.Page = page
// 		records, resp, err := c.List(ctx, opts)

// 		if err != nil {
// 			return resp, err
// 		}

// 		all = append(all, records...)
// 		return resp, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return all, nil
// }

// RecordCreateOpts specifies options for creating a new record.
// type RecordCreateOpts struct {
// 	Domain   string
// 	Name     string
// 	Type     string
// 	Value    string
// 	TTL      int
// 	Priority int
// }

// // Validate checks if options are valid.
// func (o RecordCreateOpts) Validate() error {
// 	if o.Domain == "" {
// 		return errors.New("missing domain")
// 	}

// 	if o.Type == "" {
// 		return errors.New("missing type")
// 	}

// 	if o.Value == "" {
// 		return errors.New("missing value")
// 	}

// 	return nil
// }

// // RecordCreateResult is the result of a create record call.
// type RecordCreateResult struct {
// 	Record       *Record
// }

// // Create creates a new record.
// func (c *RecordClient) Create(ctx context.Context, opts RecordCreateOpts) (RecordCreateResult, *Response, error) {
// 	if err := opts.Validate(); err != nil {
// 		return RecordCreateResult{}, nil, err
// 	}

// 	var reqBody schema.RecordCreateRequest
// 	reqBody.Name = opts.Name

// 	// TODO: maybe set defaults?

// 	// reqBodyData, err := json.Marshal(reqBody)

// 	// if err != nil {
// 	// 	return RecordCreateResult{}, nil, err
// 	// }

// 	req, err := c.client.NewRequest(ctx, "", nil) //, bytes.NewReader(reqBodyData))

// 	if err != nil {
// 		return RecordCreateResult{}, nil, err
// 	}

// 	var respBody schema.RecordCreateResponse
// 	resp, err := c.client.Do(req, &respBody)

// 	if err != nil {
// 		return RecordCreateResult{}, resp, err
// 	}

// 	result := RecordCreateResult{
// 		Record: RecordFromSchema(respBody.Record),
// 	}

// 	return result, resp, nil
// }

// // RecordUpdateOpts specifies options for updating a record.
// type RecordUpdateOpts struct {
// 	Name     string
// 	Type     string
// 	Value    string
// 	TTL      int
// 	Priority int
// }

// // Update updates a record.
// func (c *RecordClient) Update(ctx context.Context, server *Record, opts RecordUpdateOpts) (*Record, *Response, error) {
// 	reqBody := schema.RecordUpdateRequest{
// 		Name: opts.Name,
// 	}

// 	if opts.Type != "" {
// 		reqBody.Type = opts.Type
// 	}

// 	if opts.Value != "" {
// 		reqBody.Value = opts.Value
// 	}

// 	if opts.TTL != 0 {
// 		reqBody.TTL = opts.TTL
// 	}

// 	if opts.Priority != 0 {
// 		reqBody.Priority = opts.Priority
// 	}

// 	// reqBodyData, err := json.Marshal(reqBody)

// 	// if err != nil {
// 	// 	return nil, nil, err
// 	// }

// 	path := fmt.Sprintf("/records/%d", server.ID)
// 	req, err := c.client.NewRequest(ctx, path, nil) //, bytes.NewReader(reqBodyData))

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	respBody := schema.RecordUpdateResponse{}
// 	resp, err := c.client.Do(req, &respBody)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return RecordFromSchema(respBody.Record), resp, nil
// }

// // Delete deletes a record.
// func (c *RecordClient) Delete(ctx context.Context, record *Record) (*Response, error) {
// 	req, err := c.client.NewRequest(ctx, fmt.Sprintf("/records/%d", record.ID), nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return c.client.Do(req, nil)
// }
