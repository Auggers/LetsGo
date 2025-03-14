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

func inputCalcPart1(input []int) []int {
	//Entering forloop to loop through the slice
	for i := 0; i < len(input); i += 4 {
		if input[i] == 99 {
			//if data is 99 the code exits the for loop
			break
		} else if input[i] == 1 {
			//if data is 2 it multiplies position 1 and 2 and replaces 3rd position with answer. It then restarts the loop.
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
			//if data is 2 it multiplies position 1 and 2 and replaces 3rd position with answer. It then restarts the loop.
		} else if input[i] == 2 {
			//if data is 2 it adds position 1st and 2nd and replaces 3rd position with answer. It then restarts the loop.
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		}
	}
	//returning the new array with changed values
	return input
}

func inputCalcPart2(input []int) int {
	//Entering forloop to loop through the slice
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			newSlice := make([]int, len(input))
			copy(newSlice, input)

			newSlice[1], newSlice[2] = noun, verb

			newSlice2 := inputCalcPart1(newSlice)
			if newSlice2[0] == 19690720 {
				answer := 100*noun + verb
				return answer
			}

		}
	}
	//returning the new array with changed values
	return 0
}

func main() {
	//Opens the input.csv file
	f, err := os.Open("input.csv")

	//Prints out error if something goes wrong
	var stderr bytes.Buffer
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	//creates a new reader which takes the input from input.csv
	r := csv.NewReader(f)
	var data []int
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
			intRecord, err := strconv.Atoi(stringRecord)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, intRecord)
		}

		part2Data := make([]int, len(data))
		copy(part2Data, data)
		data[1], data[2] = 12, 2
		//Printing Part 1 answer
		fmt.Printf("\nAnswer for part 1: %d\n\n", inputCalcPart1(data)[0])

		//Printing Part 2 answer
		fmt.Printf("\nAnswer for part 2: %d\n\n", inputCalcPart2(part2Data))

	}
}
