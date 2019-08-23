package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	fmt.Println("Data atual é:", dateStr)

	dateHoraStr := now.Format("2006-01-02 15:04:02")
	fmt.Println("Data e hora atual é:", dateHoraStr)

}
