package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func prepare(input string) (repo, path string) {
	u, err := url.Parse(input)
	if err != nil {
		return input, input
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}
	if !strings.HasSuffix(u.Path, ".git") {
		u.Path = u.Path + ".git"
	}
	repo = u.String()

	if strings.HasSuffix(u.Path, ".git") {
		path = u.Host + strings.Replace(u.Path, ".git", "", -1)
	}

	return repo, path
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: gitclone <repo>")
		os.Exit(1)
	}

	repo, path := prepare(os.Args[1])
	var cmd *exec.Cmd

	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		dir := filepath.Join(gopath, "src", path)
		cmd = exec.Command("git", "clone", repo, dir)
	} else {
		fmt.Println("You didn't set GOPATH before, so just clone directly.")
		cmd = exec.Command("git", "clone", repo)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
