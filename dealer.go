package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode int
}

type Dealer struct {
	DealerId int
	Address
}

// initializes dealers.csv with sample data
func DealerInitialize() {
	a, err := os.OpenFile("dealers.csv", os.O_APPEND, 0644)
	Check(err)
	defer a.Close()
	_, err = a.Write([]byte("DealerId,Street,City,State,Zip\n"))
	Check(err)
	_, err = a.Write([]byte("1,13694 E Iliff Ave,Aurora,CO,80014\n"))
	Check(err)
	_, err = a.Write([]byte("2,575 Clayton St,Denver,CO,80206\n"))
	Check(err)
	_, err = a.Write([]byte("3,5440 W 29th Ave,Denver,CO,80214\n"))
	Check(err)
}

// appends new entry to dealers.csv with values from parameters
func DealersAppend(Dealers [][]string, street string, city string, state string, zipcode string) {
	latestDealerId, err := strconv.Atoi(Dealers[(len(Dealers) - 1)][0])
	Check(err)
	newZipCode, err := strconv.Atoi(zipcode)
	Check(err)
	newAddress := Address{Street: street, City: city, State: state, ZipCode: newZipCode}
	newDealer := Dealer{DealerId: (latestDealerId + 1), Address: newAddress}
	DealerId := strconv.Itoa(newDealer.DealerId)
	newEntry := [][]string{{DealerId, newDealer.Street, newDealer.City, newAddress.State, zipcode}}
	a, err := os.OpenFile("dealers.csv", os.O_APPEND, 0644)
	Check(err)
	defer a.Close()
	writer := csv.NewWriter(a)
	defer writer.Flush()
	for _, v := range newEntry {
		err := writer.Write(v)
		Check(err)
	}
	Check(err)
	fmt.Println()
	fmt.Println("New entry added to Dealers records successfully.")
	time.Sleep(1 * time.Second)
}

// recieves user input to get values for variables and store new entries in dealers.csv
func GetDealerInfoFromUser(Dealers [][]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter Dealer Street Address:")
	street, err := reader.ReadString('\n')
	Check(err)
	street = strings.TrimSpace(street)
	fmt.Println()
	fmt.Print("Enter Dealer City:")
	city, err := reader.ReadString('\n')
	Check(err)
	city = strings.TrimSpace(city)
	fmt.Println()
	fmt.Print("Enter Dealer State:")
	state, err := reader.ReadString('\n')
	Check(err)
	state = strings.TrimSpace(state)
	fmt.Println()
	fmt.Print("Enter Dealer ZIP Code:")
	zipcode, err := reader.ReadString('\n')
	Check(err)
	zipcode = strings.TrimSpace(zipcode)
	userResponse := ""
append:
	for userResponse != "quit" {
		fmt.Println()
		fmt.Println("Is the following information correct?")
		fmt.Println()
		fmt.Println("Dealer Street Address:", street)
		fmt.Println("City:", city)
		fmt.Println("State:", state)
		fmt.Println("Zipcode:", zipcode)
		time.Sleep(2 * time.Second)
		fmt.Println()
		fmt.Print("(Y/N)?:")
		userResponse, err := reader.ReadString('\n')
		Check(err)
		userResponse = strings.TrimSpace(userResponse)
		userResponse = strings.ToLower(userResponse)
		if userResponse == "y" {
			DealersAppend(Dealers, street, city, state, zipcode)
			break
		} else if userResponse == "n" {
			editResponse := ""
			for editResponse != "6" {
				fmt.Println("Which attribute would you like to edit?")
				fmt.Println()
				fmt.Println("1. Street:", street)
				fmt.Println("2. City:", city)
				fmt.Println("3. State:", state)
				fmt.Println("4. ZIP Code", zipcode)
				fmt.Println("5. Submit new entry")
				fmt.Println("6. Cancel new entry")
				time.Sleep(1 * time.Second)
				fmt.Println()
				fmt.Print("Your Response: ")
				editResponse, err = reader.ReadString('\n')
				Check(err)
				editResponse = strings.TrimSpace(editResponse)
				if editResponse == "1" {
					street = EditAttribute(street, "Dealer Street Address")
				} else if editResponse == "2" {
					city = EditAttribute(city, "Dealer City")
				} else if editResponse == "3" {
					state = EditAttribute(state, "Dealer State")
				} else if editResponse == "4" {
					zipcode = EditAttribute(zipcode, "Dealer ZIP Code")
				} else if editResponse == "5" {
					DealersAppend(Dealers, street, city, state, zipcode)
					break append
				} else if editResponse == "6" {
					fmt.Println("New entry cancelled. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break append
				} else {
					InvalidResponse()
				}
			}
		} else {
			InvalidYesNo()
		}
	}
}
