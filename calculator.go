package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	number, _ := strconv.Atoi(n)

	for i := 0; i < number; i++ {
		fmt.Println("Hello, World!")
	}
}
