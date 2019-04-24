package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "helloWorld/pb"
	"log"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {

	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("./ssl/server.pem", "CN")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
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
	//switch  pb.Profile{}.Avatar.(type){
	//case *pb.Profile_ImageData:
	//		fmt.Println("aaaaa")
	//case *pb.Profile_ImageUrl:
	//		fmt.Println("bbbbb")
	//default:
	//		fmt.Println("cccc")
	//
	//}
	//t := pb.SearchRequest{}
	//t.List = make(map[int32]*pb.ListDate)
	//l1 := &pb.ListDate{}
	//l1.List = make(map[int32]int32)
	//l1.List[1] = 1
	//l1.List[2] = 1
	//l1.List[3] = 1
	//l1.List[4] = 1
	//
	//l2 := &pb.ListDate{}
	//l2.List = make(map[int32]int32)
	//l2.List[1] = 10
	//l2.List[2] = 10
	//l2.List[3] = 10
	//l2.List[4] = 10
	//t.List = map[int32]*pb.ListDate{
	//	1: l1,
	//	2: l2,
	//}
	//fmt.Println(t.List)
	//for key, val := range t.List {
	//	fmt.Printf("key=%v,val=%v\n", key, val)
	//	for boxID, res := range val.List {
	//		fmt.Printf("----------boxID=%v,result=%v\n", boxID, res)
	//	}
	//}
}
