package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, conversionErr := strconv.Atoi(n)
		if conversionErr != nil {
			fmt.Println(conversionErr)
		} else {
			sum += num
		}
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println(inputErr)
		}
	}

	fmt.Println(sum)
}
