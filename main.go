package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Hello World")

	reqBytes, err := json.Marshal(request)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Error in marshal: %v", err),
		}, err
	}

	log.Print(string(reqBytes))

	var person Person

	err = json.Unmarshal([]byte(request.Body), &person)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Error in unmarshal: %v", err),
		}, err
	}

	message := fmt.Sprintf("Hello %v %v", *person.FirstName, *person.LastName)

	response := Response{
		Message: &message,
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Error in marshalling: %v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseJson),
	}, nil
}

type Person struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type Response struct {
	Message *string `json:"message"`
}
