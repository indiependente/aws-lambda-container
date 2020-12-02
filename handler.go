package main

import "context"

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
