package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/pb"
	"io"
	"log"
	"net"
)

type Server struct{}

// SayUnaryHello 简单 RPC（Unary RPC）
// 一般的 RPC 调用，传入一个请求对象，返回一个返回对象
func (s *Server) SayUnaryHello(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	log.Println(req.Name, req.Message)
	return &pb.ResponseMessage{Name: "服务器大爷", Message: "收到，吱"}, nil
}

// SayServerStreamHello 服务端流式 RPC（Server streaming RPC）
// 传入一个请求对象，服务端可以返回多个结果对象
func (s *Server) SayServerStreamHello(request *pb.RequestMessage, server pb.Greets_SayServerStreamHelloServer) error {
	fmt.Println("request: ", request)
	var err error
	menu := [2]string{"红烧肉", "黄花鱼"}

	for i := 0; i < 2; i++ {
		err = server.Send(&pb.ResponseMessage{Name: "服务器大爷", Message: "晚上吃:" + " " + menu[i]})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// SayClientStreamHello 客户端流式 RPC（Clientstreaming RPC）
// 客户端传入多个请求对象，服务端返回一个结果对象
func (s *Server) SayClientStreamHello(server pb.Greets_SayClientStreamHelloServer) error {
	for {
		recv, err := server.Recv()
		// 接收完所有数据之后再响应
		if err == io.EOF {
			err := server.SendAndClose(&pb.ResponseMessage{Name: "服务器大爷", Message: "太哆嗦了，结束"})
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		} else if err != nil {
			return err
		}
		fmt.Println("recv: ", recv)
	}
}

// SayBiDirectionalStreamHello 双向流式 RPC（Bi-directional streaming RPC）
// 结合客户端流式RPC和服务端流式RPC，可以传⼊多个请求对象，返回多个结果对
func (s *Server) SayBiDirectionalStreamHello(server pb.Greets_SayBiDirectionalStreamHelloServer) error {
	for {
		recv, err := server.Recv()
		// 接收完所有数据之后再响应
		if err == io.EOF {
			fmt.Println("数据接收完毕！")
			// 返回两个菜
			menu := [2]string{"红烧肉", "黄花鱼"}
			err = server.Send(&pb.ResponseMessage{Name: "服务器大爷", Message: "晚上吃:" + " " + menu[0]})
			err = server.Send(&pb.ResponseMessage{Name: "服务器大爷", Message: "晚上吃:" + " " + menu[1]})
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		} else if err != nil {
			return err
		}
		fmt.Println("recv: ", recv)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error listening", err)
		return
	}

	// 定义一个 rpc 的 server
	server := grpc.NewServer()
	// 注册服务，注册 sayHello 接口
	pb.RegisterGreetsServer(server, &Server{})
	// 映射绑定
	reflection.Register(server)

	// 启动服务
	err = server.Serve(listen)
	if err != nil {
		fmt.Println("Error serving", err)
		return
	}
}
