package main

import "fmt"

func main() {

	// an array in Golang has a predetermined length
	var teamArray [5]string

	teamArray[0] = "Frank"
	teamArray[1] = "Mike"
	teamArray[2] = "Daniel"
	teamArray[3] = "Jacob"
	teamArray[4] = "Patrick"

	// Will cause an out of bounds error
	// teamArray[5] = "Rhett"

	for index := 0; index < len(teamArray); index++ {
		fmt.Println(teamArray[index])
	}

	initalizedTeamArray := [5]string{
		"Frank", "Mike", "Daniel", "Jacob", "Patrick"}

	// For Each/Enumerate equivalent
	for index, value := range initalizedTeamArray {
		fmt.Println(index, value)
	}

	// for each without index
	for _, value := range initalizedTeamArray {
		fmt.Println(value)
	}

	// Making a bigger array
	// var teamWithDeliveryLead [6]string
	// for i, v := range teamArray {
	// 	teamWithDeliveryLead[i] = v
	// }
	// teamWithDeliveryLead[5] = "Rhett"
	// fmt.Println("Team with DL:\n", teamWithDeliveryLead)

	// built in make() function creates a slice, optional second argument pre-allocates capacity
	var teamSlice = make([]string, 5)
	teamSlice[0] = "Frank"
	teamSlice[1] = "Mike"
	teamSlice[2] = "Daniel"
	teamSlice[3] = "Jacob"
	teamSlice[4] = "Patrick"
	// built in append() operation can
	teamSlice = append(teamSlice, "Rhett")
	fmt.Println("Team Slice:\n", teamSlice)

	// Supports slicing operator
	fmt.Println("Slice of a slice:\n", teamSlice[2:5])

}
