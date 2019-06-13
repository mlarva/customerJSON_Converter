package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Customer struct {
	ID           string
	CustomerName string
	CustomerKey  string
	OEStartDate  string
	OEEndDate    string
	PlanYear     string
}

func main() {
	var customer Customer
	var customerList []Customer
	csvfile, err := os.Open("customer_key_name_oe_dates.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		customer.ID = record[0]
		customer.CustomerKey = record[1]
		customer.CustomerName = record[2]
		customer.OEStartDate, err = convertStringUnixToTime(record[3])
		if err != nil {
			fmt.Println(err)
		}
		customer.OEEndDate, err = convertStringUnixToTime(record[4])
		if err != nil {
			fmt.Println(err)
		}
		customer.PlanYear = record[5]
		customerList = append(customerList, customer)
		convertStringUnixToTime(customer.OEStartDate)
	}
	jsonCustomerList, err := json.Marshal(customerList)
	if err != nil {
		fmt.Println(err)
	} else {
		f, err := os.Create("output.json")
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write(jsonCustomerList)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func convertStringUnixToTime(dateString string) (string, error) {
	i, err := strconv.ParseInt(dateString, 10, 64)
	if err != nil {
		return "", err
	} else {
		tm := time.Unix(i, 0)
		sTime := tm.Format("2006-01-02")
		return sTime, nil
	}
}
