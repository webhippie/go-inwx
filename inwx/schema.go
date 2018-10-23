package inwx

import (
	"github.com/webhippie/go-inwx/inwx/schema"
)

// AccountFromSchema converts a schema.Account to an Account.
func AccountFromSchema(s schema.Account) *Account {
	return &Account{
		ID:         s.ID,
		Customer:   s.Customer,
		Username:   s.Username,
		Salutation: s.Salutation,
		Company:    s.Company,
		Firstname:  s.Firstname,
		Lastname:   s.Lastname,
		Street:     s.Street,
		Zip:        s.Zip,
		City:       s.City,
		Country:    s.Country,
		Phone:      s.Phone,
		Fax:        s.Fax,
		Website:    s.Website,
		Email:      s.Email,
	}
}

// // RecordFromSchema converts a schema.Record to a Record.
// func RecordFromSchema(s schema.Record) *Record {
//   record := &Record{
//     ID: s.ID,
//   }

//   return record
// }
