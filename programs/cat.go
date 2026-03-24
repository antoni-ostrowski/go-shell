package programs

import (
	"fmt"
	"os"
)

type CatRunner struct {
}

func (r *CatRunner) Exec(stdIn []string) (*os.File, *os.File) {
	err := os.Chdir(stdIn[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd error: %v\n", err)
	}
	return nil, nil
}
