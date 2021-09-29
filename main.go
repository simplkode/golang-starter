package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("ERROR!")
		return
	}
	timesToPrint, err := strconv.Atoi(os.Args[1])
	if err != nil || timesToPrint < 0 || timesToPrint > 10 {
		fmt.Println("ERROR!")
		return
	}
	for count := 0; count < timesToPrint; count += 1 {
		fmt.Println("famous-coyote")
	}
}
