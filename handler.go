package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type EventRequest struct {
	N int64 `json:"n"`
}
type EventResponse struct {
	Result uint64 `json:"result"`
}

func handler(ctx context.Context, e EventRequest) (EventResponse, error) {
	fibN, err := fibonacci(e.N)
	if err != nil {
		return EventResponse{}, err
	}
	return EventResponse{Result: fibN}, nil
}

func apiGWHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var evtReq EventRequest
	err := json.Unmarshal([]byte(req.Body), &evtReq)
	if err != nil {
		return apiGWResponse(0, http.StatusInternalServerError, err), nil
	}
	fibN, err := fibonacci(evtReq.N)
	if err != nil {
		if errors.Is(err, ErrInvalidInput) {
			return apiGWResponse(0, http.StatusBadRequest, err), nil
		}
		return apiGWResponse(0, http.StatusInternalServerError, err), nil
	}
	return apiGWResponse(fibN, http.StatusOK, nil), nil
}

func apiGWResponse(result uint64, status int, err error) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	if err != nil {
		resp.Body = err.Error()
		return resp
	}

	evtResp := EventResponse{
		Result: result,
	}
	data, err := json.Marshal(evtResp)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = err.Error()
	}
	resp.Body = string(data)
	return resp
}
