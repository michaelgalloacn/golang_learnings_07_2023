package main

import "fmt"

// commmon implementation for enums

type founder string

const (
	Linus_Torvalds   founder = "Linus Torvalds"
	Clarence_Delany  founder = "Clarence Delany"
	Richard_Fairbank founder = "Richard Fairbank"
	Bill_Gates       founder = "Bill Gates"
)

type fakeType string

func main() {

	foundersToCompanies := make(map[founder]string)

	foundersToCompanies[Linus_Torvalds] = "Linux Foundation"
	foundersToCompanies[Clarence_Delany] = "Accenture"
	foundersToCompanies[Richard_Fairbank] = "Capital One"

	unknownCompany, notInMap := foundersToCompanies[Bill_Gates]

	fmt.Println(unknownCompany, notInMap)

	// Example of a company that IS in the map
	knownCompany, isInMap := foundersToCompanies[Linus_Torvalds]
	fmt.Println(knownCompany, isInMap)

	// Length of a map
	fmt.Println("Number of known companies:\n", len(foundersToCompanies))

	// Delete operation for a hashmap
	delete(foundersToCompanies, Linus_Torvalds)
	exKnownCompany, noLongerInMap := foundersToCompanies[Linus_Torvalds]
	fmt.Println(exKnownCompany, noLongerInMap)

	// Iterating through a map
	for name, company := range foundersToCompanies {
		fmt.Println(name, ":", company)
	}

}
