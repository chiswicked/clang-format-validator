package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if !isCommandAvailable("clang-format") {
		fmt.Fprintln(os.Stderr, "clang-format missing")
		os.Exit(1)
	}
	protos, err := listFilesWithExt(".", ".proto")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, p := range protos {
		formatProtoFile(p)
	}
	fmt.Println(protos)
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/bash", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func listFilesWithExt(root string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				if filepath.Ext(path) == ext {
					files = append(files, path)
				}
			}
			return nil
		})
	if err != nil {
		return []string{}, err
	}
	return files, nil
}

func formatProtoFile(p string) {
	fmt.Println(p)
	out, err := exec.Command("/bin/bash", "-c", "clang-format -style=google -i "+p).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}
