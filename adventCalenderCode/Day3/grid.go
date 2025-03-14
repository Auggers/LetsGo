package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func dataToSlice(input string) []string {
	//Opens the input.csv file
	f, err := os.Open(input)

	//Prints out error if something goes wrong
	var stderr bytes.Buffer
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	//creates a new reader which takes the input from input.csv
	r := csv.NewReader(f)
	var data []string
	//reads the data from each position in each position and stores it in stringRecord
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//loops through record, coverts it to an integer, and adds it to an a slice called data
		for i := 0; i < len(record); i++ {
			stringRecord := record[i]
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, stringRecord)
		}
	}
	return data
}

//Seperates the letters from the numbers in the slice
func parseNum(s string) (letters, numbers string) {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}
	return string(l), string(n)
}

//Adds the seperated numbers and letter from parseNum and adds to new slice
func newSlice(input []string) []string {
	var newSlice []string
	for _, num := range input {
		letters, numbers := parseNum(num)
		newSlice = append(newSlice, letters, numbers)
	}
	return newSlice
}

//Plots all points that have been used into a slice
func arrange(input []string) []int {
	y := 0
	x := 0
	var data []int

	for i := 0; i < len(input); i += 2 {
		if input[i] == "L" {
			stringRecord := input[i+1]
			intRecord, err := strconv.Atoi(stringRecord)
			if err != nil {
				log.Fatal(err)
			}
			z := x
			x -= intRecord
			for z >= x {
				data = append(data, y)
				data = append(data, z)
				z--
			}
		} else if input[i] == "R" {
			stringRecord := input[i+1]
			intRecord, err := strconv.Atoi(stringRecord)
			if err != nil {
				log.Fatal(err)
			}
			z := x
			x += intRecord
			for z <= x {
				data = append(data, y)
				data = append(data, z)
				z++
			}
		} else if input[i] == "D" {
			stringRecord := input[i+1]
			intRecord, err := strconv.Atoi(stringRecord)
			if err != nil {
				log.Fatal(err)
			}
			z := y
			y -= intRecord
			for z >= y {
				data = append(data, z)
				data = append(data, x)
				z--
			}
		} else if input[i] == "U" {
			stringRecord := input[i+1]
			intRecord, err := strconv.Atoi(stringRecord)
			if err != nil {
				log.Fatal(err)
			}
			z := y
			y += intRecord
			for z <= y {
				data = append(data, z)
				data = append(data, x)
				z++
			}
		}
	}
	return data
}

//Compares both wires, if there's a match in coordinates then it adds to a slice
func compare(input1 []int, input2 []int) []int {
	var data []int
	for i := 0; i < len(input1); i += 2 {
		for j := 0; j < len(input2); j += 2 {
			if input1[i] == input2[j] && input1[i+1] == input2[j+1] {
				data = append(data, input1[i])
				data = append(data, input1[i+1])
				fmt.Printf("\nMatch at %d %d", input1[i], input1[i+1])
			} else {
				continue
			}
		}
	}
	return data
}

//converts integer top float
func intToFloat(n int) float64 {
	f := float64(n)
	return f
}

//converts float to Integer
func floatToInt(n float64) int {
	f := int(n)
	return f
}

//converts coordinate to float to be used in math.abs function, calculates answer, and converts it back to an integer to be returned
func calc(input []int) int {
	answer := 1000000
	for i := 2; i < len(input); i += 2 {
		a := intToFloat(input[i])
		b := intToFloat(input[i+1])
		x := math.Abs(a) + math.Abs(b)
		y := floatToInt(x)
		if y < answer {
			answer = y
		}
	}
	return answer
}

func main() {
	wire1Raw := dataToSlice("wire1.csv")
	wire2Raw := dataToSlice("wire2.csv")

	wire1 := newSlice(wire1Raw)
	wire2 := newSlice(wire2Raw)

	wire1Plot := arrange(wire1)
	wire2Plot := arrange(wire2)

	input := compare(wire1Plot, wire2Plot)

	fmt.Printf("\nAnswer for part 1: %d\n\n", calc(input))
}
