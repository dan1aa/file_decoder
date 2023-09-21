package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Command() (string, string) {
	var command string

	fmt.Print("Enter a command: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	command = strings.TrimSpace(input)

	splitted_command := strings.Split(command, " ")

	if len(splitted_command) > 2 {
		return "", ""
	}

	return splitted_command[0], splitted_command[1]
}
