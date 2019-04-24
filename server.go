package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "helloWorld/pb"
	"net"
	"net/http"
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

// 这个方法里面的token认证，只是保存在某个方法里面的，如果有大量方法的话，在每个方法里面写大量的认证很麻烦
func (u *user) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {

	//md, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	//}
	//var (
	//	appuid string
	//	appkey string
	//)
	//if val, ok := md["appuid"]; ok {
	//	appuid = val[0]
	//}
	//if val, ok := md["appkey"]; ok {
	//	appkey = val[0]
	//}
	//if appuid != "100" || appkey != "i am key" {
	//	return nil, status.Errorf(codes.Unauthenticated, "Token认证失败")
	//}

	return &pb.RegisterReply{Uid: "token认证成功"}, nil
}

// auth 验证Token
func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appuid string
		appkey string
	)
	if val, ok := md["appuid"]; ok {
		appuid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appuid != "100" || appkey != "i am key" {
		return status.Errorf(codes.Unauthenticated, "Token认证失败")
	}
	return nil
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
	grpc.EnableTracing = true
	var (
		opts []grpc.ServerOption
	)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatal(err)
	}

	// 添加TLS认证
	creds, err := credentials.NewServerTLSFromFile("./ssl/server.pem", "./ssl/server.key")
	if err != nil {
		grpclog.Fatal(err)
	}
	opts = append(opts, grpc.Creds(creds))

	// 注册interceptor
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		// 继续处理其他请求
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	s := grpc.NewServer(opts...)
	// 注册 GreeterService
	pb.RegisterGreeterServer(s, &server{})
	// 注册 UserService
	pb.RegisterUserServer(s, &user{})
	grpclog.Infof("Listen on %s with TLS", port)

	// 开启trace
	go startTrace()
	err = s.Serve(lis)
	if err != nil {
		grpclog.Fatal(err)
	}
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":50052", nil)
	grpclog.Infoln("Trace listen on 50052")
}
