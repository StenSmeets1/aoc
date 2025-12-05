package main

import (
	"bufio"
	"fmt"	
	"math/big"
	"os"
	"strings"
)

func maxSubsequence(line string, k int) string {
	stack := []byte{}
	remaining := len(line)

	for i := 0; i < len(line); i++ {
		c := line[i]
		for len(stack) > 0 && stack[len(stack)-1] < c && len(stack)-1+remaining >= k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		remaining--
	}

	if len(stack) > k {
		stack = stack[:k]
	}

	return string(stack)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := big.NewInt(0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue 
		}
		numStr := maxSubsequence(line, 12)

		num := new(big.Int)
		num.SetString(numStr, 10)
		total.Add(total, num)
	}
	fmt.Println(total)
}

