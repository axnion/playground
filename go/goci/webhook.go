package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type payload struct {
	After      string
	Repository repository
}

type repository struct {
	Clone_url string
	Ssh_url   string
}

var secret = []byte(os.Getenv("WEBHOOK_SECRET"))

func handleWebhook(res http.ResponseWriter, req *http.Request) (*payload, error) {
	payload, err := parseRequest(req)

	res.Header().Set("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	log.Println(payload.Repository.Clone_url)
	res.WriteHeader(http.StatusOK)
	return payload, nil
}

func parseRequest(req *http.Request) (*payload, error) {
	signature := req.Header.Get("X-Hub-Signature")
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	if signature == "" {
		return nil, errors.New("Missing signature")
	}

	if !verifySignature(secret, signature, body) {
		return nil, errors.New("Invalid signature")
	}

	parsedJSON := new(payload)

	// var prettyJSON bytes.Buffer
	// _ = json.Indent(&prettyJSON, body, "", "\t")
	// fmt.Println(string(prettyJSON.Bytes()))

	err = json.Unmarshal(body, parsedJSON)

	if err != nil {
		return nil, err
	}

	return parsedJSON, nil
}

func verifySignature(secret []byte, signature string, body []byte) bool {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(body)

	expectedMac := mac.Sum(nil)
	expectedSig := "sha1=" + hex.EncodeToString(expectedMac)

	return hmac.Equal([]byte(expectedSig), []byte(signature))
}
