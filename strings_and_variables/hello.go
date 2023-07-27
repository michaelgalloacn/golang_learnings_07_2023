package main

import "fmt"

// custom type - anything assigned to this type BEHHAVES like a string but is not comparable to a string
// Example of why you'd want a custom type : https://www.sohamkamani.com/golang/enums/
type customMessageType string

// Type Alias - anything assigned to this type IS a string
type aliasedStringMessage = string

func main() {

	// CONST keyword for immutable values, constants
	const constHelloMessage string = "Hello World"
	fmt.Println(constHelloMessage)
	// Trying to update a const will cause a compilation error
	// constHelloMessage = "Hello world updated"

	// Declare  variables with var keyword. Compiler will infer type when possible
	var varMessage = "I don't have a type annotation"
	fmt.Println(varMessage)
	// variables CAN be updated
	// varMessage = "I have changed"
	// fmt.Println(varMessage)

	// Types can be defined explicitly with var [Name] [type]
	var typeAnnotatedMessage string = "I do have a type annotation"
	fmt.Println(typeAnnotatedMessage)

	// Types in Golang have defined zero values. Not a compilation error to define a variable and use immediately
	var emptyString string
	fmt.Println("Uninitialized string:")
	fmt.Println(emptyString)
	fmt.Println("Uninitialized int:")
	var emptyInt int
	fmt.Println(emptyInt)

	// Walrus ( := ) operator is syntactic sugar for declaring and initializing a var
	walrusMessage := "Walrus Hello World"
	fmt.Println(walrusMessage)

	// Variables declared with custom types can operate as base type but are not directly comparable
	var customTypedMessage customMessageType = "I use a custom type"
	fmt.Println(customTypedMessage)
	// attempting to compare without a conversion will be treated as an error
	// var customTypedAsString = "I use a custom type"
	// fmt.Println(customTypedMessage == customTypedAsString)
	// A custom type can be converted back into the base type or visa versa
	// fmt.Println(string(customTypedMessage) == customTypedAsString)
	// fmt.Println(customTypedMessage == customMessageType(customTypedAsString))

	// An aliased type CAN be compared directly to its base type, this val IS type String
	var aliasedMessage aliasedStringMessage = "I am a string under the hood"
	fmt.Println(aliasedMessage)
	var aliasedMessageAsString string = "I am a string under the hood"
	fmt.Println(aliasedMessage == aliasedMessageAsString)

	// var wrong_annotation int = "I have the wrong type annotation and will not compile"
	// fmt.Println(wrong_annotation)

	// var unused_string = "I am unused and by default this will stop execution"

}
