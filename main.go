package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("clang-format", "--help").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}
