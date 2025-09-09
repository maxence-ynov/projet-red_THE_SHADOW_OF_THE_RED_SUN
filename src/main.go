package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("")
	input, _ := reader.ReadString('\n') // lit jusqu'au retour chariot
	input = strings.TrimSpace(input)    // supprime \n et espaces
	fmt.Println("Tu as écrit :", input)
}
