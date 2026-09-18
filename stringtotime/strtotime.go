package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type date struct {
	Date time.Time
}

func main() {
	var err error
	dte := date{}
	dateStr := "2019-08-31"
	dte.Date, err = convertStringToTime(dateStr)
	if err != nil {
		log.Println(fmt.Sprintf("date [%s], error: [%s]", dte.Date.Format("2006-01-02"), err))
	}

	log.Println("Date", dte.Date)
}

func convertStringToTime(dte string) (time.Time, error) {
	date := strings.Split(dte, "-")
	if len(date) != 3 {
		return time.Now(), errors.New("Invalid date. Needs to be YYYY-MM-DD")

	}
	y, err := strconv.Atoi(date[0])
	if err != nil {
		return time.Now(), err
	}

	m, err := strconv.Atoi(date[1])
	if err != nil {
		return time.Now(), err
	}

	d, err := strconv.Atoi(date[2])
	if err != nil {
		return time.Now(), err
	}

	return time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC), nil

}
