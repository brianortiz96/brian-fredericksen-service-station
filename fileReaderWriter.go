package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// checks if the four required CSV files for the program are in the directory
// if any file is not found, this function calls it's initialize function to initialize it
func fileCheck(fileName string) {
	CsvCreated := false
	files, err := os.ReadDir(".")
	Check(err)

	for _, file := range files {
		directoryFile := strings.TrimSpace(file.Name())
		if directoryFile == fileName {
			CsvCreated = true
		}
	}

	if !CsvCreated {
		err := os.WriteFile(fileName, []byte(""), 0666)
		Check(err)
		if fileName == "customers.csv" {
			fmt.Println()
			fmt.Println("Initializing customers.csv...")
			time.Sleep(1 * time.Second)
			CustomerInitialize()
		} else if fileName == "dealers.csv" {
			fmt.Println()
			fmt.Println("Initializing dealers.csv...")
			time.Sleep(1 * time.Second)
			DealerInitialize()
		} else if fileName == "serviceTransactionHistory.csv" {
			fmt.Println()
			fmt.Println("Initializing serviceTransactionHistory.csv...")
			time.Sleep(1 * time.Second)
			TransHistInitialize()
		} else if fileName == "couponNotificationHistory.csv" {
			fmt.Println()
			fmt.Println("Initializing couponNotificationHistory.csv...")
			time.Sleep(1 * time.Second)
			CoupNotifHistInitialize()
		}
	}
}

// reads files from csv file into slice
func ReadCsv(fileName string) ([]string, [][]string, error) {
	c, err := os.Open(fileName)
	csvReader := csv.NewReader(c)
	Check(err)
	defer c.Close()
	header, err := csvReader.Read()
	Check(err)
	lines, err := csvReader.ReadAll()
	Check(err)
	return header, lines, nil
}

// updates csv file by deleting it and replacing it with the newer data present in the respective slice
func UpdateCSV(records [][]string, fileName string) {
	d := os.Remove(fileName)
	Check(d)
	a, err := os.Create(fileName)
	Check(err)
	f, err := os.OpenFile(fileName, os.O_APPEND, 0644)
	Check(err)
	defer a.Close()
	defer f.Close()
	if fileName == "customers.csv" {
		_, err = f.Write([]byte("CustomerId,Name,VehicleMake,VehicleModel,VehicleYear,Phone\n"))
		Check(err)
	} else if fileName == "dealers.csv" {
		_, err = f.Write([]byte("DealerId,Street,City,State,Zip\n"))
		Check(err)
	} else if fileName == "serviceTransactionHistory" {
		_, err = f.Write([]byte("CustomerID,DealerID,Oil Change,Car Wash,Year,Month,Day\n"))
		Check(err)
	}
	err = csv.NewWriter(f).WriteAll(records)
	Check(err)
}
