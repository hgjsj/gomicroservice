package transport

import (
	"context"
	"encoding/json"
	"go-microservice/model"
	"net/http"
)

func DecodeVMRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.VirtualMachine
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeCloudResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeDiskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.Disk
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
