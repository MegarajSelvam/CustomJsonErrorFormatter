package main

import (
	"fmt"
	"strings"
)

var (
	errorMessage = `json: cannot unmarshal number into Go struct field Customer.customer.custId of type string`
)

// Aim is to remove the Struct Name which is getting appended. in below case Customer.customer.custId -> Customer is struct Name.
// input string: "json: cannot unmarshal number into Go struct field Customer.customer.custId of type string"
// output string: "json: cannot unmarshal number into Go struct field customer.custId of type string"
func main() {
	fmt.Println("Before formatting: ", errorMessage)
	fmt.Println("After formatting: ", formatErrorMessage(errorMessage))
}

func formatErrorMessage(errMessage string) string {
	splittedString := strings.Fields(errorMessage)
	for index, value := range splittedString {
		if strings.Contains(value, ".") {
			stackTrace := strings.Split(value, ".")
			if stackTrace != nil && len(stackTrace) >= 1 {
				splittedString[index] = strings.Join(stackTrace[1:], ".")
			}
		}
	}
	return strings.Join(splittedString, " ")
}
