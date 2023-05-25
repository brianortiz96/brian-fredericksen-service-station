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

// initializes couponNotificationHistory.csv, adds in sample data
func CoupNotifHistInitialize() {
	a, err := os.OpenFile("couponNotificationHistory.csv", os.O_APPEND, 0644)
	Check(err)
	defer a.Close()
	_, err = a.Write([]byte("Customer ID,Dealer ID,Service Year,Service Month,Service Day,Notification Year,Notification Month,Notification Day\n"))
	Check(err)
	_, err = a.Write([]byte("1,1,2022,2,22,2022,8,22\n"))
	Check(err)
	_, err = a.Write([]byte("2,1,2022,4,19,2022,10,19\n"))
	Check(err)
	_, err = a.Write([]byte("3,2,2022,5,31,2022,12,1\n"))
	Check(err)
	_, err = a.Write([]byte("4,1,2022,6,29,2022,12,29\n"))
	Check(err)
	_, err = a.Write([]byte("5,3,2022,7,31,2022,1,31\n"))
	Check(err)
}

// finds customers who haven't been for an oil change in over 6 months, notifies them of their oil change coupon
func CouponNotify(coupons [][]string, transHist [][]string, customers [][]string) {
	reader := bufio.NewReader(os.Stdin)
	currentDate := time.Now()
	for x := len(transHist) - 1; x >= 0; x-- {
		transHistID := x
		lastServiceYear, err := strconv.Atoi(transHist[x][4])
		Check(err)
		lastServiceMonth, err := strconv.Atoi(transHist[x][5])
		Check(err)
		lastServiceDay, err := strconv.Atoi(transHist[x][6])
		Check(err)
		lastServiceDate := time.Date(lastServiceYear, time.Month(lastServiceMonth), lastServiceDay, 0, 0, 0, 0, time.Local)
		dayDifference := currentDate.Sub(lastServiceDate)
		if dayDifference.Hours() > 4320 {
			if transHist[x][2] == "true" {
				for x := len(coupons) - 1; x >= 0; x-- {
					checkCustomerId := transHist[transHistID][0]
					customerId, err := strconv.Atoi(transHist[transHistID][0])
					Check(err)
					customerId = customerId - 1
					couponCustomerId, err := strconv.Atoi(coupons[x][0])
					Check(err)
					couponCustomerId = couponCustomerId - 1
					couponServiceYear, err := strconv.Atoi(coupons[x][2])
					Check(err)
					couponServiceMonth, err := strconv.Atoi(coupons[x][3])
					Check(err)
					couponServiceDay, err := strconv.Atoi(coupons[x][4])
					Check(err)
					if lastServiceYear == couponServiceYear && lastServiceMonth == couponServiceMonth && lastServiceDay == couponServiceDay && checkCustomerId == coupons[x][0] {
						continue
					} else if dayDifference.Hours() > 6480 {
						continue
					} else {
						if customerId == couponCustomerId {
							couponChoice := ""
							for couponChoice != "quit" {
								fmt.Println()
								Check(err)
								fmt.Println(customers[customerId][1], "has not had an oil change on their", customers[customerId][4], customers[customerId][2], customers[customerId][3], "since:")
								fmt.Print(lastServiceMonth)
								fmt.Print("-")
								fmt.Print(lastServiceDay)
								fmt.Print("-")
								fmt.Println(lastServiceYear)
								fmt.Println()
								fmt.Println("Would you like to send them a coupon for an oil change service?")
								fmt.Println()
								fmt.Print("(Y/N): ")
								couponChoice, err = reader.ReadString('\n')
								Check(err)
								couponChoice = strings.TrimSpace(couponChoice)
								couponChoice = strings.ToLower(couponChoice)
								if couponChoice == "y" {
									currentDate := time.Now()
									year := strconv.Itoa(currentDate.Year())
									month := strconv.Itoa(int(currentDate.Month()))
									day := strconv.Itoa(currentDate.Day())
									newCoupon := []string{transHist[transHistID][0], transHist[transHistID][1], transHist[transHistID][4], transHist[transHistID][5], transHist[transHistID][6], year, month, day}
									Coupons := [][]string{newCoupon}
									a, err := os.OpenFile("couponNotificationHistory.csv", os.O_APPEND, 0644)
									Check(err)
									defer a.Close()
									writer := csv.NewWriter(a)
									defer writer.Flush()
									err = writer.WriteAll(Coupons)
									Check(err)
									fmt.Println()
									fmt.Println("Coupon was sent.")
									time.Sleep(1 * time.Second)
									break
								} else if couponChoice == "n" {
									fmt.Println()
									fmt.Println("Okay, we won't send them a coupon for now.")
									time.Sleep(1 * time.Second)
									break
								} else {
									InvalidYesNo()
								}
							}
						}
					}
				}
			}
		}
	}
}

// takes a given customer id and shows the coupons they've been sent
func LastCouponDetails(givenId int, CoupHist [][]string, Customers [][]string) {
	for x := 0; x < len(Customers); x++ {
		customer := Customers[x]
		customerId, err := strconv.Atoi(customer[0])
		Check(err)
		if customerId == givenId {
			for x := len(CoupHist) - 1; x >= 0; x-- {
				coupon := CoupHist[x]
				customerId, err := strconv.Atoi(coupon[0])
				Check(err)
				if customerId == givenId {
					for x := 0; x < len(Customers); x++ {
						customer := Customers[x]
						customersCustomerId, err := strconv.Atoi(customer[0])
						Check(err)
						if customerId == customersCustomerId {
							customerName := customer[1]
							customerVehicleMake := customer[2]
							customerVehicleModel := customer[3]
							customerVehicleYear := customer[4]
							fmt.Println()
							fmt.Print("Customer ", customerName, " was notified about their coupon for an oil change on their ")
							fmt.Print(customerVehicleYear, " ", customerVehicleMake, " ", customerVehicleModel, " on ")
							fmt.Print(coupon[6])
							fmt.Print("-")
							fmt.Print(coupon[7])
							fmt.Print("-")
							fmt.Print(coupon[5])
							fmt.Print(".")
							time.Sleep(1 * time.Second)
							fmt.Println()
							break
						}
					}
				}
			}
		}
	}
}

// menu for getting user input to feed to LastCouponDetails
func CouponNotifCheckMenu(customerHeader []string, customers [][]string, coupNotifHist [][]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("Here is the list of customers:")
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	fmt.Println(customerHeader)
	fmt.Println()
	for x := 0; x < len(customers); x++ {
		fmt.Println(customers[x])
	}
	time.Sleep(1 * time.Second)
	fmt.Println()
	fmt.Println("Enter the ID of the customer whose coupon notification history you'd like to check.")
	fmt.Println()
	fmt.Print("Customer ID:")
	userCustomerID, err := reader.ReadString('\n')
	Check(err)
	userCustomerID = strings.TrimSpace(userCustomerID)
	customerID, err := strconv.Atoi(userCustomerID)
	Check(err)
	latestCustomerId, err := strconv.Atoi(customers[(len(customers) - 1)][0])
	Check(err)
	if customerID > latestCustomerId {
		fmt.Println()
		fmt.Println("Error. Customer ID entered does not correspond to an existing customer.")
	} else {
		LastCouponDetails(customerID, coupNotifHist, customers)
	}
}
