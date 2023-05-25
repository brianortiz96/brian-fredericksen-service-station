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

type Customer struct {
	CustomerId   int
	Name         string
	VehicleMake  string
	VehicleModel string
	VehicleYear  string
	Phone        string
}

// initializes customers.csv after it's created by populating it with sample data
func CustomerInitialize() {
	a, err := os.OpenFile("customers.csv", os.O_APPEND, 0644)
	Check(err)
	defer a.Close()
	_, err = a.Write([]byte("CustomerId,Name,VehicleMake,VehicleModel,VehicleYear,Phone\n"))
	Check(err)
	_, err = a.Write([]byte("1,John Smith,Mercedes,ML320,2003,(720)123-4567\n"))
	Check(err)
	_, err = a.Write([]byte("2,Kevin James,Toyota,Prius,2014,(808)808-8008\n"))
	Check(err)
	_, err = a.Write([]byte("3,Barbara Esparza,Toyota,RAV4,2018,(970)623-7452\n"))
	Check(err)
	_, err = a.Write([]byte("4,Laura Hill,Audi,A4,2023,(303)653-2746\n"))
	Check(err)
	_, err = a.Write([]byte("5,Raul Gomez,Hyundai,Sonata,2009,(719)908-7238\n"))
	Check(err)
	_, err = a.Write([]byte("6,Carol Jones,Suzuki,Swift,2009,(303)886-4193\n"))
	Check(err)
	_, err = a.Write([]byte("7,Robert Otero,Toyota,Yaris,2020,(303)034-2085\n"))
	Check(err)
	_, err = a.Write([]byte("8,Charles Yi,BMW,X6,2017,(720)909-5309\n"))
	Check(err)
	_, err = a.Write([]byte("9,Michelle Rollins,Honda,Odyssey,2005,(719)895-0468\n"))
	Check(err)
	_, err = a.Write([]byte("10,Kentavious Caldwell,Suburu,Crosstek,2015,(970)160-5593\n"))
	Check(err)
}

// appends new entries to customers.csv
func CustomersAppend(Customers [][]string, name string, make string, model string, year string, phone string) {
	latestCustomerId, err := strconv.Atoi(Customers[(len(Customers) - 1)][0])
	Check(err)
	newCustomer := Customer{CustomerId: (latestCustomerId + 1), Name: name, VehicleMake: make, VehicleModel: model, VehicleYear: year, Phone: phone}
	customerId := strconv.Itoa(newCustomer.CustomerId)
	newEntry := [][]string{{customerId, newCustomer.Name, newCustomer.VehicleMake, newCustomer.VehicleModel, newCustomer.VehicleYear, newCustomer.Phone}}
	a, err := os.OpenFile("customers.csv", os.O_APPEND, 0644)
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
	fmt.Println("New entry added to Customers records successfully.")
	time.Sleep(1 * time.Second)
}

// recieves user input to get values for variables and store new entries in customers.csv
func GetCustomerInfoFromUser(Customers [][]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter Customer Name: ")
	name, err := reader.ReadString('\n')
	Check(err)
	name = strings.TrimSpace(name)
	fmt.Println()
	fmt.Print("Enter Vehicle Make: ")
	make, err := reader.ReadString('\n')
	Check(err)
	make = strings.TrimSpace(make)
	fmt.Println()
	fmt.Print("Enter Vehicle Model: ")
	model, err := reader.ReadString('\n')
	Check(err)
	model = strings.TrimSpace(model)
	fmt.Println()
	fmt.Print("Enter Vehicle Year: ")
	year, err := reader.ReadString('\n')
	Check(err)
	year = strings.TrimSpace(year)
	fmt.Println()
	fmt.Print("Enter Customer Phone Number: ")
	phone, err := reader.ReadString('\n')
	Check(err)
	phone = strings.TrimSpace(phone)
	userResponse := ""
append:
	for userResponse != "quit" {
		fmt.Println()
		fmt.Println("Is the following information correct?")
		fmt.Println()
		fmt.Println("Name: ", name)
		fmt.Println("Make: ", make)
		fmt.Println("Model: ", model)
		fmt.Println("Year: ", year)
		fmt.Println("Phone: ", phone)
		time.Sleep(2 * time.Second)
		fmt.Println()
		fmt.Print("(Y/N)?:")
		userResponse, err := reader.ReadString('\n')
		Check(err)
		userResponse = strings.TrimSpace(userResponse)
		userResponse = strings.ToLower(userResponse)
		if userResponse == "y" {
			CustomersAppend(Customers, name, make, model, year, phone)
			break
		} else if userResponse == "n" {
			editResponse := ""
			for editResponse != "6" {
				fmt.Println()
				fmt.Println("Which attribute would you like to edit?")
				fmt.Println()
				fmt.Println("1: Name:", name)
				fmt.Println("2: Vehicle Make:", make)
				fmt.Println("3: Vehicle Model:", model)
				fmt.Println("4: Vehicle Year:", year)
				fmt.Println("5: Customer Phone Number:", phone)
				fmt.Println("6: Submit new entry")
				fmt.Println("7: Cancel new entry")
				time.Sleep(1 * time.Second)
				fmt.Println()
				fmt.Print("Your Response: ")
				editResponse, err = reader.ReadString('\n')
				Check(err)
				editResponse = strings.TrimSpace(editResponse)
				if editResponse == "1" {
					name = EditAttribute(name, "Customer Name")
				} else if editResponse == "2" {
					make = EditAttribute(make, "Vehicle Make")
				} else if editResponse == "3" {
					model = EditAttribute(model, "Vehicle Model")
				} else if editResponse == "4" {
					year = EditAttribute(year, "Vehicle Year")
				} else if editResponse == "5" {
					phone = EditAttribute(phone, "Customer Phone Number")
				} else if editResponse == "6" {
					CustomersAppend(Customers, name, make, model, year, phone)
					break append
				} else if editResponse == "7" {
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
