package main

import (
	"fmt"
	"log"
	"strconv"

	calculator "github.com/SzymekN/currency-exchange-calculator/pkg/calculator"
)

func main() {

	d := calculator.DefaultHttpGetter{}
	for {
		mid, err := calculator.GetCurrentRate(d, "GBP", calculator.GBPDefaultURL)

		if err != nil {
			log.Fatal(err)
		}

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

			v, err := calculator.CalculateReceivedAmount(value, mid)

			if err != nil {
				fmt.Println("Could not calculate how much to send, continuing")
			}

			fmt.Println("They will receive: " + strconv.FormatFloat(v, 'f', -1, 64) + "GBP")

		case 2:
			fmt.Print("Enter amount of GBP you want to get: ")
			_, err = fmt.Scanln(&value)

			if err != nil {
				fmt.Println("Invalid input")
				continue
			}

			v := calculator.CalculateSentAmount(value, mid)

			fmt.Println("You will have to send: " + strconv.FormatFloat(v, 'f', -1, 64) + "PLN")

		default:
			fmt.Println("Invalid choice")

		}

	}

}
