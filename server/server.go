package main

import (
	"context"
	"fmt"
	"gae-grpc-example/echo"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func main() {
	grpcSrv := grpc.NewServer()
	echo.RegisterEchoServiceServer(grpcSrv, &echoService{})
	wrappedServer := grpcweb.WrapServer(grpcSrv, grpcweb.WithOriginFunc(func(origin string) bool {
		// 手抜き。本当は引数で受け取ったoriginを元にアクセス許可の処理を書く必要あり。
		return true
	}))

	http.Handle("/", wrappedServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *echo.EchoRequest) (res *echo.EchoResponse, err error) {
	res = &echo.EchoResponse{
		Message: req.Message,
	}
	return
}
