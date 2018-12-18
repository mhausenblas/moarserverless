package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// see https://github.com/bmc/fortunes/
	messages := []string{"It's bad luck to be superstitious.", "A city is a large community where people are lonesome together.", "A critic is a legless man who teaches running.", "Dobbin's Law: When in doubt, use a bigger hammer."}
	rand.Seed(time.Now().Unix())
	fmessage := messages[rand.Intn(len(messages))]
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(fmessage),
	}, nil

}

func main() {
	lambda.Start(handler)
}
