package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		maxJolt := 0

		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				a := int(line[i] - '0')
				b := int(line[j] - '0')
				val := a*10 + b
				if val > maxJolt {
					maxJolt = val
				}
			}
		}
		total += maxJolt
	}
	fmt.Println(total)
}
