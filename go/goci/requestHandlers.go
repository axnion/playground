package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Fprintf(res, "Front page")
	} else if req.Method == "POST" {
		// Handle webhook request
		webhookPayload, err := handleWebhook(res, req)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Clone repository and fetch dependencies
		getRepository(webhookPayload, new(clone))

		// Run tests
		stacktrace, err := executeCommand(exec.Command("go", "test"))
		log.Println(stacktrace)

		if err != nil {
			log.Println(err)
			return
		}

		// Build binary
		executeCommand(exec.Command("go", "install"))
		log.Println(stacktrace)

		if err != nil {
			log.Println(err)
			return
		}
	}
}

func executeCommand(cmd command) (string, error) {
	out, err := cmd.Output()
	return string(out), err
}

/**
 * Interface used in executeCommand method. Allows mocking of exec.Cmd object in Testing
 * executeCommand
 */
type command interface {
	Output() ([]byte, error)
}
