package repository

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	//"github.com/bootcamp-go/desafio-go-bases/internal/tickets"

	. "github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func GetTickets(filePath string) ([]Ticket, error) {

	// open a file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("The indicated file [%v] was not found or is damaged", filePath)
		return nil, errors.New("Error: The indicated file was not found or is damaged")
	}
	// defer close file
	defer file.Close()

	var lstTickets []Ticket

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		//fmt.Println(values)
		ticket, _ := parserValues(values)
		//fmt.Println(ticket, err2)
		lstTickets = append(lstTickets, ticket)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lstTickets, nil
}

func parserValues(values []string) (Ticket, error) {

	var t Ticket

	t.Id = values[0]
	t.Name = values[1]
	t.Email = values[2]
	t.DestinationCountry = values[3]
	t.FlightTime = values[4]
	t.Price = values[5]

	return t, nil
}
