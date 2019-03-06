package client

import (
	"blockchainRpc/proto"
	"flag"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.BlockChainClient

func main() {
	fmt.Println("BlockChainRpc Client start")

	addFlag := flag.Bool("add", false, "Add new block")
	//data := flag.String("data", "btc", "btc data")
	listFlag := flag.Bool("list", false, "List all blocks")
	flag.Parse()

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		panic("remote Server not use")
	}

	client := proto.NewBlockChainClient(conn)

	if *addFlag {
		rsp, _ := client.AddBlock(context.Background(), &proto.AddBlockRequest{Data: time.Now().Format("20060102150405")})
		fmt.Println(rsp)
	}

	if *listFlag {
		rsp, _ := client.GetBlock(context.Background(), nil)
		fmt.Println("返回内容:")
		fmt.Println(rsp.Blocklist)
	}
}
