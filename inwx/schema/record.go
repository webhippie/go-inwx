package schema

// Record defines the schema of a record.
type Record struct {
	ID int `json:"id"`
}

// RecordCreateRequest defines the schema for the request to create a record.
type RecordCreateRequest struct {
	Name string `json:"name"`
}

// RecordCreateResponse defines the schema of the response when creating a record.
type RecordCreateResponse struct {
	Record Record `json:"record"`
}

// RecordUpdateRequest defines the schema of the request to update a record.
type RecordUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Value    string `json:"value,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

// RecordUpdateResponse defines the schema of the response when updating a record.
type RecordUpdateResponse struct {
	Record Record `json:"record"`
}
