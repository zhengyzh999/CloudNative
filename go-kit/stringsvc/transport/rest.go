package transport

import (
	"context"
	"encoding/json"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"stringsvc/endpoint"
	"stringsvc/service"
)

func RestRun() {
	svc := service.StringService{}
	r := makeHttpHandler(svc)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func makeHttpHandler(svc service.StringService) http.Handler {
	router := mux.NewRouter()
	endpoints := endpoint.MakeEndpoints(svc)
	router.Methods("GET").Path("/uppercase").Handler(httpTransport.NewServer(
		endpoints.GetUpperCase,
		decodeUpperCaseRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/uppercase").Handler(httpTransport.NewServer(
		endpoints.PostUpperCase,
		decodeUpperCaseRequest,
		encodeResponse,
	))
	router.Methods("GET").Path("/count").Handler(httpTransport.NewServer(
		endpoints.GetCount,
		decodeCountRequest,
		encodeResponse,
	))
	return router
}

func decodeUpperCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.UpperCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
