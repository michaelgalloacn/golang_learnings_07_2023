package main

import (
	"fmt"
	"strconv"
)

type TrialBalance struct {
	companyName    string
	debitsInCents  uint64
	creditsInCents uint64
}
type Checkbook struct {
	bankBalance    uint64
	debitsInCents  uint64
	creditsInCents uint64
}

type Balanceable interface {
	isBalanced() bool
}

func isBalanced(balance TrialBalance) bool {

	return balance.creditsInCents == balance.debitsInCents
}

func (checkBook Checkbook) isBalanced() bool {
	sum := checkBook.debitsInCents - checkBook.creditsInCents
	return sum == checkBook.bankBalance
}

func (balance TrialBalance) isBalanced() bool {

	return balance.creditsInCents == balance.debitsInCents
}

// func (checkBook Checkbook) getCredits() uint64{

// }

func isBalanceable(balanceable Balanceable) string {
	return ("Something balanceable was sent, balanced status :\n" + strconv.FormatBool(balanceable.isBalanced()))
}

func main() {

	FranksDeli := TrialBalance{
		companyName:    "Frank's Deli",
		debitsInCents:  50_459_27,
		creditsInCents: 50_459_27,
	}

	var franksIsBalanced bool = isBalanced(FranksDeli)
	fmt.Println("Frank's Deli is Balanced:\n", franksIsBalanced)

	MikesNonsenseSupply := TrialBalance{
		companyName:   "Mike's Nonsense Supply",
		debitsInCents: 145_12,
		// Equivalent to creditsInCents being 0
	}
	mikesIsBalanced := isBalanced(MikesNonsenseSupply)
	fmt.Println("Mike's Nonsense Supply is initially Balanced\n", mikesIsBalanced)

	// Assigning a value to an already created struct
	MikesNonsenseSupply.creditsInCents = 145_12
	mikesIsBalanced = isBalanced(MikesNonsenseSupply)
	fmt.Println("Mike's Nonsense Supply is now Balanced\n", mikesIsBalanced)

	mikesIsBalanced = MikesNonsenseSupply.isBalanced()
	fmt.Println(mikesIsBalanced)

	// Example of using an interface method across types
	fmt.Println(isBalanceable(MikesNonsenseSupply))

	MikesCheckBook := Checkbook{
		bankBalance:    100,
		debitsInCents:  150,
		creditsInCents: 500,
	}
	fmt.Println(isBalanceable(MikesCheckBook))

}
