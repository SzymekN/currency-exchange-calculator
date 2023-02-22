package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// 	W przypadku braku danych dla prawidłowo określonego zakresu czasowego zwracany jest komunikat 404 Not Found

	// W przypadku zadania nieprawidłowo sformułowanych zapytań serwis zwraca komunikat 400 Bad Request

	// W przypadku zapytania przekraczającego limit zwracanych danych serwis zwróci komunikat 400 Bad Request - Przekroczony limit

	for {
		resp, err := calculator.getCurrentGBPRate()

		if err != nil {
			log.Fatal(err)
		}

		mid := resp.Rates[0].Mid

		midStr := fmt.Sprintf("%f", mid)

		fmt.Println("Current exchange rate: 1GBP = " + midStr + "PLN")
		fmt.Println("\nChoose option:\n1. Calculate from PLN to GBP\n2. Calculate from GBP to PLN")

		choice, value := 0, 0.0
		_, err = fmt.Scanln(&choice)

		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter amount of PLN you want to send: ")
			_, err = fmt.Scanln(&value)

			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Invalid input")
				continue
			}

			fmt.Println("They will receive: " + strconv.FormatFloat(value*mid, 'f', -1, 64) + "GBP")

		case 2:
			fmt.Print("Enter amount of GBP you want to get: ")
			_, err = fmt.Scanln(&value)

			if err != nil {
				fmt.Println("Invalid input")
				continue
			}

			fmt.Println("You will have to send: " + strconv.FormatFloat(value/mid, 'f', -1, 64) + "PLN")

		default:
			fmt.Println("Invalid choice")

		}

	}

}
