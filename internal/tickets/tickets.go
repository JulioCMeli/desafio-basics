package tickets

import (
	"fmt"
	"time"
)

type Ticket struct {
	Id                 string
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              string
}

// ejemplo 1
func GetTotalTickets(destination string, lstTickets *[]Ticket) (int, error) {

	var count = 0
	for _, t := range *lstTickets {
		if t.DestinationCountry == destination {
			count++
		}
	}
	return count, nil

}

// ejemplo 2
func GetCountByPeriod(from string, to string, lstTickets *[]Ticket) (int, error) {
	var count = 0

	formatDate := "2006-01-02 15:04"
	defaultDate := "2024-01-01 "

	fromString := defaultDate + from
	toString := defaultDate + to
	// fmt.Printf("From: %v \n", fromString)
	// fmt.Printf("To: %v \n", toString)

	fromDate, err := time.Parse(formatDate, fromString)
	toDate, _ := time.Parse(formatDate, toString)
	if err != nil {
		panic(err)
	}

	for _, t := range *lstTickets {
		t_date_str := defaultDate + t.FlightTime
		t_date, _ := time.Parse(formatDate, t_date_str)

		if t_date.After(fromDate) && t_date.Before(toDate) {
			count++
			fmt.Printf("\t %v \n", t)
		}
	}

	return count, nil
}

// ejemplo 3
func AverageDestination(destination string, lstTickets *[]Ticket) (float64, error) {
	count, _ := GetTotalTickets(destination, lstTickets)
	total := len(*lstTickets)
	if count < 1 {
		return 0.0, nil
	}

	var avg float64 = float64(count) / float64(total)
	// fmt.Printf(" AverageDestination total %v \n", total)
	// fmt.Printf(" AverageDestination count %v \n", count)
	// fmt.Printf(" AverageDestination avg %f \n", avg)

	return avg, nil
}
