package main

import "fmt"

type customMessageType string

type aliasedStringMessage = string

func main() {

	const constHelloMessage string = "Hello World"
	fmt.Println(constHelloMessage)

	walrusMessage := "Walrus Hello World"
	fmt.Println(walrusMessage)

	var notTypeAnnotatedMessage = "I don't have a type annotation"
	fmt.Println(notTypeAnnotatedMessage)

	var typeAnnotatedMessage string = "I do have a type annotation"
	fmt.Println(typeAnnotatedMessage)

	var customTypedMessage customMessageType = "I use a custom type"
	fmt.Println(customTypedMessage)

	// A custom type can not be compared directly to its base type
	// var customTypedAsString = "I use a custom type"
	// fmt.Println(customTypedMessage == customTypedAsString)
	// A custom type can be converted back into the base type or visa versa
	// fmt.Println(string(customTypedMessage) == customTypedAsString)
	// fmt.Println(customTypedMessage == customMessageType(customTypedAsString))

	// An aliased type CAN be compared directly to its base type
	var aliasedMessage aliasedStringMessage = "I am a string under the hood"
	fmt.Println(aliasedMessage)
	var aliasedMessageAsString string = "I am a string under the hood"
	fmt.Println(aliasedMessage == aliasedMessageAsString)

	// var wrong_annotation int = "I have the wrong type annotation and will not compile"
	// fmt.Println(wrong_annotation)

	// var unused_string = "I am unused and by default this will stop execution"

	var emptyString string
	fmt.Println(emptyString)
	var emptyInt int
	fmt.Println(emptyInt)

}
