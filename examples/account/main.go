package main

import (
	"fmt"
	"os"

	"github.com/webhippie/go-inwx/inwx"
)

func main() {
	client, err := inwx.NewClient(
		inwx.WithUsername("username"),
		inwx.WithPassword("password"),
	)

	if err != nil {
		fmt.Println("Failed to login user:", err)
		os.Exit(1)
	}

	account, _, err := client.Account.Info()

	if err != nil {
		fmt.Println("Failed to get account:", err)
		os.Exit(1)
	}

	fmt.Printf("Username: %s\n", account.Username)
	fmt.Printf("Salutation: %s\n", account.Salutation)
	fmt.Printf("Firstname: %s\n", account.Firstname)
	fmt.Printf("Lastname: %s\n", account.Lastname)
	fmt.Printf("Company: %s\n", account.Company)
}
