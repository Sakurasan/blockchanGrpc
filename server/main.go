package main

import (
	pb "blockChainRpc/proto"
	"blockChainRpc/server/blockchain"
	"encoding/json"
	"log"
	"net"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type Server struct {
	Server *blockchain.BlockChain
}

func newsrv() *Server {
	return new(Server)
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("监听服务地址失败")
	}

	svr := grpc.NewServer()
	pb.RegisterBlockChainServer(svr, &Server{
		Server: blockchain.NewBlockChain(),
	})

	svr.Serve(listener)
}

func (t *Server) AddBlock(ctx context.Context, in *pb.AddBlockRequest) (*pb.AddBlockResponse, error) {
	resp := new(pb.AddBlockResponse)
	t.Server.AddBlock(in.Data)

	byteStr, err := json.Marshal(t.Server)
	if err != nil {
		return nil, err
	}
	resp.Hash = string(byteStr)
	return resp, nil
}

func (t *Server) GetBlock(ctx context.Context, in *pb.GetBlockRequest) (*pb.GetBlockResponse, error) {
	resp := new(pb.GetBlockResponse)
	byresp, err := json.Marshal(t.Server)
	if err != nil {
		return nil, err
	}
	resp.Blocklist = string(byresp)
	return resp, nil
}

// // ptrof
// func ptrof(p interface{}, fieldName string) unsafe.Pointer {
// 	V := reflect.ValueOf(p).Elem()
// 	field := v.FieldByName(fieldName)
// 	return unsafe.Pointer(field.UnsafeAddr())
// }
