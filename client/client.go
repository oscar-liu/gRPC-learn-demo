package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/pb"
	"io"
	"log"
	"time"
)

func sayUnaryClient() {
	// 创建一个 grpc 连接
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 创建 grpc 客户端
	client := pb.NewGreetsClient(conn)
	// 设置超时时间
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用方法
	sendMessage := &pb.RequestMessage{
		Name: "client1", Message: "server 大爷，收到我信息了没有，收到了请吱一声",
	}
	log.Println("client send:", sendMessage)
	reply, err := client.SayUnaryHello(context.Background(), sendMessage)
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
	}
	log.Printf("serve reply：name: %s, message: %s\n", reply.Name, reply.Message)
}

func sayStreamClient() {
	// 创建一个 grpc 连接
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 创建 grpc 客户端
	client := pb.NewGreetsClient(conn)
	// 设置超时时间
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用方法
	sendMessage := &pb.RequestMessage{
		Name: "client1", Message: "server 大爷，晚上吃什么？",
	}
	log.Println("client send:", sendMessage)
	reply, err := client.SayServerStreamHello(context.Background(), sendMessage)
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
	}

	for {
		recv, err := reply.Recv()
		if err != nil {
			log.Fatalf("couldn not greet: %v", err)
			break
		}
		fmt.Println(recv)
	}

}

func sayClientStreamHello() {
	// 创建一个 grpc 连接
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 创建 grpc 客户端
	client := pb.NewGreetsClient(conn)
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//调用rpc方法，得到一个客户端用于循环发送数据
	gclient, err := client.SayClientStreamHello(ctx)

	requestNumber := 2 // 请求2次
	for i := 0; i < requestNumber; i++ {
		// 调用方法
		sendMessage := &pb.RequestMessage{
			Name: "client1", Message: "大爷你好？",
		}
		sendMessage2 := &pb.RequestMessage{
			Name: "client1", Message: "大爷在吗？",
		}

		if i == 0 {
			err = gclient.Send(sendMessage)
			log.Println("client send:", sendMessage)
		} else {
			err = gclient.Send(sendMessage2)
			log.Println("client send:", sendMessage2)
		}
		if err != nil {
			log.Fatalf("couldn not greet: %v", err)
			return
		}
		if i >= 1 {
			res, err := gclient.CloseAndRecv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(res)
			break
		}
	}

}

func sayBiDirectionalStreamHello() {
	// 创建一个 grpc 连接
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 创建 grpc 客户端
	client := pb.NewGreetsClient(conn)
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//调用rpc方法，得到一个客户端用于循环发送数据
	gclient, err := client.SayBiDirectionalStreamHello(ctx)

	requestNumber := 2 // 请求2次
	for i := 0; i < requestNumber; i++ {
		// 调用方法
		sendMessage := &pb.RequestMessage{
			Name: "client1", Message: "大爷你好？晚上吃什么？",
		}
		sendMessage2 := &pb.RequestMessage{
			Name: "client1", Message: "大爷在吗？晚上吃什么？",
		}

		if i == 0 {
			err = gclient.Send(sendMessage)
			log.Println("client send:", sendMessage)
		} else {
			err = gclient.Send(sendMessage2)
			log.Println("client send:", sendMessage2)
		}
		if err != nil {
			log.Fatalf("couldn not greet: %v", err)
			return
		}
		if i >= 1 {
			err := gclient.CloseSend()
			if err != nil {
				fmt.Println(err)
				break
			}
			break
		}
	}

	// 接收
	for {
		recv, err := gclient.Recv()
		if err == io.EOF {
			log.Println("数据接收完毕")
			break
		}
		fmt.Println(recv)
	}
}

func main() {
	//sayUnaryClient()

	//sayStreamClient()

	//sayClientStreamHello()

	sayBiDirectionalStreamHello()
}
