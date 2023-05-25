package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// called to inform a user when they've provided an invalid response to a menu with numbered options
func InvalidResponse() {
	fmt.Println()
	fmt.Println("Error. Invalid response. Please choose a number from the menu.")
	time.Sleep(1 * time.Second)
}

// called to inform a user when they've provided an invalid response to a yes or no option
func InvalidYesNo() {
	fmt.Println()
	fmt.Println("Error. Invalid response. Please respond with Y for Yes or N for No.")
	time.Sleep(1 * time.Second)
}

// reads data from a slice
func ReadRecords(records [][]string, header []string) {
	fmt.Println()
	fmt.Println(header)
	fmt.Println()
	for x := 0; x < len(records); x++ {
		fmt.Println(records[x])
	}
}

// edits a value that is going to be appended to a slice and subsequently a CSV file
func EditAttribute(editInput string, attributeName string) (editOutput string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter ", attributeName)
	fmt.Print(": ")
	editOutput, err := reader.ReadString('\n')
	Check(err)
	editOutput = strings.TrimSpace(editOutput)
	return editOutput
}

// user menu for editing the existing records in the slices and respective CSV Files
func UpdateRecords(records [][]string, header []string, fileName string) {
	reader := bufio.NewReader(os.Stdin)
	if fileName == "customers.csv" {
		fmt.Println()
		fmt.Print("Enter the Customer ID of the record you want to edit: ")
		customerId, err := reader.ReadString('\n')
		Check(err)
		customerId = strings.TrimSpace(customerId)
		editCustomerID, err := strconv.Atoi(customerId)
		Check(err)
		if editCustomerID < len(records) {
			editResponse := ""
			for editResponse != "7" {
				fmt.Println("Which attribute would you like to edit?")
				fmt.Println()
				fmt.Println("1. Name:", records[editCustomerID-1][1])
				fmt.Println("2. Vehicle Make:", records[editCustomerID-1][2])
				fmt.Println("3. Vehicle Model:", records[editCustomerID-1][3])
				fmt.Println("4. Vehicle Year:", records[editCustomerID-1][4])
				fmt.Println("5. Customer Phone Number:", records[editCustomerID-1][5])
				fmt.Println("6. Save edits and quit")
				fmt.Println("7. Cancel edits and quit")
				time.Sleep(1 * time.Second)
				fmt.Println()
				fmt.Print("Your Response: ")
				editResponse, err = reader.ReadString('\n')
				Check(err)
				editResponse = strings.TrimSpace(editResponse)
				if editResponse == "1" {
					records[editCustomerID-1][1] = EditAttribute(records[editCustomerID-1][1], "Customer Name")
				} else if editResponse == "2" {
					records[editCustomerID-1][2] = EditAttribute(records[editCustomerID-1][2], "Vehicle Make")
				} else if editResponse == "3" {
					records[editCustomerID-1][3] = EditAttribute(records[editCustomerID-1][3], "Vehicle Model")
				} else if editResponse == "4" {
					records[editCustomerID-1][4] = EditAttribute(records[editCustomerID-1][4], "Vehicle Year")
				} else if editResponse == "5" {
					records[editCustomerID-1][5] = EditAttribute(records[editCustomerID-1][5], "Customer Phone Number")
				} else if editResponse == "6" {
					UpdateCSV(records, fileName)
					fmt.Println()
					fmt.Print("Edits to entry with Customer ID ", editCustomerID, " saved successfully to ", fileName)
					fmt.Println("...")
					time.Sleep(1 * time.Second)
					fmt.Println()
					fmt.Println("Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else if editResponse == "7" {
					fmt.Println()
					fmt.Println("Editing cancelled. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else {
					InvalidResponse()
				}
			}
		} else {
			fmt.Println()
			fmt.Println("No record exists with that Customer ID.")
			time.Sleep(1 * time.Second)
			fmt.Println()
			fmt.Println("Returning to Main Menu...")
			time.Sleep(1 * time.Second)
		}
	} else if fileName == "dealers.csv" {
		fmt.Println()
		fmt.Print("Enter the Dealer ID of the record you want to edit: ")
		dealerId, err := reader.ReadString('\n')
		Check(err)
		dealerId = strings.TrimSpace(dealerId)
		editDealerID, err := strconv.Atoi(dealerId)
		Check(err)
		if editDealerID < len(records) {
			editResponse := ""
			for editResponse != "7" {
				fmt.Println("Which attribute would you like to edit?")
				fmt.Println()
				fmt.Println("1. Street Address:", records[editDealerID-1][1])
				fmt.Println("2. City:", records[editDealerID-1][2])
				fmt.Println("3. State:", records[editDealerID-1][3])
				fmt.Println("4. ZIP Code:", records[editDealerID-1][4])
				fmt.Println("5. Save edits and quit")
				fmt.Println("6. Cancel edits and quit")
				time.Sleep(1 * time.Second)
				fmt.Println()
				fmt.Print("Your Response: ")
				editResponse, err = reader.ReadString('\n')
				Check(err)
				editResponse = strings.TrimSpace(editResponse)
				if editResponse == "1" {
					records[editDealerID-1][1] = EditAttribute(records[editDealerID-1][1], "Street Address")
				} else if editResponse == "2" {
					records[editDealerID-1][2] = EditAttribute(records[editDealerID-1][2], "City")
				} else if editResponse == "3" {
					records[editDealerID-1][3] = EditAttribute(records[editDealerID-1][3], "State")
				} else if editResponse == "4" {
					records[editDealerID-1][4] = EditAttribute(records[editDealerID-1][4], "ZIP Code")
				} else if editResponse == "5" {
					UpdateCSV(records, fileName)
					fmt.Println()
					fmt.Println("Edits to entry with Dealer ID ", editDealerID, " saved successfully to ", fileName)
					fmt.Print("...")
					time.Sleep(1 * time.Second)
					fmt.Println()
					fmt.Println("Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else if editResponse == "6" {
					fmt.Println()
					fmt.Println("Editing cancelled. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else {
					InvalidResponse()
				}
			}
		} else {
			fmt.Println()
			fmt.Println("No record exists with that Customer ID.")
			time.Sleep(1 * time.Second)
			fmt.Println()
			fmt.Println("Returning to Main Menu...")
			time.Sleep(1 * time.Second)
		}
	} else if fileName == "Service Transactions" {
		fmt.Println()
		fmt.Println("Here is a numbered list of every entry in the Service Transaction History:")
		fmt.Println()
		fmt.Println(header)
		for x := 0; x < len(records); x++ {
			fmt.Print((x + 1))
			fmt.Println(":", records[x])
		}
		fmt.Println()
		fmt.Print("Enter the number corresponding to the entry you'd like to edit: ")
		userChoice, err := reader.ReadString('\n')
		Check(err)
		choice, err := strconv.Atoi(userChoice)
		Check(err)
		choice = choice - 1
		if choice > (len(records)-1) || choice < 0 {
			fmt.Println()
			fmt.Print("There is no entry number", choice)
			fmt.Println("...")
			time.Sleep(1 * time.Second)
			fmt.Println()
			fmt.Println("Aborting edit process and returning to Main Menu...")
			time.Sleep(1 * time.Second)
		} else {
			editResponse := ""
			for editResponse != "6" {
				fmt.Println("Which attribute would you like to edit?")
				fmt.Println()
				fmt.Println("1: Oil Change:", records[choice][2])
				fmt.Println("2: Car Wash:", records[choice][3])
				fmt.Println("3: Service Date Year:", records[choice][4])
				fmt.Println("4: Service Date Month:", records[choice][5])
				fmt.Println("5: Service Date Day:", records[choice][6])
				fmt.Println("6. Save edits and quit")
				fmt.Println("7. Cancel edits and quit")
				time.Sleep(1 * time.Second)
				fmt.Println()
				fmt.Print("Your Response: ")
				editResponse, err = reader.ReadString('\n')
				Check(err)
				if editResponse == "1" {
					records[choice][2] = EditAttribute(records[choice][2], "oilChange")
				} else if editResponse == "2" {
					records[choice][3] = EditAttribute(records[choice][3], "carWash")
				} else if editResponse == "3" {
					records[choice][4] = EditAttribute(records[choice][4], "Service Date Year")
				} else if editResponse == "4" {
					records[choice][5] = EditAttribute(records[choice][5], "Service Date Month")
				} else if editResponse == "5" {
					records[choice][6] = EditAttribute(records[choice][6], "Service Date Day")
				} else if editResponse == "6" {
					UpdateCSV(records, fileName)
					fmt.Println()
					fmt.Println("Edits to entry number", choice+1, "saved successfully to", fileName)
					fmt.Print("...")
					time.Sleep(1 * time.Second)
					fmt.Println()
					fmt.Println("Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else if editResponse == "7" {
					fmt.Println()
					fmt.Println("Editing cancelled. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else {
					InvalidResponse()
				}
			}
		}
	}
}

// user menu for deleting the existing records in the slices and respective CSV Files
func DeleteRecord(records [][]string, header []string, fileName string, recordsName string) {
	reader := bufio.NewReader(os.Stdin)
	if fileName == "customers.csv" || fileName == "dealers.csv" {
		fmt.Println()
		fmt.Print("Enter the ", recordsName, " ID of the record you want to delete: ")
		iD, err := reader.ReadString('\n')
		Check(err)
		iD = strings.TrimSpace(iD)
		editId, err := strconv.Atoi(iD)
		Check(err)
		editId = editId - 1
		if editId < len(records) {
			deleteResponse := ""
			for deleteResponse != "quit" {
				fmt.Println()
				fmt.Println("Are you SURE you want to delete the following record?:")
				fmt.Println(records[editId])
				fmt.Println()
				fmt.Print("(Y/N)?:")
				deleteResponse, err := reader.ReadString('\n')
				Check(err)
				deleteResponse = strings.TrimSpace(deleteResponse)
				deleteResponse = strings.ToLower(deleteResponse)
				if deleteResponse == "y" {
					for x := 1; x < len(records[editId]); x++ {
						records[editId][x] = "deleted"
					}
					UpdateCSV(records, fileName)
					fmt.Println()
					fmt.Println("Record deleted successfully. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else if deleteResponse == "n" {
					fmt.Println()
					fmt.Println("Record was not deleted. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else {
					InvalidYesNo()
				}
			}
		} else {
			fmt.Println()
			fmt.Print("No ", recordsName, " exists with ", recordsName, " ID: ", (editId + 1))
			fmt.Println("...")
			time.Sleep(1 * time.Second)
		}
	} else if fileName == "serviceTransactionHistory.csv" {
		fmt.Println()
		fmt.Println("Here is a numbered list of every entry in the Service Transaction History:")
		fmt.Println()
		fmt.Println(header)
		for x := 0; x < len(records); x++ {
			fmt.Print((x + 1))
			fmt.Println(":", records[x])
		}
		fmt.Println()
		fmt.Print("Enter the number corresponding to the entry you'd like to delete: ")
		userChoice, err := reader.ReadString('\n')
		Check(err)
		choice, err := strconv.Atoi(userChoice)
		Check(err)
		choice = choice - 1
		if choice > (len(records)-1) || choice < 0 {
			fmt.Println()
			fmt.Print("There is no entry number", choice)
			fmt.Println("...")
			time.Sleep(1 * time.Second)
			fmt.Println()
			fmt.Println("Aborting deletion process and returning to Main Menu...")
			time.Sleep(1 * time.Second)
		} else {
			deleteResponse := ""
			for deleteResponse != "quit" {
				fmt.Println()
				fmt.Println("Are you SURE you want to delete the following record?:")
				fmt.Println(records[choice])
				fmt.Println()
				fmt.Print("(Y/N)?:")
				deleteResponse, err := reader.ReadString('\n')
				Check(err)
				deleteResponse = strings.TrimSpace(deleteResponse)
				deleteResponse = strings.ToLower(deleteResponse)
				if deleteResponse == "y" {
					for x := 0; x < len(records[choice]); x++ {
						records[choice][x] = "deleted"
					}
					UpdateCSV(records, fileName)
					fmt.Println()
					fmt.Println("Record deleted successfully. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else if deleteResponse == "n" {
					fmt.Println()
					fmt.Println("Record was not deleted. Returning to Main Menu...")
					time.Sleep(1 * time.Second)
					break
				} else {
					InvalidYesNo()
				}
			}
		}
	}
}

// user menu for navigating records and using append, edit, and delete functionality
func RecordsMenu(recordsName string, trasnHist [][]string, coupHist [][]string, dealers [][]string, customers [][]string, recordsHeader []string, fileName string) {
	userExit := ""
menu:
	for userExit != "quit" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println()
		fmt.Println("What would you like to do with the", recordsName, "records?")
		time.Sleep(1 * time.Second)
		fmt.Println()
		fmt.Println("1: View Records")
		fmt.Println("2: Add New Record")
		fmt.Println("3: Edit Record")
		fmt.Println("4: Delete Record")
		fmt.Println("5: Quit to Main Menu")
		time.Sleep(1 * time.Second)
		fmt.Println()
		fmt.Print("Your response:")
		menuChoice, err := reader.ReadString('\n')
		Check(err)
		menuChoice = strings.TrimSpace(menuChoice)
		if menuChoice == "1" {
			if recordsName == "Customers" {
				ReadRecords(customers, recordsHeader)
			} else if recordsName == "Dealers" {
				ReadRecords(dealers, recordsHeader)
			} else if recordsName == "Service Transactions" {
				dateChoice := ""
				for dateChoice != "quit" {
					fmt.Println()
					fmt.Println("You can choose between viewing records from before a specific date, or all records prior to the current date.")
					fmt.Println()
					fmt.Println("1: Current Date")
					fmt.Println("2: Specific Date")
					fmt.Println("3: Quit to Main Menu")
					fmt.Println()
					time.Sleep(1 * time.Second)
					fmt.Print("Your response: ")
					dateChoice, err := reader.ReadString('\n')
					Check(err)
					dateChoice = strings.TrimSpace(dateChoice)
					if dateChoice == "1" {
						currentTime := time.Now()
						DateReport(currentTime.Year(), int(currentTime.Month()), currentTime.Day(), customers, trasnHist)
						break menu
					} else if dateChoice == "2" {
						fmt.Println()
						fmt.Print("Enter the year of your chosen date:")
						userYear, err := reader.ReadString('\n')
						Check(err)
						userYear = strings.TrimSpace(userYear)
						year, err := strconv.Atoi(userYear)
						Check(err)
						fmt.Println()
						fmt.Print("Enter the month of your chosen date:")
						userMonth, err := reader.ReadString('\n')
						Check(err)
						userMonth = strings.TrimSpace(userMonth)
						month, err := strconv.Atoi(userMonth)
						Check(err)
						fmt.Println()
						fmt.Print("Enter the day of your chosen date:")
						userDay, err := reader.ReadString('\n')
						Check(err)
						userDay = strings.TrimSpace(userDay)
						day, err := strconv.Atoi(userDay)
						Check(err)
						DateReport(year, month, day, customers, trasnHist)
						break menu
					} else if dateChoice == "3" {
						fmt.Println()
						fmt.Println("Returning to Main Menu...")
						time.Sleep(1 * time.Second)
						break menu
					} else {
						InvalidResponse()
					}
				}
			}
		} else if menuChoice == "2" {
			if recordsName == "Customers" {
				GetCustomerInfoFromUser(customers)
				break menu
			} else if recordsName == "Dealers" {
				GetDealerInfoFromUser(dealers)
				break menu
			} else if recordsName == "Service Transactions" {
				GetServTransInfoFromUser(trasnHist, customers, dealers)
				break menu
			}
		} else if menuChoice == "3" {
			if recordsName == "Customers" {
				UpdateRecords(customers, recordsHeader, fileName)
				break menu
			} else if recordsName == "Dealers" {
				UpdateRecords(dealers, recordsHeader, fileName)
				break menu
			} else if recordsName == "Service Transactions" {
				UpdateRecords(trasnHist, recordsHeader, fileName)
				break menu
			}
		} else if menuChoice == "4" {
			if recordsName == "Customers" {
				DeleteRecord(customers, recordsHeader, fileName, "Customer")
				break menu
			} else if recordsName == "Dealers" {
				DeleteRecord(dealers, recordsHeader, fileName, "Dealer")
				break menu
			} else if recordsName == "Service Transactions" {
				DeleteRecord(trasnHist, recordsHeader, fileName, "")
				break menu
			}
		} else if menuChoice == "5" {
			fmt.Println()
			fmt.Println("Returning to Main Menu...")
			time.Sleep(1 * time.Second)
			break menu
		} else {
			InvalidResponse()
		}
	}
}
