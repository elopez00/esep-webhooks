package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Issue struct {
	HTMLURL string `json:"html_url"`
}

type RequestPayload struct {
	Issue Issue `json:"issue"`
}

func HandleRequest(ctx context.Context, payload RequestPayload) (string, error) {
	// get the link from the payload
	link := payload.Issue.HTMLURL

	// create the message to send to slack
	data := []byte("{'text': 'Issue Created: " + link + "'}")

	// make a request to the webhook under the SLACK_URI env var
	res, err := http.Post(os.Getenv("SLACK_URL"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// read the contents of the response if request successful
	buff, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// return the response
	return string(buff), nil
}

func main() {
	lambda.Start(HandleRequest)
}
