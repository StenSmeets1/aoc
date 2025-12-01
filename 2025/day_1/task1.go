package main

import ( 
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 50
	zeroCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
	
		if direction == 'L' {
			position -= distance
		} else if direction == 'R' {
			position += distance
		}
		
		position %= 100
		if position < 0 {
			position += 100
		}

		if position == 0 {
			zeroCount++
		}
	}

	fmt.Println("pass: ", zeroCount)
}
