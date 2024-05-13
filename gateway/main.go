package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/myservice/protoservice"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "")

func main() {
	// 网关使用
	flag.Parse()
	mux := runtime.NewServeMux()
	mux.HandlePath("GET", "/ping", pingHandle)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := protoservice.RegisterMyServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func pingHandle(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	fmt.Fprintf(w, "{\"msg\":\"pong\"}")
}
