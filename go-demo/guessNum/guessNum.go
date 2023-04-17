package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxNum := 100

	randomNum := rand.Intn(maxNum)

	fmt.Println(randomNum)

	fmt.Println("please input your number:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error1", err)
		return
	}

	input = strings.TrimSuffix(input, "\r\n")

	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error2", err)
		return
	}
	fmt.Println("your guess number is:", guess)
}