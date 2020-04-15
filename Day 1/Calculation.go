package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func calcMass(number int) int {
	mass := 0
	for number > 0 {
		calcMass := number/3 - 2
		mass += calcMass
		break
	}
	return mass
}

func calcFuel(number int) int {
	total := 0
	for {
		if number > 0 {
			fuelCalcDiv := number/3 - 2
			if fuelCalcDiv <= 0 {
				break
			} else {
				number = fuelCalcDiv
				total += fuelCalcDiv
			}
		} else {
			break
		}
	}
	return total
}

func main() {
	f, err := os.Open("input.csv")

	var stderr bytes.Buffer
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	r := csv.NewReader(f)
	mass := 0
	fuel := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		stringRecord := record[0]

		intRecord, err := strconv.Atoi(stringRecord)
		mass += calcMass(intRecord)
		fuel += calcFuel(intRecord)
	}
	fmt.Println("Total mass is: ", mass)
	fmt.Println("Total mass is: ", fuel)
}
