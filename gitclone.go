package main

import (
	"errors"
	"fmt"
	"go/build"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func prepare(target string) (repo, path string) {
	if strings.Contains(target, "@") {
		return prepareSSH(target)
	}
	return prepareHTTPS(target)
}

func prepareSSH(target string) (repo, path string) {
	s := strings.Split(target, "@")[1]

	path = strings.Replace(s, ":", "/", -1)
	if strings.HasSuffix(path, ".git") {
		path = strings.Replace(path, ".git", "", -1)
	}
	return target, path
}

func prepareHTTPS(target string) (repo, path string) {
	u, err := url.Parse(target)
	if err != nil {
		return target, target
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

func getFirstDir(gopath string) (string, error) {
	buildContext := build.Default
	list := filepath.SplitList(buildContext.GOPATH)
	if len(list) == 0 {
		return "", errors.New("no gopath set")
	}
	// Guard against people setting GOPATH=$GOROOT.
	if list[0] == buildContext.GOROOT {
		return "", errors.New("gopath can not be goroot")
	}

	return list[0], nil
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
		gopath, err := getFirstDir(gopath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		dir := filepath.Join(gopath, "src", path)
		cmd = exec.Command("git", "clone", repo, dir)
	} else {
		fmt.Println("You didn't set GOPATH before, so just clone directly.")
		cmd = exec.Command("git", "clone", repo)
	}
	fmt.Printf("repo: %s\npath: %s\n", repo, path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
