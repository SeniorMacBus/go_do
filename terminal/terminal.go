package terminal

import (
	"bufio"
	"fmt"
	"os"
)

func PrintToTerminal(line string) error {
	_, err := fmt.Printf(">>> %s", line)

	return err
}

func ReadCommand() (line string, err error) {
	reader := bufio.NewReader(os.Stdin)
	line, err = reader.ReadString('\n')
	return line, err
}

func Session() {
	var active bool = true
	for {
		PrintToTerminal("")
		line, err := ReadCommand()
		if err != nil {
			fmt.Println(err)
			active = false
		}
		PrintToTerminal(line)
		if !active {
			break
		}
	}
}
