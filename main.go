package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/antoni-ostrowski/go-shell/programs"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	programsMap := make(map[string]Program)
	programsMap["cd"] = Program{
		runner: &programs.CdRunner{},
	}

	for {

		curWorkingDir, err := os.Getwd()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Printf("(%s)> ", curWorkingDir)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		command, args := cleanOutStdIn(input)

		bin, ok := programsMap[command]
		if !ok {
			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			bin.runner.Exec(args)
		}

	}
}

func cleanOutStdIn(input string) (string, []string) {
	cleared := strings.TrimSuffix(input, "\n")
	splitted := strings.Split(cleared, " ")
	return splitted[0], splitted[1:]
}
