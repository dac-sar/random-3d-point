package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.ibm.com/Tomonori-Mukai1/random-3d-point/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

//interface実装のため
type server1 struct {
	pb.UnimplementedRandom3DPointServiceServer
}

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRandom3DPointServiceServer(server, &server1{})

	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(server)
	// サーバーを起動
	server.Serve(listenPort)
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
