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

// intializes serviceTransactionHistory.csv and provides sample data to work with
func TransHistInitialize() {
	a, err := os.OpenFile("serviceTransactionHistory.csv", os.O_APPEND, 0644)
	Check(err)
	defer a.Close()
	_, err = a.Write([]byte("CustomerID,DealerID,Oil Change,Car Wash,Year,Month,Day\n"))
	Check(err)
	_, err = a.Write([]byte("1,1,true,true,2022,2,22\n"))
	Check(err)
	_, err = a.Write([]byte("2,1,true,false,2022,4,19\n"))
	Check(err)
	_, err = a.Write([]byte("3,2,true,true,2022,5,31\n"))
	Check(err)
	_, err = a.Write([]byte("2,2,false,true,2022,6,7\n"))
	Check(err)
	_, err = a.Write([]byte("4,1,true,true,2022,6,29\n"))
	Check(err)
	_, err = a.Write([]byte("1,3,false,true,2022,7,4\n"))
	Check(err)
	_, err = a.Write([]byte("5,3,true,false,2022,7,31\n"))
	Check(err)
	_, err = a.Write([]byte("6,2,true,true,2022,8,22\n"))
	Check(err)
	_, err = a.Write([]byte("1,1,true,true,2022,8,25\n"))
	Check(err)
	_, err = a.Write([]byte("7,3,true,true,2022,9,10\n"))
	Check(err)
	_, err = a.Write([]byte("4,1,false,true,2022,9,15\n"))
	Check(err)
	_, err = a.Write([]byte("8,2,true,true,2022,9,31\n"))
	Check(err)
	_, err = a.Write([]byte("5,3,false,true,2022,10,11\n"))
	Check(err)
	_, err = a.Write([]byte("2,1,true,true,2022,10,31\n"))
	Check(err)
	_, err = a.Write([]byte("3,3,false,true,2022,11,5\n"))
	Check(err)
	_, err = a.Write([]byte("9,2,true,false,2022,11,17\n"))
	Check(err)
	_, err = a.Write([]byte("6,1,false,true,2022,11,30\n"))
	Check(err)
	_, err = a.Write([]byte("3,2,true,false,2022,12,6\n"))
	Check(err)
	_, err = a.Write([]byte("4,1,true,true,2023,1,5\n"))
	Check(err)
	_, err = a.Write([]byte("5,3,true,true,2023,2,7\n"))
	Check(err)
}

// takes either a user selected date or the current date and shows all records in serviceTransactionHistory.csv from before that date
func DateReport(year int, month int, day int, Customers [][]string, TransHist [][]string) {
	dateReportHeader := []string{"Customer Name", "Vehicle Make", "Vehicle Model", "Vehicle Year", "Customer Phone Number", "Oil Change", "Car Wash", "Year", "Month", "Day"}
	userDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	fmt.Println()
	fmt.Print("All service transactions recorded before ")
	fmt.Print(month)
	fmt.Print("-")
	fmt.Print(day)
	fmt.Print("-")
	fmt.Print(year)
	fmt.Println()
	fmt.Println(dateReportHeader)
	fmt.Println()
	for x := len(TransHist) - 1; x >= 0; x-- {
		year, err := strconv.Atoi(TransHist[x][4])
		Check(err)
		month, err := strconv.Atoi(TransHist[x][5])
		Check(err)
		day, err := strconv.Atoi(TransHist[x][6])
		Check(err)
		recordDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		customerID, err := strconv.Atoi(TransHist[x][0])
		Check(err)
		if userDate.Sub(recordDate) > 0 {
			serviceRecord := make([]string, 10)
			serviceRecord[0] = Customers[customerID][0]
			serviceRecord[1] = Customers[customerID][1]
			serviceRecord[2] = Customers[customerID][2]
			serviceRecord[3] = Customers[customerID][3]
			serviceRecord[4] = Customers[customerID][4]
			serviceRecord[5] = TransHist[x][2]
			serviceRecord[6] = TransHist[x][4]
			serviceRecord[7] = TransHist[x][4]
			serviceRecord[8] = TransHist[x][5]
			serviceRecord[9] = TransHist[x][6]
			fmt.Println(serviceRecord)
		}
	}
}

// appends new entries to serviceTransactionHistory.csv
func ServiceTransactionAppend(TransHist [][]string, customerId string, dealerId string, oilChange string, carWash string, year string, month string, day string) {
	newEntry := [][]string{{customerId, dealerId, oilChange, carWash, year, month, day}}
	a, err := os.OpenFile("serviceTransactionHistory.csv", os.O_APPEND, 0644)
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
	fmt.Println("New entry added to Service Transaction History successfully.")
	time.Sleep(1 * time.Second)
}

// user menu for inputting values that can be used to append new service transactions to the respective slice and CSV file
func GetServTransInfoFromUser(TransHist [][]string, Customers [][]string, Dealers [][]string) {
	append := false
append:
	for !append {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println()
		fmt.Print("Enter Customer ID:")
		customerId, err := reader.ReadString('\n')
		Check(err)
		customerId = strings.TrimSpace(customerId)
		checkedCustomerId, err := strconv.Atoi(customerId)
		Check(err)
		if checkedCustomerId > len(Customers) {
			fmt.Println()
			fmt.Println("Customer ID", checkedCustomerId, "does not exist. A customer must first be registered before")
			fmt.Println("they can be serviced.")
			time.Sleep(1 * time.Second)
			break append
		}
		fmt.Println()
		fmt.Print("Enter Dealer ID:")
		dealerId, err := reader.ReadString('\n')
		Check(err)
		dealerId = strings.TrimSpace(dealerId)
		checkedDealerId, err := strconv.Atoi(dealerId)
		Check(err)
		if checkedDealerId > len(Dealers) {
			fmt.Println("Dealer ID", checkedDealerId, "does not exist. A dealer must first be registered before")
			fmt.Println("they can provide service.")
			time.Sleep(1 * time.Second)
			break append
		}
		fmt.Println()
		fmt.Print("Oil Change? (true/false):")
		oilChange, err := reader.ReadString('\n')
		Check(err)
		oilChange = strings.TrimSpace(oilChange)
		oilChange = strings.ToLower(oilChange)
		fmt.Println()
		fmt.Print("Car Wash? (true/false):")
		carWash, err := reader.ReadString('\n')
		Check(err)
		carWash = strings.TrimSpace(carWash)
		carWash = strings.ToLower(carWash)
		fmt.Println()
		fmt.Print("Enter Year of Service Date:")
		year, err := reader.ReadString('\n')
		Check(err)
		year = strings.TrimSpace(year)
		fmt.Println()
		fmt.Print("Enter Month of Service Date:")
		month, err := reader.ReadString('\n')
		Check(err)
		month = strings.TrimSpace(month)
		fmt.Println()
		fmt.Print("Enter Year of Service Day:")
		day, err := reader.ReadString('\n')
		Check(err)
		day = strings.TrimSpace(day)
		fmt.Println()
		userResponse := ""
		for userResponse != "quit" {
			fmt.Println()
			fmt.Println("Is the following information correct?")
			fmt.Println("(Note: Oil Change and Car Wash should have values of either 'true' or 'false')")
			fmt.Println()
			fmt.Println("Customer ID:", customerId)
			fmt.Println("Dealer ID:", dealerId)
			fmt.Println("Oil Change:", oilChange)
			fmt.Println("Car Wash:", carWash)
			fmt.Println("Service Date Year:", year)
			fmt.Println("Service Date Month:", month)
			fmt.Println("Service Date Day:", day)
			time.Sleep(2 * time.Second)
			fmt.Println()
			fmt.Print("(Y/N)?:")
			userResponse, err := reader.ReadString('\n')
			Check(err)
			userResponse = strings.TrimSpace(userResponse)
			userResponse = strings.ToLower(userResponse)
			if userResponse == "y" {
				latestCustomerId := strconv.Itoa(checkedCustomerId)
				latestDealerID := strconv.Itoa(checkedDealerId)
				ServiceTransactionAppend(TransHist, latestCustomerId, latestDealerID, oilChange, carWash, year, month, day)
				break append
			} else if userResponse == "n" {
				editResponse := ""
				for editResponse != "6" {
					fmt.Println("Which attribute would you like to edit?")
					fmt.Println()
					fmt.Println("Customer ID and Dealer ID cannot be edited in this menu.")
					fmt.Println("To edit either of these attributes, please quit and start")
					fmt.Println("from the beginning of the add entry process.")
					fmt.Println()
					fmt.Println("1: Oil Change:", oilChange)
					fmt.Println("2: Car Wash:", carWash)
					fmt.Println("3: Service Date Year:", year)
					fmt.Println("4: Service Date Month:", month)
					fmt.Println("5: Service Date Day:", day)
					fmt.Println("6: Submit new entry")
					fmt.Println("7: Cancel new entry")
					time.Sleep(1 * time.Second)
					fmt.Println()
					fmt.Print("Your Response: ")
					editResponse, err = reader.ReadString('\n')
					Check(err)
					editResponse = strings.TrimSpace(editResponse)
					if editResponse == "1" {
						oilChange = EditAttribute(oilChange, "oilChange")
					} else if editResponse == "2" {
						carWash = EditAttribute(carWash, "carWash")
					} else if editResponse == "3" {
						year = EditAttribute(year, "Service Date Year")
					} else if editResponse == "4" {
						month = EditAttribute(month, "Service Date Month")
					} else if editResponse == "5" {
						day = EditAttribute(day, "Service Date Day")
					} else if editResponse == "6" {
						latestCustomerId := strconv.Itoa(checkedCustomerId)
						latestDealerID := strconv.Itoa(checkedDealerId)
						ServiceTransactionAppend(TransHist, latestCustomerId, latestDealerID, oilChange, carWash, year, month, day)
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
}
