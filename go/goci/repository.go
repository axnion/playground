package main

import (
	"log"
	"os"
	"os/exec"

	"gopkg.in/src-d/go-git.v4"
)

func getRepository(payload *payload, clone iClone) {
	path := os.Getenv("GOPATH") + "/src/tmp/" + payload.After
	log.Printf("Cloning into %s ...", path)

	_, err := clone.PlainClone(path, false, &git.CloneOptions{
		URL:      payload.Repository.Clone_url,
		Progress: os.Stdout,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Successful clone into " + path)

	clone.Chdir(path)
	executeCommand(exec.Command("go", "get", "."))
}

type iClone interface {
	PlainClone(string, bool, *git.CloneOptions) (*git.Repository, error)
	Chdir(string) error
}

type clone struct{}

func (clone) Chdir(dir string) error {
	return os.Chdir(dir)
}

func (clone) PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	return git.PlainClone(path, isBare, o)
}
