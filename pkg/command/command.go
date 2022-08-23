package command

import (
	"bufio"
	"fmt"
	"os/exec"
)

// Runs the given command with the given arguments and returns the output of the given executable.
func Run(cmd string, args ...string) (output string, err error) {
	c := exec.Command(cmd, args...)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return
	}
	if err = c.Start(); err != nil {
		return
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		output += fmt.Sprintln(scanner.Text())
	}
	if err = c.Wait(); err != nil {
		return
	}
	return
}
