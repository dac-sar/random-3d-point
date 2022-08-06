package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.ibm.com/Tomonori-Mukai1/random-3d-point/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

//interface実装のため
type server1 struct {
	pb.UnimplementedRandom3DPointServiceServer
}

func main() {

	grpcServer := grpc.NewServer()

	wrappedServer := grpcweb.WrapServer(
		grpcServer,
		// CORSの設定
		grpcweb.WithOriginFunc(func(origin string) bool {
			return origin == "http://localhost:8080"
		}),
	)
	pb.RegisterRandom3DPointServiceServer(grpcServer, &server1{})
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(wrappedServer.ServeHTTP))
	// ポート50051で起動
	hs := &http.Server{
		Addr:    ":50051",
		Handler: mux,
	}
	log.Fatal(hs.ListenAndServe())
}

//interfaceの実装
func (s *server1) Get3DVector(context.Context, *emptypb.Empty) (*pb.Random3DVector, error) {
	now := time.Now()
	return &pb.Random3DVector{
		X: float32(rand.Intn(3) + 1),
		Y: float32(rand.Intn(3) + 1),
		Z: float32(rand.Intn(3) + 1),
		CreateTime: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}, nil
}
