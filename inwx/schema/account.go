package schema

// Account defines the schema of an account.
type Account struct {
	ID         int    `mapstructure:"accountId"`
	Customer   int    `mapstructure:"customerId"`
	Username   string `mapstructure:"username"`
	Salutation string `mapstructure:"title"`
	Company    string `mapstructure:"org"`
	Firstname  string `mapstructure:"firstname"`
	Lastname   string `mapstructure:"lastname"`
	Street     string `mapstructure:"street"`
	Zip        string `mapstructure:"pc"`
	City       string `mapstructure:"city"`
	Country    string `mapstructure:"cc"`
	Phone      string `mapstructure:"voice"`
	Fax        string `mapstructure:"fax"`
	Website    string `mapstructure:"www"`
	Email      string `mapstructure:"email"`
}
