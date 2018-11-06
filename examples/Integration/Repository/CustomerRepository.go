package testdata

// Customer Object to hold customer data for the BulkFromDataSourceWithMerge example
type Customer struct {
	FirstName     string
	LastName      string
	EmailAddress  string
	FavoriteColor string
}

// GetCustomers returns a list of mock Customers for the BulkFromDataSourceWithMerge example
func GetCustomers() []Customer {
	customer1 := Customer{FirstName: "Recipient", LastName: "One", EmailAddress: "recipient1@example.com", FavoriteColor: "Green"}
	customer2 := Customer{FirstName: "Recipient", LastName: "Two", EmailAddress: "recipient2@example.com", FavoriteColor: "Red"}
	customer3 := Customer{FirstName: "Recipient", LastName: "Three", EmailAddress: "recipient3@example.com", FavoriteColor: "Blue"}
	customer4 := Customer{FirstName: "Recipient", LastName: "Four", EmailAddress: "recipient4@example.com", FavoriteColor: "Orange"}

	var list []Customer
	list = append(list, customer1, customer2, customer3, customer4)

	return list
}
