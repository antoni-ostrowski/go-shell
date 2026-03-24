package main

import "os"

type Program struct {
	runner    Executable
	getStdOut func() *os.File
	getStdErr func() *os.File
}

func (p *Program) Run(stdIn []string) (stdOut *os.File, stdErr *os.File) {
	return p.runner.Exec(stdIn)
}

type Executable interface {
	Exec([]string) (*os.File, *os.File)
}
