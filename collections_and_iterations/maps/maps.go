package main

import "fmt"

// commmon implementation for enums is to declare a set of constants in the same block, often with a custom type

type founder string

const (
	Linus_Torvalds   founder = "Linus Torvalds"
	Clarence_Delany  founder = "Clarence Delany"
	Richard_Fairbank founder = "Richard Fairbank"
	Bill_Gates       founder = "Bill Gates"
)

type fakeType string

func main() {

	// Syntax to create a map is make(map[keyType]valType)
	foundersToCompanies := make(map[founder]string)

	foundersToCompanies[Linus_Torvalds] = "Linux Foundation"
	foundersToCompanies[Clarence_Delany] = "Accenture"
	foundersToCompanies[Richard_Fairbank] = "Capital One"

	// Golang supports multiple returns
	// Accessing a value from a map returns two values
	// First value is zero value for return type if key is not found
	// Second value is bool representing whether the key is in the map
	unknownCompany, notInMap := foundersToCompanies[Bill_Gates]
	fmt.Println(unknownCompany, notInMap)

	// Example of a company that IS in the map
	knownCompany, isInMap := foundersToCompanies[Linus_Torvalds]
	fmt.Println(knownCompany, isInMap)

	// Length of a map is found with built in len function
	fmt.Println("Number of known companies:\n", len(foundersToCompanies))

	// Deletion operations use built in delete function with map as its first argument and key as its second
	delete(foundersToCompanies, Linus_Torvalds)
	exKnownCompany, noLongerInMap := foundersToCompanies[Linus_Torvalds]
	fmt.Println(exKnownCompany, noLongerInMap)

	// Iterating through the range operator returns two values, key and value
	for name, company := range foundersToCompanies {
		fmt.Println(name, ":", company)
	}

}
