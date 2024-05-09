package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"stringsvc/service"
)

type Endpoints struct {
	GetUpperCase  endpoint.Endpoint
	PostUpperCase endpoint.Endpoint
	GetCount      endpoint.Endpoint
}

func MakeEndpoints(svc service.StringService) Endpoints {
	return Endpoints{
		GetCount:      MakeGetCountEndpoint(svc),
		PostUpperCase: MakePostUpperEndpoint(svc),
		GetUpperCase:  MakeGetUpperEndpoint(svc),
	}
}

type UpperCaseRequest struct {
	S string `json:"s"`
}
type UpperCaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
type CountRequest struct {
	S string `json:"s"`
}
type CountResponse struct {
	V int `json:"v"`
}

func MakeGetUpperEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpperCaseRequest)
		v, err := svc.GetUpperCase(req.S)
		if err != nil {
			return UpperCaseResponse{v, err.Error()}, nil
		}
		return UpperCaseResponse{v, ""}, nil
	}
}

func MakePostUpperEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpperCaseRequest)
		v, err := svc.PostUpperCase(req.S)
		if err != nil {
			return UpperCaseResponse{v, err.Error()}, nil
		}
		return UpperCaseResponse{v, ""}, nil
	}
}

func MakeGetCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.GetCount(req.S)
		return CountResponse{v}, nil
	}
}
