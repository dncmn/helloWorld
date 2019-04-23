package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "helloWorld/pb"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//
	//r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("greeting:", r.Message)
	//loginReply, err := c.Login(context.Background(), &pb.LoginRequest{Name: "root", Password: "123456"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	////fmt.Println(loginReply.Message)
	////fmt.Println(loginReply.Color)
	////fmt.Println(loginReply.RewardMap)
	//fmt.Println(loginReply.DateList)
	//
	//canSetResp, err := c.CanSet(context.Background(), &pb.CanSetRequest{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(canSetResp.CanSet)
	//
	//re, err := c.CanUpdate(context.Background(), &pb.CanUpdateRequest{Username: "root"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("final=", re)
	//fmt.Println(pb.UserRegisterType_GuestRegister, pb.UserRegisterType_NormalRegister)

	u := pb.NewUserClient(conn)
	r, err := u.Register(context.Background(), &pb.RegisterRequest{Username: "manan", Password: "123456", Country: 1, PhoneNum: "15737345574"})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(r.Uid)

}
