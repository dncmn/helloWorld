package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	pb "helloWorld/pb"
	"net"
	"strings"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	msg := "login error"
	color := make([]string, 0)
	rewardMap := make(map[int32]bool)
	DateList := make(map[int32]*pb.List)
	if in.Username == "root" && in.Password == "123456" {
		msg = "login success"
		color = []string{"black", "red", "yellow"}
		rewardMap[1001] = true
		rewardMap[1002] = false
		rewardMap[1003] = false
		rewardMap[1004] = true
		DateList = map[int32]*pb.List{
			1: &pb.List{Id: []int32{1, 2, 3, 4}},
		}
	}
	fmt.Println(msg, color, DateList)
	return &pb.LoginReply{}, nil
}

func (s *server) CanSet(ctx context.Context, in *pb.CanSetRequest) (*pb.CanSetReply, error) {
	return &pb.CanSetReply{CanSet: true}, nil
	//os.Exit(0)
}

type user struct{}

func (u *user) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{Uid: "123456"}, nil
}

func (u *user) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}

func (u *user) UserByUID(ctx context.Context, in *pb.UserByUIDRequest) (*pb.UserByUIDReply, error) {
	return &pb.UserByUIDReply{}, nil
}

func (s *server) CanUpdate(ctx context.Context, in *pb.CanUpdateRequest) (resp *pb.CanUpdateReply, err error) {
	resp = &pb.CanUpdateReply{CanUpdate: false}
	if !strings.Contains(in.Username, "root") {
		resp.CanUpdate = true
	}

	if strings.Contains(in.Username, "err") {
		//err = errors.New(fmt.Sprintf("username=%v is illegal", in.Username))
		err = errors.New("abcd")
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatal(err)
	}

	// 添加TLS认证
	creds, err := credentials.NewServerTLSFromFile("./ssl/server.pem", "./ssl/server.key")
	if err != nil {
		grpclog.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterUserServer(s, &user{})
	grpclog.Infof("Listen on %s with TLS", port)
	err = s.Serve(lis)
	if err != nil {
		grpclog.Fatal(err)
	}
}
