package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if !isCommandAvailable("clang-format") {
		fmt.Fprintln(os.Stderr, "clang-format missing")
		os.Exit(1)
	}

	out, err := exec.Command("/bin/bash", "-c", "clang-format -style=google -i *.proto").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/bash", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
