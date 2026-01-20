package creationalPatterns

import (
	"fmt"
	// "sync"
)

// Singleton Pattern

// Ensure a class has only one instance and provide a global point of access to it.

// Key Go Differences from Typescript:
// - No private constructors, so we use unexported struct + exported getter
// - sync.Once ensures thread-safe lazy initialisation
// - Package-level variable holds the singleton instance

// Database connection holds the connection details
type DatabaseConnection struct {
	url string
	api string
}

// Database is our singleton
type Database struct {
	connection *DatabaseConnection
}

// Package level variables for singleton
var (
	instance *Database
	once     Sync.Once
)

// Get instance returns the singleton Database instance
// sync.Once guarantees this runs exactly once, even with concurrent calls
func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Creating singleton instance")
		instance = &Database{
			connection: &DatabaseConnection{
				url: "someUrl",
				api: "SomeApiKey",
			},
		}
	})
	return instance
}

// GetConnectionInfo returns connection details (for demonstration)
func (d *Database) GetConnectionInfo() string {
	return fmt.Sprintf("URL: %s, API: %s", d.connection.url, d.connection.api)
}

// Client code example
func ExampleSingleton() {
	// Both Calls return the same instance
	db1 := GetInstance()
	db2 := GetInstance()

	fmt.Printf("db1 == db2 %v\n", db1 == db2) // true
	fmt.Println(db1.GetConenctionInfo())
}
