package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/repository"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	//total, err := tickets.GetTotalTickets("Brazil")

	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	fileName := "tickets.csv"

	lstTickets, err := repository.GetTickets(fileName)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Total tickets retrieved :%v \n", len(lstTickets))

	var destination = "China"
	count, _ := tickets.GetTotalTickets(destination, &lstTickets)
	fmt.Printf("Total tickets to detination [%v] are [%v] \n", destination, count)

	destination = "Brazil"
	count, _ = tickets.GetTotalTickets(destination, &lstTickets)
	fmt.Printf("Total tickets to detination [%v] are [%v] \n", destination, count)

	destination = "Honduras"
	count, _ = tickets.GetTotalTickets(destination, &lstTickets)
	fmt.Printf("Total tickets to detination [%v] are [%v] \n", destination, count)

	from := "00:00"
	to := "01:00"
	fmt.Printf("Getting tickets from [%v] to [%v] ...\n", from, to)
	count, _ = tickets.GetCountByPeriod(from, to, &lstTickets)
	fmt.Printf("Total tickets from [%v] to [%v] are [%v] \n", from, to, count)

	avg, _ := tickets.AverageDestination(destination, &lstTickets)
	fmt.Printf("Avg tickets to [%v] is [%f], total %v \n", destination, avg, len(lstTickets))
	destination = "China"
	avg, _ = tickets.AverageDestination(destination, &lstTickets)
	fmt.Printf("Avg tickets to [%v] is [%f], total %v \n", destination, avg, len(lstTickets))

}
