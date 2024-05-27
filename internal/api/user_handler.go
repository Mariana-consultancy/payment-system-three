package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Customer struct {
	FullName           string
	DateOfBirth        time.Time
	GovernmentIssuedID string
	Password           string
	Email              string
	PhoneNumber        string
	ProofOfAddress     string
}

var customers []Customer

func main() {
	
	// Customer registeration 
	reader := bufio.NewReader(os.Stdin)
	customer := registerCustomer(reader)
	if customer != nil {
		customers = append(customers, *customer)
		fmt.Printf("Customer Registered: %+v\n", *customer)
		fmt.Println("All Customers:", customers)
	}

}

func registerCustomer(reader *bufio.Reader) *Customer {
	fullName := readInput(reader, "Enter Full Name: ")

	dobStr := readInput(reader, "Enter Date of Birth (YYYY-MM-DD): ")
	dateOfBirth, err := parseDate(dobStr)
	if err != nil {
		fmt.Println("Invalid date format.")
		return nil
	}

	govID := readInput(reader, "Enter Government Issued Identification: ")
	password := readInput(reader, "Enter Password: ")
	email := readInput(reader, "Enter Email Address: ")
	phone := readInput(reader, "Enter Phone Number: ")
	proofOfAddress := readInput(reader, "Enter Proof of Address: ")

	return &Customer{
		FullName:           fullName,
		DateOfBirth:        dateOfBirth,
		GovernmentIssuedID: govID,
		Password:           password,
		Email:              email,
		PhoneNumber:        phone,
		ProofOfAddress:     proofOfAddress,
	}
}

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func parseDate(dateStr string) (time.Time, error) {
	layout := "2019-19-12"
	return time.Parse(layout, dateStr)
}
